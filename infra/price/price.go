package price

import (
	"context"
	"strconv"
	"time"

	"github.com/piovani/digital_wallet_2/infra/redis"
)

const (
	DURATION = time.Second * 60

	// BTC
	KEY_BTC_USD = "BTC_USD"
	KEY_BTC_EUR = "BTC_EUR"

	// ETH
	KEY_ETH_USD = "ETH_USD"
	KEY_ETH_EUR = "ETH_EUR"

	// ADA
	KEY_ADA_USD = "ADA_USD"
	KEY_ADA_EUR = "ADA_EUR"

	// XRP
	KEY_XRP_USD = "XRP_USD"
	KEY_XRP_EUR = "XRP_EUR"

	// DOGE
	KEY_DOGE_USD = "DOGE_USD"
	KEY_DOGE_EUR = "DOGE_EUR"
)

type Price struct {
	Ctx      context.Context
	Redis    *redis.Redis
	Cex      *Cex
	Coinbase *Coinbase
}

func NewPrice(ctx context.Context) *Price {
	return &Price{
		Ctx:      ctx,
		Redis:    redis.NewRedis(),
		Cex:      NewCex(),
		Coinbase: NewCoinbase(),
	}
}

// BTC
func (p Price) CollectBtcUsd() {
	value := p.Cex.getValue("BTC", "USD")
	p.Redis.Set(p.Ctx, KEY_BTC_USD, value, DURATION)
}

func (p Price) GetBtcUsd() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_BTC_USD)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

func (p Price) CollectBtcEur() {
	value := p.Cex.getValue("BTC", "EUR")
	p.Redis.Set(p.Ctx, KEY_BTC_EUR, value, DURATION)
}

func (p Price) GetBtcEur() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_BTC_EUR)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

// ETH
func (p Price) CollectEthUsd() {
	value := p.Cex.getValue("ETH", "USD")
	p.Redis.Set(p.Ctx, KEY_ETH_USD, value, DURATION)
}

func (p Price) GetEthUsd() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_ETH_USD)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

func (p Price) CollecEtcEur() {
	value := p.Cex.getValue("ETH", "EUR")
	p.Redis.Set(p.Ctx, KEY_ETH_EUR, value, DURATION)
}

func (p Price) GetEthEur() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_ETH_EUR)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

// ADA
func (p Price) CollectAdaUsd() {
	value := p.Cex.getValue("ADA", "USD")
	p.Redis.Set(p.Ctx, KEY_ADA_USD, value, DURATION)
}

func (p Price) GetAdaUsd() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_ADA_USD)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

func (p Price) CollectAdaEur() {
	value := p.Cex.getValue("ADA", "EUR")
	p.Redis.Set(p.Ctx, KEY_ADA_EUR, value, DURATION)
}

func (p Price) GetAdaEur() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_ADA_EUR)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

//XRP
func (p Price) CollectXrpUsd() {
	value := p.Cex.getValue("XRP", "USD")
	p.Redis.Set(p.Ctx, KEY_XRP_USD, value, DURATION)
}

func (p Price) GetXrpUsd() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_XRP_USD)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

func (p Price) CollectXrpEur() {
	value := p.Cex.getValue("XRP", "EUR")
	p.Redis.Set(p.Ctx, KEY_XRP_EUR, value, DURATION)
}

func (p Price) GetXrpEur() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_XRP_EUR)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

// DOGE
func (p Price) CollectDogeUsd() {
	value := p.Cex.getValue("DOGE", "USD")
	p.Redis.Set(p.Ctx, KEY_DOGE_USD, value, DURATION)
}

func (p Price) GetDogeUsd() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_DOGE_USD)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

func (p Price) CollectDogeEur() {
	value := p.Coinbase.GetDogeEur()
	p.Redis.Set(p.Ctx, KEY_DOGE_EUR, value, DURATION)
}

func (p Price) GetDogeEur() float64 {
	valueString, err := p.Redis.Get(p.Ctx, KEY_DOGE_EUR)
	if err != nil {
		return 1
	}

	return p.stringToFloat(valueString)
}

// ======== AUXILIARES ========
func (p Price) stringToFloat(valueString string) float64 {
	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return 1
	}

	return value
}
