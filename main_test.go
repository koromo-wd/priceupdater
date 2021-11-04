package main

import (
	"testing"
	"time"

	"github.com/koromo-wd/priceupdater/oracle"
	"github.com/stretchr/testify/assert"
)

func TestCreateTradingPairs(t *testing.T) {
	quoteItems := []oracle.QuoteItem{
		{
			Symbol:      "A",
			Name:        "A Token",
			Slug:        "a",
			LastUpdated: time.UnixMilli(1),
			USDPrice:    1,
		},
		{
			Symbol:      "B",
			Name:        "B Token",
			Slug:        "b",
			LastUpdated: time.UnixMilli(3),
			USDPrice:    0.8,
		},
	}

	pairs := createTradingPairs(quoteItems)

	for i, pair := range pairs {
		quoteItem := quoteItems[i]
		assert.Equal(t, quoteItem.Symbol, pair.BaseSymbol)
		assert.Equal(t, usd, pair.QuoteSymbol)
		assert.Equal(t, quoteItem.USDPrice, pair.Price)
		assert.Equal(t, quoteItem.LastUpdated, pair.UpdatedTime)
	}
}
