package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/config"
)

type Web3Service struct {
	clients map[string]*ethclient.Client
	db      *gorm.DB
	config  *config.Config
}

type BalanceInfo struct {
	Address       string         `json:"address"`
	Network       string         `json:"network"`
	Balance       string         `json:"balance"`
	TokenBalances []TokenBalance `json:"token_balances"`
}

type TokenBalance struct {
	TokenAddress string `json:"token_address"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Amount       string `json:"amount"`
	Decimals     uint8  `json:"decimals"`
	Price        string `json:"price"`
	Value        string `json:"value"`
}

type TransactionInfo struct {
	TxHash       string    `json:"tx_hash"`
	Network      string    `json:"network"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	Value        string    `json:"value"`
	TokenAddress string    `json:"token_address"`
	BlockNumber  uint64    `json:"block_number"`
	Timestamp    time.Time `json:"timestamp"`
}

// Add a map of known token decimals
var tokenDecimals = map[string]uint8{
	// Ethereum
	"0xdAC17F958D2ee523a2206206994597C13D831ec7": 6,  // USDT
	"0xA0b86a33E6441b8C4C8C8C8C8C8C8C8C8C8C8C8C": 6,  // USDC (demo address)
	"0x6B175474E89094C44Da98b954EedeAC495271d0F": 18, // DAI
	"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": 18, // WETH
	// Polygon
	"0xc2132D05D31c914a87C6611C10748AEb04B58e8F": 6,  // USDT
	"0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174": 6,  // USDC
	"0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063": 18, // DAI
	"0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270": 18, // WMATIC
	// BSC
	"0x55d398326f99059fF775485246999027B3197955": 18, // USDT
	"0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d": 18, // USDC
	"0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3": 18, // DAI
	"0x2170Ed0880ac9A755fd29B2688956BD959F933F8": 18, // WETH
	// Arbitrum (add more as needed)
}

func NewWeb3Service(cfg *config.Config) *Web3Service {
	clients := make(map[string]*ethclient.Client)

	// Initialize Ethereum client
	if cfg.EthereumRPCURL != "" {
		log.Printf("Connecting to Ethereum RPC: %s", cfg.EthereumRPCURL)
		if client, err := ethclient.Dial(cfg.EthereumRPCURL); err == nil {
			// Test the connection
			if _, err := client.BlockNumber(context.Background()); err == nil {
				clients["ethereum"] = client
				log.Printf("✅ Ethereum client connected successfully")
			} else {
				log.Printf("❌ Ethereum client connection test failed: %v", err)
			}
		} else {
			log.Printf("❌ Failed to connect to Ethereum RPC: %v", err)
		}
	}

	// Initialize Polygon client
	if cfg.PolygonRPCURL != "" {
		log.Printf("Connecting to Polygon RPC: %s", cfg.PolygonRPCURL)
		if client, err := ethclient.Dial(cfg.PolygonRPCURL); err == nil {
			// Test the connection
			if _, err := client.BlockNumber(context.Background()); err == nil {
				clients["polygon"] = client
				log.Printf("✅ Polygon client connected successfully")
			} else {
				log.Printf("❌ Polygon client connection test failed: %v", err)
			}
		} else {
			log.Printf("❌ Failed to connect to Polygon RPC: %v", err)
		}
	}

	// Initialize BSC client
	if cfg.BSCRPCURL != "" {
		log.Printf("Connecting to BSC RPC: %s", cfg.BSCRPCURL)
		if client, err := ethclient.Dial(cfg.BSCRPCURL); err == nil {
			// Test the connection
			if _, err := client.BlockNumber(context.Background()); err == nil {
				clients["bsc"] = client
				log.Printf("✅ BSC client connected successfully")
			} else {
				log.Printf("❌ BSC client connection test failed: %v", err)
			}
		} else {
			log.Printf("❌ Failed to connect to BSC RPC: %v", err)
		}
	}

	// Initialize Arbitrum client
	if cfg.ArbitrumRPCURL != "" {
		log.Printf("Connecting to Arbitrum RPC: %s", cfg.ArbitrumRPCURL)
		if client, err := ethclient.Dial(cfg.ArbitrumRPCURL); err == nil {
			// Test the connection
			if _, err := client.BlockNumber(context.Background()); err == nil {
				clients["arbitrum"] = client
				log.Printf("✅ Arbitrum client connected successfully")
			} else {
				log.Printf("❌ Arbitrum client connection test failed: %v", err)
			}
		} else {
			log.Printf("❌ Failed to connect to Arbitrum RPC: %v", err)
		}
	}

	log.Printf("Web3Service initialized with %d connected networks", len(clients))
	return &Web3Service{
		clients: clients,
		config:  cfg,
	}
}

// GetBalance gets the native token balance for an address
func (s *Web3Service) GetBalance(address, network string) (*big.Int, error) {
	client, exists := s.clients[network]
	if !exists {
		return nil, fmt.Errorf("network %s not supported or not connected", network)
	}

	// Validate address
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid address format: %s", address)
	}

	addr := common.HexToAddress(address)
	log.Printf("Getting balance for address %s on network %s", address, network)

	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Printf("Error getting balance for %s on %s: %v", address, network, err)
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	log.Printf("Balance for %s on %s: %s", address, network, balance.String())
	return balance, nil
}

