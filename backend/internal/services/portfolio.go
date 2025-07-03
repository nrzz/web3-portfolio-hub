package services

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/models"
)

type PortfolioService struct {
	db          *gorm.DB
	web3Service *Web3Service
}

type PortfolioSummary struct {
	TotalValue        string            `json:"total_value"`
	TotalChange24h    string            `json:"total_change_24h"`
	TotalChange7d     string            `json:"total_change_7d"`
	TotalChange30d    string            `json:"total_change_30d"`
	AssetCount        int               `json:"asset_count"`
	NetworkCount      int               `json:"network_count"`
	TopAssets         []PortfolioAsset  `json:"top_assets"`
	NetworkAllocation map[string]string `json:"network_allocation"`
}

type PortfolioAsset struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Value     string `json:"value"`
	Change24h string `json:"change_24h"`
	Network   string `json:"network"`
}

type PortfolioPerformance struct {
	Period      string                 `json:"period"`
	Data        []PerformanceDataPoint `json:"data"`
	TotalReturn string                 `json:"total_return"`
	BestDay     string                 `json:"best_day"`
	WorstDay    string                 `json:"worst_day"`
}

type PerformanceDataPoint struct {
	Date   string `json:"date"`
	Value  string `json:"value"`
	Change string `json:"change"`
}

type PortfolioAllocation struct {
	ByNetwork map[string]NetworkAllocation `json:"by_network"`
	ByAsset   map[string]AssetAllocation   `json:"by_asset"`
}

type NetworkAllocation struct {
	Value      string `json:"value"`
	Percentage string `json:"percentage"`
	AssetCount int    `json:"asset_count"`
}

type AssetAllocation struct {
	Value      string `json:"value"`
	Percentage string `json:"percentage"`
	Amount     string `json:"amount"`
	Network    string `json:"network"`
}

type PortfolioHistory struct {
	Period string             `json:"period"`
	Data   []HistoryDataPoint `json:"data"`
}

type HistoryDataPoint struct {
	Date   string `json:"date"`
	Value  string `json:"value"`
	Change string `json:"change"`
	Volume string `json:"volume"`
}

func NewPortfolioService(db *gorm.DB, web3 *Web3Service) *PortfolioService {
	return &PortfolioService{
		db:          db,
		web3Service: web3,
	}
}

// GetPortfolios retrieves all portfolios for a user
func (s *PortfolioService) GetPortfolios(userID string) ([]models.Portfolio, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var portfolios []models.Portfolio
	err = s.db.Where("user_id = ?", userUUID).Find(&portfolios).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get portfolios: %w", err)
	}

	// Manually load addresses for each portfolio
	for i := range portfolios {
		var addresses []models.Address
		err = s.db.Where("portfolio_id = ?", portfolios[i].ID).Find(&addresses).Error
		if err != nil {
			return nil, fmt.Errorf("failed to get addresses for portfolio %s: %w", portfolios[i].ID, err)
		}
		portfolios[i].Addresses = addresses
	}

	return portfolios, nil
}

// CreatePortfolio creates a new portfolio for a user
func (s *PortfolioService) CreatePortfolio(userID, name string) (*models.Portfolio, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	portfolio := &models.Portfolio{
		UserID: userUUID,
		Name:   name,
	}

	err = s.db.Create(portfolio).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio: %w", err)
	}

	return portfolio, nil
}

// GetPortfolio retrieves a specific portfolio
func (s *PortfolioService) GetPortfolio(userID, portfolioID string) (*models.Portfolio, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	portfolioUUID, err := uuid.Parse(portfolioID)
	if err != nil {
		return nil, fmt.Errorf("invalid portfolio ID: %w", err)
	}

	var portfolio models.Portfolio
	err = s.db.Where("id = ? AND user_id = ?", portfolioUUID, userUUID).First(&portfolio).Error
	if err != nil {
		return nil, fmt.Errorf("portfolio not found: %w", err)
	}

	// Manually load addresses for the portfolio
	var addresses []models.Address
	err = s.db.Where("portfolio_id = ?", portfolio.ID).Find(&addresses).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get addresses for portfolio: %w", err)
	}
	portfolio.Addresses = addresses

	return &portfolio, nil
}

// UpdatePortfolio updates a portfolio
func (s *PortfolioService) UpdatePortfolio(userID, portfolioID, name string) (*models.Portfolio, error) {
	portfolio, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	portfolio.Name = name
	err = s.db.Save(portfolio).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update portfolio: %w", err)
	}

	return portfolio, nil
}

// DeletePortfolio deletes a portfolio
func (s *PortfolioService) DeletePortfolio(userID, portfolioID string) error {
	portfolio, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return err
	}

	err = s.db.Delete(portfolio).Error
	if err != nil {
		return fmt.Errorf("failed to delete portfolio: %w", err)
	}

	return nil
}

// GetPortfolioAddresses retrieves all addresses for a portfolio
func (s *PortfolioService) GetPortfolioAddresses(userID, portfolioID string) ([]models.Address, error) {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	portfolioUUID, _ := uuid.Parse(portfolioID)
	var addresses []models.Address
	err = s.db.Where("portfolio_id = ?", portfolioUUID).Find(&addresses).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get addresses: %w", err)
	}

	return addresses, nil
}

