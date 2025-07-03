package api

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Auth handlers
func (s *Server) registerHandler(c *gin.Context) {
	var req struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
		DiscordID string `json:"discord_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := s.authService.Register(req.Email, req.Password, req.DiscordID)
	if err != nil {
		s.logger.Error("Registration failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

func (s *Server) loginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		s.logger.Error("Login failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (s *Server) refreshTokenHandler(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.authService.RefreshToken(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) logoutHandler(c *gin.Context) {
	// In a real implementation, you might want to blacklist the token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// User handlers
func (s *Server) getUserProfileHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := s.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *Server) updateUserProfileHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		Email     string `json:"email"`
		DiscordID string `json:"discord_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.authService.UpdateUser(userID, req.Email, req.DiscordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *Server) deleteUserAccountHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := s.authService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

// Portfolio handlers
func (s *Server) getPortfoliosHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	portfolios, err := s.portfolioService.GetPortfolios(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolios": portfolios})
}

func (s *Server) createPortfolioHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio, err := s.portfolioService.CreatePortfolio(userID, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"portfolio": portfolio})
}

func (s *Server) getPortfolioHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	portfolio, err := s.portfolioService.GetPortfolio(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolio": portfolio})
}

func (s *Server) updatePortfolioHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio, err := s.portfolioService.UpdatePortfolio(userID, portfolioID, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolio": portfolio})
}

func (s *Server) deletePortfolioHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	if err := s.portfolioService.DeletePortfolio(userID, portfolioID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted successfully"})
}

// Portfolio address handlers
func (s *Server) getPortfolioAddressesHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	addresses, err := s.portfolioService.GetPortfolioAddresses(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

func (s *Server) addPortfolioAddressHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	var req struct {
		Address string `json:"address" binding:"required"`
		Network string `json:"network" binding:"required"`
		Label   string `json:"label"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := s.portfolioService.AddAddress(userID, portfolioID, req.Address, req.Network, req.Label)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"address": address})
}

func (s *Server) updatePortfolioAddressHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")
	addressID := c.Param("addressId")

	var req struct {
		Label string `json:"label"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := s.portfolioService.UpdateAddress(userID, portfolioID, addressID, req.Label)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}

func (s *Server) deletePortfolioAddressHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")
	addressID := c.Param("addressId")

	if err := s.portfolioService.DeleteAddress(userID, portfolioID, addressID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}

// Portfolio balance handlers
func (s *Server) getPortfolioBalancesHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	balances, err := s.portfolioService.GetPortfolioBalances(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balances": balances})
}

func (s *Server) refreshPortfolioBalancesHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	balances, err := s.portfolioService.RefreshPortfolioBalances(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balances": balances})
}

// Portfolio transaction handlers
func (s *Server) getPortfolioTransactionsHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	transactions, total, err := s.portfolioService.GetPortfolioTransactions(userID, portfolioID, page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (s *Server) refreshPortfolioTransactionsHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	transactions, err := s.portfolioService.RefreshPortfolioTransactions(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// Analytics handlers
func (s *Server) getPortfolioSummaryHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	summary, err := s.portfolioService.GetPortfolioSummary(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

func (s *Server) getPortfolioPerformanceHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")
	period := c.DefaultQuery("period", "30d")

	performance, err := s.portfolioService.GetPortfolioPerformance(userID, portfolioID, period)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"performance": performance})
}

func (s *Server) getPortfolioAllocationHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")

	allocation, err := s.portfolioService.GetPortfolioAllocation(userID, portfolioID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"allocation": allocation})
}

func (s *Server) getPortfolioHistoryHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	portfolioID := c.Param("id")
	period := c.DefaultQuery("period", "30d")

	history, err := s.portfolioService.GetPortfolioHistory(userID, portfolioID, period)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": history})
}

// Alert handlers
func (s *Server) getAlertsHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	alerts, err := s.alertService.GetAlerts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alerts": alerts})
}

