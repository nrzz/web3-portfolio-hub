package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/config"
	"web3-portfolio-dashboard/backend/internal/services"
)

type Server struct {
	engine           *gin.Engine
	config           *config.Config
	logger           *logrus.Logger
	db               *gorm.DB
	portfolioService *services.PortfolioService
	authService      *services.AuthService
	alertService     *services.AlertService
	web3Service      *services.Web3Service
}

func NewServer(
	cfg *config.Config,
	logger *logrus.Logger,
	db *gorm.DB,
	portfolioService *services.PortfolioService,
	authService *services.AuthService,
	alertService *services.AlertService,
	web3Service *services.Web3Service,
) *Server {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	s := &Server{
		engine:           r,
		config:           cfg,
		logger:           logger,
		db:               db,
		portfolioService: portfolioService,
		authService:      authService,
		alertService:     alertService,
		web3Service:      web3Service,
	}

	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	// Add middleware
	s.engine.Use(corsMiddleware())
	s.engine.Use(loggerMiddleware(s.logger))
	s.engine.Use(rateLimitMiddleware())
	s.engine.Use(requestIDMiddleware())
	s.engine.Use(securityHeadersMiddleware())
	s.engine.Use(errorHandlingMiddleware())
	s.engine.Use(recoveryMiddleware(s.logger))

	// Public routes
	s.engine.GET("/api/health", s.healthHandler)
	s.engine.GET("/api/version", s.versionHandler)

	// Auth routes
	auth := s.engine.Group("/api/v1/auth")
	{
		auth.POST("/register", s.registerHandler)
		auth.POST("/login", s.loginHandler)
		auth.POST("/refresh", s.refreshTokenHandler)
		auth.POST("/logout", s.logoutHandler)
	}

	// Protected routes
	protected := s.engine.Group("/api/v1")
	protected.Use(authMiddleware(s.authService))
	{
		// User management
		protected.GET("/user/profile", s.getUserProfileHandler)
		protected.PUT("/user/profile", s.updateUserProfileHandler)
		protected.DELETE("/user/account", s.deleteUserAccountHandler)
		// Subscription management
		protected.GET("/user/subscription", s.getSubscriptionHandler)
		protected.PUT("/user/subscription", s.updateSubscriptionHandler)

		// Portfolio management
		portfolios := protected.Group("/portfolios")
		{
			portfolios.GET("", s.getPortfoliosHandler)
			portfolios.POST("", s.createPortfolioHandler)
			portfolios.GET("/:id", s.getPortfolioHandler)
			portfolios.PUT("/:id", s.updatePortfolioHandler)
			portfolios.DELETE("/:id", s.deletePortfolioHandler)

			// Portfolio addresses
			portfolios.GET("/:id/addresses", s.getPortfolioAddressesHandler)
			portfolios.POST("/:id/addresses", s.addPortfolioAddressHandler)
			portfolios.PUT("/:id/addresses/:addressId", s.updatePortfolioAddressHandler)
			portfolios.DELETE("/:id/addresses/:addressId", s.deletePortfolioAddressHandler)

			// Portfolio balances
			portfolios.GET("/:id/balances", s.getPortfolioBalancesHandler)
			portfolios.GET("/:id/balances/refresh", s.refreshPortfolioBalancesHandler)

			// Portfolio transactions
			portfolios.GET("/:id/transactions", s.getPortfolioTransactionsHandler)
			portfolios.GET("/:id/transactions/refresh", s.refreshPortfolioTransactionsHandler)
		}

		// Analytics
		analytics := protected.Group("/analytics")
		analytics.Use(planMiddleware(s.db))
		{
			analytics.GET("/portfolio/:id/summary", s.getPortfolioSummaryHandler)
			analytics.GET("/portfolio/:id/performance", s.getPortfolioPerformanceHandler)
			analytics.GET("/portfolio/:id/allocation", s.getPortfolioAllocationHandler)
			analytics.GET("/portfolio/:id/history", s.getPortfolioHistoryHandler)
		}

		// Alerts
		alerts := protected.Group("/alerts")
		alerts.Use(planMiddleware(s.db))
		{
			alerts.GET("", s.getAlertsHandler)
			alerts.POST("", s.createAlertHandler)
			alerts.GET("/:id", s.getAlertHandler)
			alerts.PUT("/:id", s.updateAlertHandler)
			alerts.DELETE("/:id", s.deleteAlertHandler)
			alerts.POST("/:id/toggle", s.toggleAlertHandler)
		}

		// Web3 data
		web3 := protected.Group("/web3")
		{
			web3.GET("/networks", s.getNetworksHandler)
			web3.GET("/networks/:network/status", s.getNetworkStatusHandler)
			web3.GET("/networks/:network/gas", s.getGasPriceHandler)
			web3.GET("/tokens/:symbol/price", s.getTokenPriceHandler)
			web3.GET("/addresses/:address/balance", s.getAddressBalanceHandler)
			web3.GET("/addresses/:address/tokens", s.getAddressTokensHandler)
		}

		// Forum routes
		forum := protected.Group("/forum")
		{
			questions := forum.Group("/questions")
			{
				questions.POST("", s.createQuestionHandler)
				questions.GET("", s.listQuestionsHandler)
				questions.GET(":id", s.getQuestionHandler)
				questions.PUT(":id", s.updateQuestionHandler)
				questions.DELETE(":id", s.deleteQuestionHandler)
			}
		}
	}
}

func (s *Server) Start(addr string) error {
	return s.engine.Run(addr)
}

// Health and version handlers
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"timestamp": gin.H{
			"unix": time.Now().Unix(),
			"iso":  time.Now().Format(time.RFC3339),
		},
		"services": gin.H{
			"database": "ok",
			"web3":     s.web3Service.GetNetworkStatus(),
		},
	})
}

func (s *Server) versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":     "1.0.0",
		"build":       os.Getenv("BUILD_VERSION"),
		"environment": s.config.Environment,
	})
}