// AddAddress adds a new address to a portfolio
func (s *PortfolioService) AddAddress(userID, portfolioID, address, network, label string) (*models.Address, error) {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	portfolioUUID, _ := uuid.Parse(portfolioID)
	newAddress := &models.Address{
		PortfolioID: portfolioUUID,
		Address:     address,
		Network:     network,
		Label:       label,
	}

	err = s.db.Create(newAddress).Error
	if err != nil {
		return nil, fmt.Errorf("failed to add address: %w", err)
	}

	return newAddress, nil
}

// UpdateAddress updates an address
func (s *PortfolioService) UpdateAddress(userID, portfolioID, addressID, label string) (*models.Address, error) {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	addressUUID, err := uuid.Parse(addressID)
	if err != nil {
		return nil, fmt.Errorf("invalid address ID: %w", err)
	}

	var address models.Address
	err = s.db.Where("id = ?", addressUUID).First(&address).Error
	if err != nil {
		return nil, fmt.Errorf("address not found: %w", err)
	}

	address.Label = label
	err = s.db.Save(&address).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update address: %w", err)
	}

	return &address, nil
}

// DeleteAddress deletes an address
func (s *PortfolioService) DeleteAddress(userID, portfolioID, addressID string) error {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return err
	}

	addressUUID, err := uuid.Parse(addressID)
	if err != nil {
		return fmt.Errorf("invalid address ID: %w", err)
	}

	err = s.db.Where("id = ?", addressUUID).Delete(&models.Address{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}

	return nil
}

// GetPortfolioBalances retrieves current balances for a portfolio
func (s *PortfolioService) GetPortfolioBalances(userID, portfolioID string) ([]models.Balance, error) {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	portfolioUUID, _ := uuid.Parse(portfolioID)
	var balances []models.Balance
	err = s.db.Preload("Address").Where("address_id IN (SELECT id FROM addresses WHERE portfolio_id = ?)", portfolioUUID).Find(&balances).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get balances: %w", err)
	}

	return balances, nil
}

// RefreshPortfolioBalances updates balances for all addresses in a portfolio
func (s *PortfolioService) RefreshPortfolioBalances(userID, portfolioID string) ([]models.Balance, error) {
	addresses, err := s.GetPortfolioAddresses(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	var allBalances []models.Balance

	for _, address := range addresses {
		// Get native token balance
		balance, err := s.web3Service.GetBalance(address.Address, address.Network)
		if err != nil {
			continue // Skip failed addresses
		}

		// Get token balances
		tokenBalances, err := s.web3Service.GetTokenBalances(address.Address, address.Network)
		if err != nil {
			continue
		}

		// Save native token balance
		if balance.Cmp(big.NewInt(0)) > 0 {
			balanceRecord := models.Balance{
				AddressID:    address.ID,
				TokenAddress: "", // Native token
				Symbol:       getNativeTokenSymbol(address.Network),
				Name:         getNativeTokenName(address.Network),
				Amount:       balance.String(),
				Decimals:     18,
				UpdatedAt:    time.Now(),
			}
			allBalances = append(allBalances, balanceRecord)
		}

		// Save token balances
		for _, token := range tokenBalances {
			balanceRecord := models.Balance{
				AddressID:    address.ID,
				TokenAddress: token.TokenAddress,
				Symbol:       token.Symbol,
				Name:         token.Name,
				Amount:       token.Amount,
				Decimals:     token.Decimals,
				Price:        token.Price,
				Value:        token.Value,
				UpdatedAt:    time.Now(),
			}
			allBalances = append(allBalances, balanceRecord)
		}
	}

	// Save to database
	for _, balance := range allBalances {
		s.db.Save(&balance)
	}

	return allBalances, nil
}

// GetPortfolioTransactions retrieves transactions for a portfolio
func (s *PortfolioService) GetPortfolioTransactions(userID, portfolioID string, page, limit int) ([]models.Transaction, int64, error) {
	_, err := s.GetPortfolio(userID, portfolioID)
	if err != nil {
		return nil, 0, err
	}

	portfolioUUID, _ := uuid.Parse(portfolioID)
	offset := (page - 1) * limit

	var transactions []models.Transaction
	var total int64

	err = s.db.Model(&models.Transaction{}).Where("portfolio_id = ?", portfolioUUID).Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count transactions: %w", err)
	}

	err = s.db.Where("portfolio_id = ?", portfolioUUID).Order("timestamp DESC").Offset(offset).Limit(limit).Find(&transactions).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get transactions: %w", err)
	}

	return transactions, total, nil
}

