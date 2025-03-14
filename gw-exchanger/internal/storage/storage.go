package storage

type Repository interface {
	GetExchangeRateForCurrency(fromCurrency, toCurrency string) (float32, error)
	GetAllExchangeRates() (map[string]float32, error)

	Start() error
	Stop() error
}
