package worker

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CurrencyExchangeRateUpdater,ConfigGetter,MessageFetcher,MessageProcessor