// RefreshPortfolioTransactions fetches and stores new transactions
func (s *PortfolioService) RefreshPortfolioTransactions(userID, portfolioID string) ([]models.Transaction, error) {
	addresses, err := s.GetPortfolioAddresses(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	portfolioUUID, _ := uuid.Parse(portfolioID)
	var allTransactions []models.Transaction

	// This is a simplified implementation
	// In production, you'd want to use blockchain APIs to fetch actual transactions
	for _, address := range addresses {
		// Mock transaction data
		transaction := models.Transaction{
			PortfolioID:  portfolioUUID,
			TxHash:       "0x" + generateRandomHash(),
			Network:      address.Network,
			TokenAddress: "",
			Amount:       "0.1",
			BlockNumber:  12345678,
			Timestamp:    time.Now(),
		}
		allTransactions = append(allTransactions, transaction)
	}

	// Save to database
	for _, transaction := range allTransactions {
		s.db.Save(&transaction)
	}

	return allTransactions, nil
}

// GetPortfolioSummary gets a summary of portfolio performance
func (s *PortfolioService) GetPortfolioSummary(userID, portfolioID string) (*PortfolioSummary, error) {
	balances, err := s.GetPortfolioBalances(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	summary := &PortfolioSummary{
		TotalValue:        "0",
		TotalChange24h:    "0",
		TotalChange7d:     "0",
		TotalChange30d:    "0",
		AssetCount:        len(balances),
		NetworkCount:      0,
		TopAssets:         []PortfolioAsset{},
		NetworkAllocation: make(map[string]string),
	}

	// Calculate totals and allocations
	networkMap := make(map[string]bool)
	totalValue := big.NewFloat(0)

	for _, balance := range balances {
		if balance.Value != "" {
			if value, err := strconv.ParseFloat(balance.Value, 64); err == nil {
				totalValue.Add(totalValue, big.NewFloat(value))
			}
		}
		networkMap[balance.Address.Network] = true
	}

	summary.TotalValue = totalValue.String()
	summary.NetworkCount = len(networkMap)

	return summary, nil
}

// GetPortfolioPerformance gets performance data for a portfolio
func (s *PortfolioService) GetPortfolioPerformance(userID, portfolioID, period string) (*PortfolioPerformance, error) {
	// This is a simplified implementation
	// In production, you'd calculate actual performance from historical data
	performance := &PortfolioPerformance{
		Period:      period,
		TotalReturn: "0",
		BestDay:     "0",
		WorstDay:    "0",
		Data:        []PerformanceDataPoint{},
	}

	// Generate mock performance data
	for i := 0; i < 30; i++ {
		date := time.Now().AddDate(0, 0, -i)
		dataPoint := PerformanceDataPoint{
			Date:   date.Format("2006-01-02"),
			Value:  "1000.00",
			Change: "0.00",
		}
		performance.Data = append(performance.Data, dataPoint)
	}

	return performance, nil
}

// GetPortfolioAllocation gets asset allocation data
func (s *PortfolioService) GetPortfolioAllocation(userID, portfolioID string) (*PortfolioAllocation, error) {
	balances, err := s.GetPortfolioBalances(userID, portfolioID)
	if err != nil {
		return nil, err
	}

	allocation := &PortfolioAllocation{
		ByNetwork: make(map[string]NetworkAllocation),
		ByAsset:   make(map[string]AssetAllocation),
	}

	// Calculate allocations
	for _, balance := range balances {
		// Network allocation
		if _, exists := allocation.ByNetwork[balance.Address.Network]; !exists {
			allocation.ByNetwork[balance.Address.Network] = NetworkAllocation{
				Value:      "0",
				Percentage: "0",
				AssetCount: 0,
			}
		}

		// Asset allocation
		if balance.Value != "" {
			allocation.ByAsset[balance.Symbol] = AssetAllocation{
				Value:      balance.Value,
				Percentage: "0", // Calculate percentage
				Amount:     balance.Amount,
				Network:    balance.Address.Network,
			}
		}
	}

	return allocation, nil
}

// GetPortfolioHistory gets historical data for a portfolio
func (s *PortfolioService) GetPortfolioHistory(userID, portfolioID, period string) (*PortfolioHistory, error) {
	history := &PortfolioHistory{
		Period: period,
		Data:   []HistoryDataPoint{},
	}

	// Generate mock historical data
	for i := 0; i < 30; i++ {
		date := time.Now().AddDate(0, 0, -i)
		dataPoint := HistoryDataPoint{
			Date:   date.Format("2006-01-02"),
			Value:  "1000.00",
			Change: "0.00",
			Volume: "100.00",
		}
		history.Data = append(history.Data, dataPoint)
	}

	return history, nil
}

// Helper functions
func getNativeTokenSymbol(network string) string {
	switch network {
	case "ethereum":
		return "ETH"
	case "polygon":
		return "MATIC"
	case "bsc":
		return "BNB"
	case "arbitrum":
		return "ETH"
	default:
		return "NATIVE"
	}
}

func getNativeTokenName(network string) string {
	switch network {
	case "ethereum":
		return "Ethereum"
	case "polygon":
		return "Polygon"
	case "bsc":
		return "Binance Smart Chain"
	case "arbitrum":
		return "Arbitrum"
	default:
		return "Native Token"
	}
}

func generateRandomHash() string {
	// Simplified random hash generation
	return fmt.Sprintf("%064x", time.Now().UnixNano())
}