// GetTokenBalance gets the ERC-20 token balance for an address
func (s *Web3Service) GetTokenBalance(address, tokenAddress, network string) (*big.Int, error) {
	client, exists := s.clients[network]
	if !exists {
		return nil, fmt.Errorf("network %s not supported", network)
	}

	if !common.IsHexAddress(address) || !common.IsHexAddress(tokenAddress) {
		return nil, fmt.Errorf("invalid address format")
	}

	// Use ABI encoding for balanceOf(address)
	erc20ABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]`))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}
	data, err := erc20ABI.Pack("balanceOf", common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("failed to pack ABI: %w", err)
	}

	to := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		To:   &to,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	if len(result) == 0 {
		return big.NewInt(0), nil
	}

	balance := new(big.Int).SetBytes(result)
	return balance, nil
}

// GetPortfolioBalances gets balances for multiple addresses
func (s *Web3Service) GetPortfolioBalances(addresses []string, network string) ([]BalanceInfo, error) {
	var balances []BalanceInfo

	for _, address := range addresses {
		balance, err := s.GetBalance(address, network)
		if err != nil {
			continue // Skip failed addresses
		}

		balanceInfo := BalanceInfo{
			Address: address,
			Network: network,
			Balance: balance.String(),
		}

		// Get token balances if supported
		if network == "ethereum" {
			tokenBalances, err := s.GetTokenBalances(address, network)
			if err == nil {
				balanceInfo.TokenBalances = tokenBalances
			}
		}

		balances = append(balances, balanceInfo)
	}

	return balances, nil
}

// GetTokenBalances gets all token balances for an address
func (s *Web3Service) GetTokenBalances(address, network string) ([]TokenBalance, error) {
	// This is a simplified version - in production, you'd want to:
	// 1. Use a token list API
	// 2. Cache results
	// 3. Handle rate limiting
	// 4. Use batch calls for efficiency

	// Common tokens for each network
	tokens := map[string][]struct {
		Address string
		Symbol  string
		Name    string
	}{
		"ethereum": {
			{"0xdAC17F958D2ee523a2206206994597C13D831ec7", "USDT", "Tether USD"},
			{"0xA0b86a33E6441b8C4C8C8C8C8C8C8C8C8C8C8C8C", "USDC", "USD Coin"},
			{"0x6B175474E89094C44Da98b954EedeAC495271d0F", "DAI", "Dai Stablecoin"},
			{"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", "WETH", "Wrapped Ether"},
		},
		"polygon": {
			{"0xc2132D05D31c914a87C6611C10748AEb04B58e8F", "USDT", "Tether USD"},
			{"0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", "USDC", "USD Coin"},
			{"0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063", "DAI", "Dai Stablecoin"},
			{"0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270", "WMATIC", "Wrapped MATIC"},
		},
		"bsc": {
			{"0x55d398326f99059fF775485246999027B3197955", "USDT", "Tether USD"},
			{"0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", "USDC", "USD Coin"},
			{"0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3", "DAI", "Dai Stablecoin"},
			{"0x2170Ed0880ac9A755fd29B2688956BD959F933F8", "WETH", "Wrapped Ether"},
		},
		"arbitrum": {
			// Add Arbitrum tokens as needed
		},
	}

	networkTokens, exists := tokens[network]
	if !exists {
		return nil, fmt.Errorf("network %s not supported for token balances", network)
	}

	var tokenBalances []TokenBalance

	for _, token := range networkTokens {
		balance, err := s.GetTokenBalance(address, token.Address, network)
		if err != nil || balance.Cmp(big.NewInt(0)) == 0 {
			continue // Skip failed or zero balances
		}

		decimals := uint8(18)
		if d, ok := tokenDecimals[strings.ToLower(token.Address)]; ok {
			decimals = d
		}
		tokenBalance := TokenBalance{
			TokenAddress: token.Address,
			Symbol:       token.Symbol,
			Name:         token.Name,
			Amount:       balance.String(),
			Decimals:     decimals,
		}

		// Get price from cache or API
		price, err := s.GetTokenPrice(token.Symbol)
		if err == nil {
			tokenBalance.Price = price
			amountFloat := new(big.Float).SetInt(balance)
			amountDiv := new(big.Float).Quo(amountFloat, new(big.Float).SetFloat64(float64(1)).SetPrec(uint(decimals*3)).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))
			priceFloat, _ := new(big.Float).SetString(price)
			valueFloat := new(big.Float).Mul(amountDiv, priceFloat)
			tokenBalance.Value = valueFloat.Text('f', 8)
		}

		tokenBalances = append(tokenBalances, tokenBalance)
	}

	return tokenBalances, nil
}

// GetTokenPrice gets the current price of a token
func (s *Web3Service) GetTokenPrice(symbol string) (string, error) {
	// In production, you'd want to:
	// 1. Use CoinGecko API
	// 2. Cache results
	// 3. Handle rate limiting
	// 4. Use WebSocket for real-time updates

	// Mock prices for demo
	prices := map[string]string{
		"USDT":   "1.00",
		"USDC":   "1.00",
		"DAI":    "1.00",
		"WETH":   "2000.00",
		"WMATIC": "0.80",
	}

	price, exists := prices[strings.ToUpper(symbol)]
	if !exists {
		return "0.00", fmt.Errorf("price not available for %s", symbol)
	}

	return price, nil
}

// GetGasPrice gets the current gas price for a network
func (s *Web3Service) GetGasPrice(network string) (*big.Int, error) {
	client, exists := s.clients[network]
	if !exists {
		return nil, fmt.Errorf("network %s not supported", network)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	return gasPrice, nil
}

// GetNetworkStatus returns the status of all networks
func (s *Web3Service) GetNetworkStatus() map[string]bool {
	status := make(map[string]bool)

	for network, client := range s.clients {
		_, err := client.BlockNumber(context.Background())
		status[network] = err == nil
	}

	return status
}
