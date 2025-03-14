package postgres

import (
	"context"
	"fmt"
	"time"
)

func (c *Connector) GetExchangeRateForCurrency(fromCurrency, toCurrency string) (float32, error) {
	// Check the cache first
	if rates, ok := c.cache.GetExchangeRates(); ok {
		c.log.Info().Msg("Read exchange rates from cache")
		rate := rates[toCurrency] / rates[fromCurrency]
		return rate, nil
	}

	rates, err := c.GetAllExchangeRates()
	if err != nil {
		return 0, err
	}

	rate := rates[toCurrency] / rates[fromCurrency]
	return rate, nil
}

func (c *Connector) GetAllExchangeRates() (map[string]float32, error) {
	// Check the cache first
	if rates, ok := c.cache.GetExchangeRates(); ok {
		c.log.Info().Msg("Read exchange rates from cache")
		return rates, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := c.Client.Query(ctx, "SELECT currency_code, exchange_rate_to_usd FROM currencies")
	if err != nil {
		return nil, fmt.Errorf("failed to query currencies: %w", err)
	}
	defer rows.Close()

	rates := make(map[string]float32)

	for rows.Next() {
		var currencyCode string
		var rate float32
		if err := rows.Scan(&currencyCode, &rate); err != nil {
			return nil, fmt.Errorf("failed to scan exchange rate: %w", err)
		}
		rates[currencyCode] = rate
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	if len(rates) == 0 {
		return nil, fmt.Errorf("no exchange rates found")
	}

	// Store exchange rates in the cache
	c.cache.StoreExchangeRates(rates)

	return rates, nil
}
