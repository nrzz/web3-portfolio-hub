package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/models"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// Register creates a new user account
func (s *AuthService) Register(email, password, discordID string) (*models.User, string, error) {
	// Check if user already exists
	var existingUser models.User
	err := s.db.Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		return nil, "", fmt.Errorf("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		Email:              email,
		Password:           string(hashedPassword),
		DiscordID:          discordID,
		IsActive:           true,
		SubscriptionTier:   "basic",
		SubscriptionStatus: "active",
	}

	err = s.db.Create(user).Error
	if err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	// Generate JWT token
	token, err := s.generateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, "", fmt.Errorf("account is deactivated")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return &user, token, nil
}

// RefreshToken generates a new token from an existing one
func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Check if user still exists and is active
	var user models.User
	err = s.db.Where("id = ? AND is_active = ?", claims.UserID, true).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("user not found or inactive")
	}

	// Generate new token
	newToken, err := s.generateToken(user.ID.String(), user.Email)
	if err != nil {
		return "", fmt.Errorf("failed to generate new token: %w", err)
	}

	return newToken, nil
}

// ValidateToken validates a JWT token and returns the user ID
func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Check if user still exists and is active
	var user models.User
	err = s.db.Where("id = ? AND is_active = ?", claims.UserID, true).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("user not found or inactive")
	}

	return claims.UserID, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(userID string) (*models.User, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var user models.User
	err = s.db.Where("id = ?", userUUID).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &user, nil
}

// UpdateUser updates user information
func (s *AuthService) UpdateUser(userID, email, discordID string) (*models.User, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Check if email is already taken by another user
	if email != "" && email != user.Email {
		var existingUser models.User
		err := s.db.Where("email = ? AND id != ?", email, user.ID).First(&existingUser).Error
		if err == nil {
			return nil, fmt.Errorf("email already taken")
		}
		user.Email = email
	}

	if discordID != "" {
		user.DiscordID = discordID
	}

	err = s.db.Save(user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// DeleteUser deactivates a user account
func (s *AuthService) DeleteUser(userID string) error {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.IsActive = false
	err = s.db.Save(user).Error
	if err != nil {
		return fmt.Errorf("failed to deactivate user: %w", err)
	}

	return nil
}

// ChangePassword changes a user's password
func (s *AuthService) ChangePassword(userID, currentPassword, newPassword string) error {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Verify current password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
	if err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	err = s.db.Save(user).Error
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

// ResetPassword initiates a password reset
func (s *AuthService) ResetPassword(email string) (string, error) {
	var user models.User
	err := s.db.Where("email = ? AND is_active = ?", email, true).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	// Generate reset token
	resetToken := generateResetToken()

	// In a real implementation, you'd store this token in the database with an expiration
	// and send it via email. For now, we'll just return it.

	return resetToken, nil
}

// ConfirmPasswordReset confirms a password reset
func (s *AuthService) ConfirmPasswordReset(email, resetToken, newPassword string) error {
	var user models.User
	err := s.db.Where("email = ? AND is_active = ?", email, true).First(&user).Error
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// In a real implementation, you'd validate the reset token
	// For now, we'll just proceed with the password change

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	err = s.db.Save(&user).Error
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

// generateToken creates a new JWT token
func (s *AuthService) generateToken(userID, email string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "web3-portfolio-dashboard",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

// generateResetToken creates a random reset token
func generateResetToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// UpdateSubscription updates a user's subscription tier and status
func (s *AuthService) UpdateSubscription(userID, tier, status string) (*models.User, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if tier != "" {
		user.SubscriptionTier = tier
	}
	if status != "" {
		user.SubscriptionStatus = status
	}

	err = s.db.Save(user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}

	return user, nil
}
