package coinmarketcap

type Coins map[string]Coin

type Coin struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int     `json:"rank,string"`
	PriceUsd         float64 `json:"price_usd,string"`
	PriceBtc         float64 `json:"price_btc,string"`
	Volume24hUsd     float64 `json:"24h_volume_usd,string"`
	MarketCapUsd     float64 `json:"market_cap_usd,string"`
	AvailableSupply  float64 `json:"available_supply,string"`
	TotalSupply      float64 `json:"total_supply,string"`
	MaxSupply        float64 `json:"max_supply,string"`
	PercentChange1H  float64 `json:"percent_change_1h,string"`
	PercentChange24H float64 `json:"percent_change_24h,string"`
	PercentChange7D  float64 `json:"percent_change_7d,string"`
	LastUpdated      string  `json:"last_updated"`
	PriceJpy         float64 `json:"price_jpy,string"`
	Volume24hJpy     float64 `json:"24h_volume_jpy,string"`
	MarketCapJpy     float64 `json:"market_cap_jpy,string"`
}
