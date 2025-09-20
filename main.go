package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"gopkg.in/yaml.v2"
)

// Config represents the YAML configuration structure
type Config struct {
	Tokens []Token `yaml:"tokens"`
}

// Token represents a single token configuration
type Token struct {
	Ticker      string  `yaml:"ticker"`
	CA          string  `yaml:"ca"`
	PriceChange float64 `yaml:"priceChange"`
}

// JupiterPriceResponse represents the Jupiter API price response
type JupiterPriceResponse struct {
	PriceChange24h float64 `json:"priceChange24h"`
	USDPrice       float64 `json:"usdPrice"`
}

// TokenPriceTracker tracks token prices and their changes
type TokenPriceTracker struct {
	config     Config
	tokens     map[string]Token
	lastPrices map[string]float64
}

// NewTokenPriceTracker creates a new token price tracker
func NewTokenPriceTracker(configPath string) (*TokenPriceTracker, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	tokensMap := make(map[string]Token)
	for _, token := range config.Tokens {
		tokensMap[token.CA] = token
	}
	return &TokenPriceTracker{
		config:     config,
		lastPrices: make(map[string]float64),
		tokens:     tokensMap,
	}, nil
}

// getTokensPrice fetches the current price of tokens using Jupiter API
func (t *TokenPriceTracker) getTokensPrice(cas []string) (map[string]*JupiterPriceResponse, error) {
	// Jupiter API endpoint for price data
	url := fmt.Sprintf("https://lite-api.jup.ag/price/v3?ids=%s", strings.Join(cas, ","))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price for tokens: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jupiter API returned status %d ", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var priceResp map[string]*JupiterPriceResponse
	if err := json.Unmarshal(body, &priceResp); err != nil {
		return nil, fmt.Errorf("failed to parse price response: %v", err)
	}

	return priceResp, nil
}

// checkPriceChanges checks all configured tokens for price changes
func (t *TokenPriceTracker) checkPriceChanges() {
	var cas []string
	for _, token := range t.config.Tokens {
		cas = append(cas, token.CA)
	}

	prices, err := t.getTokensPrice(cas)
	if err != nil {
		log.Printf("Error fetching token prices: %v", err)
		return
	}
	for ca, price := range prices {
		token, exists := t.tokens[ca]
		if !exists {
			log.Printf("Token %s not found in config", ca)
			continue
		}

		lastPrice, exists := t.lastPrices[token.CA]
		if !exists {
			// First time checking this token, store the price
			t.lastPrices[token.CA] = price.USDPrice
			log.Printf("Initial price for %s: $%.6f", token.Ticker, price.USDPrice)
			continue
		}

		// Calculate price change percentage
		priceChange := (price.USDPrice - lastPrice) / lastPrice
		absChange := math.Abs(priceChange)

		log.Printf("%s: $%.6f -> $%.6f (%.2f%% change)",
			token.Ticker, lastPrice, price.USDPrice, priceChange*100)

		// Check if change exceeds threshold
		if absChange >= token.PriceChange {
			direction := "increased"
			if priceChange < 0 {
				direction = "decreased"
			}

			title := fmt.Sprintf("Price Alert: %s", token.Ticker)
			message := fmt.Sprintf("%s has %s by %.2f%% (from $%.6f to $%.6f)",
				token.Ticker, direction, absChange*100, lastPrice, price.USDPrice)

			// Send desktop notification
			err := beeep.Notify(title, message, "")
			if err != nil {
				log.Printf("Failed to send notification for %s: %v", token.Ticker, err)
			} else {
				log.Printf("Notification sent for %s", token.Ticker)
			}
		}

		// Update last price
		t.lastPrices[token.CA] = price.USDPrice
	}
}

// Start begins monitoring token prices
func (t *TokenPriceTracker) Start(interval time.Duration) {
	log.Printf("Starting token price monitoring with %d tokens...", len(t.config.Tokens))

	// Initial check
	t.checkPriceChanges()

	// Set up ticker for periodic checks
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.checkPriceChanges()
		}
	}
}

func main() {
	log.Println("DegenMon - Token Price Monitor")

	tracker, err := NewTokenPriceTracker("config.yaml")
	if err != nil {
		log.Fatalf("Failed to initialize tracker: %v", err)
	}

	// Check prices every 5 seconds
	checkInterval := 5 * time.Second
	log.Printf("Monitoring %d tokens, checking every %v", len(tracker.config.Tokens), checkInterval)

	tracker.Start(checkInterval)
}