func (s *Server) createAlertHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		Type       string                 `json:"type" binding:"required"`
		Name       string                 `json:"name" binding:"required"`
		Conditions map[string]interface{} `json:"conditions" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alert, err := s.alertService.CreateAlert(userID, req.Type, req.Name, req.Conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"alert": alert})
}

func (s *Server) getAlertHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	alertID := c.Param("id")

	alert, err := s.alertService.GetAlert(userID, alertID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alert": alert})
}

func (s *Server) updateAlertHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	alertID := c.Param("id")

	var req struct {
		Type       string                 `json:"type"`
		Name       string                 `json:"name"`
		Conditions map[string]interface{} `json:"conditions"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alert, err := s.alertService.UpdateAlert(userID, alertID, req.Type, req.Name, req.Conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alert": alert})
}

func (s *Server) deleteAlertHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	alertID := c.Param("id")

	if err := s.alertService.DeleteAlert(userID, alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}

func (s *Server) toggleAlertHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	alertID := c.Param("id")

	alert, err := s.alertService.ToggleAlert(userID, alertID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alert": alert})
}

// Web3 handlers
func (s *Server) getNetworksHandler(c *gin.Context) {
	networks := s.web3Service.GetNetworkStatus()
	c.JSON(http.StatusOK, gin.H{"networks": networks})
}

func (s *Server) getNetworkStatusHandler(c *gin.Context) {
	network := c.Param("network")
	status := s.web3Service.GetNetworkStatus()

	if networkStatus, exists := status[network]; exists {
		c.JSON(http.StatusOK, gin.H{"network": network, "status": networkStatus})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Network not found"})
	}
}

func (s *Server) getGasPriceHandler(c *gin.Context) {
	network := c.Param("network")

	gasPrice, err := s.web3Service.GetGasPrice(network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"network":        network,
		"gas_price":      gasPrice.String(),
		"gas_price_gwei": new(big.Int).Div(gasPrice, big.NewInt(1000000000)).String(),
	})
}

func (s *Server) getTokenPriceHandler(c *gin.Context) {
	symbol := c.Param("symbol")

	price, err := s.web3Service.GetTokenPrice(symbol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol": symbol,
		"price":  price,
	})
}

func (s *Server) getAddressBalanceHandler(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "ethereum")

	balance, err := s.web3Service.GetBalance(address, network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"network": network,
		"balance": balance.String(),
	})
}

func (s *Server) getAddressTokensHandler(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "ethereum")

	tokens, err := s.web3Service.GetTokenBalances(address, network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"network": network,
		"tokens":  tokens,
	})
}

// Get current user's subscription
func (s *Server) getSubscriptionHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := s.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subscription_tier":   user.SubscriptionTier,
		"subscription_status": user.SubscriptionStatus,
	})
}

// Update current user's subscription
func (s *Server) updateSubscriptionHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req struct {
		Tier   string `json:"tier"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.authService.UpdateSubscription(userID, req.Tier, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subscription_tier":   user.SubscriptionTier,
		"subscription_status": user.SubscriptionStatus,
	})
}

// Forum Question Handlers

// CreateQuestionHandler handles POST /api/v1/forum/questions
func (s *Server) createQuestionHandler(c *gin.Context) {
	var req struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	// TODO: Call service to create question
	c.JSON(http.StatusCreated, gin.H{"message": "Question created (stub)"})
}

// ListQuestionsHandler handles GET /api/v1/forum/questions
func (s *Server) listQuestionsHandler(c *gin.Context) {
	// TODO: Call service to list questions
	c.JSON(http.StatusOK, gin.H{"questions": []string{}})
}

// GetQuestionHandler handles GET /api/v1/forum/questions/:id
func (s *Server) getQuestionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"question": gin.H{"id": c.Param("id")}})
}

// UpdateQuestionHandler handles PUT /api/v1/forum/questions/:id
func (s *Server) updateQuestionHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var req struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Call service to update question
	c.JSON(http.StatusOK, gin.H{"message": "Question updated (stub)"})
}

// DeleteQuestionHandler handles DELETE /api/v1/forum/questions/:id
func (s *Server) deleteQuestionHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	// TODO: Call service to delete question
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted (stub)"})
}
