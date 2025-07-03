package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/models"
)

type AlertService struct {
	db *gorm.DB
}

type AlertCondition struct {
	Type      string      `json:"type"`      // price, balance, transaction
	Operator  string      `json:"operator"`  // >, <, >=, <=, ==, !=
	Value     interface{} `json:"value"`
	Token     string      `json:"token,omitempty"`
	Network   string      `json:"network,omitempty"`
	Address   string      `json:"address,omitempty"`
}

type AlertNotification struct {
	AlertID   string                 `json:"alert_id"`
	Type      string                 `json:"type"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
}

func NewAlertService(db *gorm.DB) *AlertService {
	return &AlertService{db: db}
}

// GetAlerts retrieves all alerts for a user
func (s *AlertService) GetAlerts(userID string) ([]models.Alert, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var alerts []models.Alert
	err = s.db.Where("user_id = ?", userUUID).Order("created_at DESC").Find(&alerts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get alerts: %w", err)
	}

	return alerts, nil
}

// CreateAlert creates a new alert
func (s *AlertService) CreateAlert(userID, alertType, name string, conditions map[string]interface{}) (*models.Alert, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Validate alert type
	if !isValidAlertType(alertType) {
		return nil, fmt.Errorf("invalid alert type: %s", alertType)
	}

	// Validate conditions
	if err := s.validateConditions(conditions); err != nil {
		return nil, fmt.Errorf("invalid conditions: %w", err)
	}

	// Serialize conditions to JSON
	conditionsJSON, err := json.Marshal(conditions)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize conditions: %w", err)
	}

	alert := &models.Alert{
		UserID:    userUUID,
		Type:      alertType,
		Name:      name,
		Conditions: string(conditionsJSON),
		IsActive:  true,
	}

	err = s.db.Create(alert).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create alert: %w", err)
	}

	return alert, nil
}

// GetAlert retrieves a specific alert
func (s *AlertService) GetAlert(userID, alertID string) (*models.Alert, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	alertUUID, err := uuid.Parse(alertID)
	if err != nil {
		return nil, fmt.Errorf("invalid alert ID: %w", err)
	}

	var alert models.Alert
	err = s.db.Where("id = ? AND user_id = ?", alertUUID, userUUID).First(&alert).Error
	if err != nil {
		return nil, fmt.Errorf("alert not found: %w", err)
	}

	return &alert, nil
}

// UpdateAlert updates an alert
func (s *AlertService) UpdateAlert(userID, alertID, alertType, name string, conditions map[string]interface{}) (*models.Alert, error) {
	alert, err := s.GetAlert(userID, alertID)
	if err != nil {
		return nil, err
	}

	if alertType != "" {
		if !isValidAlertType(alertType) {
			return nil, fmt.Errorf("invalid alert type: %s", alertType)
		}
		alert.Type = alertType
	}

	if name != "" {
		alert.Name = name
	}

	if conditions != nil {
		if err := s.validateConditions(conditions); err != nil {
			return nil, fmt.Errorf("invalid conditions: %w", err)
		}

		conditionsJSON, err := json.Marshal(conditions)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize conditions: %w", err)
		}
		alert.Conditions = string(conditionsJSON)
	}

	err = s.db.Save(alert).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update alert: %w", err)
	}

	return alert, nil
}

// DeleteAlert deletes an alert
func (s *AlertService) DeleteAlert(userID, alertID string) error {
	alert, err := s.GetAlert(userID, alertID)
	if err != nil {
		return err
	}

	err = s.db.Delete(alert).Error
	if err != nil {
		return fmt.Errorf("failed to delete alert: %w", err)
	}

	return nil
}

// ToggleAlert toggles the active status of an alert
func (s *AlertService) ToggleAlert(userID, alertID string) (*models.Alert, error) {
	alert, err := s.GetAlert(userID, alertID)
	if err != nil {
		return nil, err
	}

	alert.IsActive = !alert.IsActive
	err = s.db.Save(alert).Error
	if err != nil {
		return nil, fmt.Errorf("failed to toggle alert: %w", err)
	}

	return alert, nil
}

// CheckAlerts checks all active alerts and triggers notifications if conditions are met
func (s *AlertService) CheckAlerts() error {
	var alerts []models.Alert
	err := s.db.Where("is_active = ?", true).Find(&alerts).Error
	if err != nil {
		return fmt.Errorf("failed to get active alerts: %w", err)
	}

	for _, alert := range alerts {
		if err := s.checkAlert(&alert); err != nil {
			// Log error but continue with other alerts
			fmt.Printf("Error checking alert %s: %v\n", alert.ID, err)
		}
	}

	return nil
}

// checkAlert checks if a specific alert should be triggered
func (s *AlertService) checkAlert(alert *models.Alert) error {
	var conditions map[string]interface{}
	err := json.Unmarshal([]byte(alert.Conditions), &conditions)
	if err != nil {
		return fmt.Errorf("failed to parse conditions: %w", err)
	}

	// Check conditions based on alert type
	switch alert.Type {
	case "price":
		return s.checkPriceAlert(alert, conditions)
	case "balance":
		return s.checkBalanceAlert(alert, conditions)
	case "transaction":
		return s.checkTransactionAlert(alert, conditions)
	default:
		return fmt.Errorf("unknown alert type: %s", alert.Type)
	}
}

// checkPriceAlert checks price-based alerts
func (s *AlertService) checkPriceAlert(alert *models.Alert, conditions map[string]interface{}) error {
	// This is a simplified implementation
	// In production, you'd fetch real-time price data from APIs
	
	token, ok := conditions["token"].(string)
	if !ok {
		return fmt.Errorf("token not specified in conditions")
	}

	operator, ok := conditions["operator"].(string)
	if !ok {
		return fmt.Errorf("operator not specified in conditions")
	}

	targetValue, ok := conditions["value"].(float64)
	if !ok {
		return fmt.Errorf("value not specified in conditions")
	}

	// Mock current price (in production, fetch from API)
	currentPrice := 1000.0 // Mock price for demonstration

	// Check condition
	shouldTrigger := false
	switch operator {
	case ">":
		shouldTrigger = currentPrice > targetValue
	case "<":
		shouldTrigger = currentPrice < targetValue
	case ">=":
		shouldTrigger = currentPrice >= targetValue
	case "<=":
		shouldTrigger = currentPrice <= targetValue
	case "==":
		shouldTrigger = currentPrice == targetValue
	case "!=":
		shouldTrigger = currentPrice != targetValue
	}

	if shouldTrigger {
		return s.triggerAlert(alert, map[string]interface{}{
			"token":         token,
			"current_price": currentPrice,
			"target_price":  targetValue,
			"operator":      operator,
		})
	}

	return nil
}

// checkBalanceAlert checks balance-based alerts
func (s *AlertService) checkBalanceAlert(alert *models.Alert, conditions map[string]interface{}) error {
	address, ok := conditions["address"].(string)
	if !ok {
		return fmt.Errorf("address not specified in conditions")
	}

	network, ok := conditions["network"].(string)
	if !ok {
		return fmt.Errorf("network not specified in conditions")
	}

	operator, ok := conditions["operator"].(string)
	if !ok {
		return fmt.Errorf("operator not specified in conditions")
	}

	targetValue, ok := conditions["value"].(float64)
	if !ok {
		return fmt.Errorf("value not specified in conditions")
	}

	// Get current balance (in production, use Web3Service)
	currentBalance := 0.5 // Mock balance for demonstration

	// Check condition
	shouldTrigger := false
	switch operator {
	case ">":
		shouldTrigger = currentBalance > targetValue
	case "<":
		shouldTrigger = currentBalance < targetValue
	case ">=":
		shouldTrigger = currentBalance >= targetValue
	case "<=":
		shouldTrigger = currentBalance <= targetValue
	case "==":
		shouldTrigger = currentBalance == targetValue
	case "!=":
		shouldTrigger = currentBalance != targetValue
	}

	if shouldTrigger {
		return s.triggerAlert(alert, map[string]interface{}{
			"address":        address,
			"network":        network,
			"current_balance": currentBalance,
			"target_balance": targetValue,
			"operator":       operator,
		})
	}

	return nil
}

// checkTransactionAlert checks transaction-based alerts
func (s *AlertService) checkTransactionAlert(alert *models.Alert, conditions map[string]interface{}) error {
	// This is a simplified implementation
	// In production, you'd check recent transactions from blockchain APIs
	
	address, ok := conditions["address"].(string)
	if !ok {
		return fmt.Errorf("address not specified in conditions")
	}

	network, ok := conditions["network"].(string)
	if !ok {
		return fmt.Errorf("network not specified in conditions")
	}

	// Mock transaction detection (in production, check actual transactions)
	hasNewTransaction := false // Mock value

	if hasNewTransaction {
		return s.triggerAlert(alert, map[string]interface{}{
			"address": address,
			"network": network,
			"message": "New transaction detected",
		})
	}

	return nil
}

// triggerAlert creates a notification for a triggered alert
func (s *AlertService) triggerAlert(alert *models.Alert, data map[string]interface{}) error {
	notification := &AlertNotification{
		AlertID:   alert.ID.String(),
		Type:      alert.Type,
		Message:   fmt.Sprintf("Alert triggered: %s", alert.Name),
		Data:      data,
		Timestamp: time.Now(),
	}

	// In production, you'd:
	// 1. Save notification to database
	// 2. Send email/SMS/push notification
	// 3. Send to webhook if configured
	// 4. Send to Discord if configured

	fmt.Printf("Alert triggered: %s - %s\n", alert.Name, notification.Message)
	
	return nil
}

// validateConditions validates alert conditions
func (s *AlertService) validateConditions(conditions map[string]interface{}) error {
	// Check required fields based on alert type
	alertType, ok := conditions["type"].(string)
	if !ok {
		return fmt.Errorf("alert type not specified")
	}

	switch alertType {
	case "price":
		return s.validatePriceConditions(conditions)
	case "balance":
		return s.validateBalanceConditions(conditions)
	case "transaction":
		return s.validateTransactionConditions(conditions)
	default:
		return fmt.Errorf("unknown alert type: %s", alertType)
	}
}

// validatePriceConditions validates price alert conditions
func (s *AlertService) validatePriceConditions(conditions map[string]interface{}) error {
	required := []string{"token", "operator", "value"}
	for _, field := range required {
		if _, ok := conditions[field]; !ok {
			return fmt.Errorf("required field missing: %s", field)
		}
	}

	operator, ok := conditions["operator"].(string)
	if !ok {
		return fmt.Errorf("operator must be a string")
	}

	if !isValidOperator(operator) {
		return fmt.Errorf("invalid operator: %s", operator)
	}

	return nil
}

// validateBalanceConditions validates balance alert conditions
func (s *AlertService) validateBalanceConditions(conditions map[string]interface{}) error {
	required := []string{"address", "network", "operator", "value"}
	for _, field := range required {
		if _, ok := conditions[field]; !ok {
			return fmt.Errorf("required field missing: %s", field)
		}
	}

	operator, ok := conditions["operator"].(string)
	if !ok {
		return fmt.Errorf("operator must be a string")
	}

	if !isValidOperator(operator) {
		return fmt.Errorf("invalid operator: %s", operator)
	}

	return nil
}

// validateTransactionConditions validates transaction alert conditions
func (s *AlertService) validateTransactionConditions(conditions map[string]interface{}) error {
	required := []string{"address", "network"}
	for _, field := range required {
		if _, ok := conditions[field]; !ok {
			return fmt.Errorf("required field missing: %s", field)
		}
	}

	return nil
}

// Helper functions
func isValidAlertType(alertType string) bool {
	validTypes := []string{"price", "balance", "transaction"}
	for _, t := range validTypes {
		if t == alertType {
			return true
		}
	}
	return false
}

func isValidOperator(operator string) bool {
	validOperators := []string{">", "<", ">=", "<=", "==", "!="}
	for _, op := range validOperators {
		if op == operator {
			return true
		}
	}
	return false
} 