package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"web3-portfolio-dashboard/backend/internal/config"
	"web3-portfolio-dashboard/backend/internal/database"
	"web3-portfolio-dashboard/backend/internal/services"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newMinimalTestServer() *Server {
	cfg := &config.Config{
		JWTSecret:          "test-secret",
		Environment:        "test",
		CorsAllowedOrigins: []string{"http://localhost:3000"},
	}
	logger := logrus.New()
	web3Service := services.NewWeb3Service(cfg)

	s := &Server{
		engine:      gin.New(),
		config:      cfg,
		logger:      logger,
		web3Service: web3Service,
	}
	s.engine.GET("/health", s.healthHandler)
	s.engine.GET("/api/health", s.healthHandler)
	s.engine.GET("/api/version", s.versionHandler)
	return s
}

func setupTestServer(t *testing.T) *Server {
	t.Helper()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
	}

	db, err := database.New(dbURL)
	if err != nil {
		t.Skipf("PostgreSQL not available for integration tests: %v", err)
	}
	require.NoError(t, database.Migrate(db))

	cfg := &config.Config{
		JWTSecret:          "test-secret",
		Environment:        "test",
		CorsAllowedOrigins: []string{"http://localhost:3000"},
	}
	logger := logrus.New()
	web3Service := services.NewWeb3Service(cfg)
	portfolioService := services.NewPortfolioService(db, web3Service)
	authService := services.NewAuthService(db, cfg.JWTSecret)
	alertService := services.NewAlertService(db)

	return NewServer(cfg, logger, db, portfolioService, authService, alertService, web3Service)
}

func TestHealthHandler(t *testing.T) {
	server := newMinimalTestServer()

	for _, path := range []string{"/health", "/api/health"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()
			server.engine.ServeHTTP(rec, req)

			require.Equal(t, http.StatusOK, rec.Code)

			var body map[string]interface{}
			require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
			require.Equal(t, "ok", body["status"])
		})
	}
}

func TestVersionHandler(t *testing.T) {
	server := newMinimalTestServer()

	req := httptest.NewRequest(http.MethodGet, "/api/version", nil)
	rec := httptest.NewRecorder()
	server.engine.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
	require.Equal(t, "1.0.0", body["version"])
}

func TestAuthRegisterAndProfile(t *testing.T) {
	server := setupTestServer(t)

	registerBody := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	payload, err := json.Marshal(registerBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	server.engine.ServeHTTP(rec, req)
	require.Equal(t, http.StatusCreated, rec.Code)

	var registerResp map[string]interface{}
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &registerResp))
	token, ok := registerResp["token"].(string)
	require.True(t, ok)
	require.NotEmpty(t, token)

	profileReq := httptest.NewRequest(http.MethodGet, "/api/v1/user/profile", nil)
	profileReq.Header.Set("Authorization", "Bearer "+token)
	profileRec := httptest.NewRecorder()
	server.engine.ServeHTTP(profileRec, profileReq)

	require.Equal(t, http.StatusOK, profileRec.Code)

	var profileResp map[string]interface{}
	require.NoError(t, json.Unmarshal(profileRec.Body.Bytes(), &profileResp))
	user, ok := profileResp["user"].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "test@example.com", user["email"])
}

func TestLoginInvalidCredentials(t *testing.T) {
	server := setupTestServer(t)

	loginBody := map[string]string{
		"email":    "missing@example.com",
		"password": "wrongpassword",
	}
	payload, err := json.Marshal(loginBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	server.engine.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
