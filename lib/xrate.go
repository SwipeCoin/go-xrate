package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
)

func NewBTCCrawler(fiatCurrency currency.Currency, exs ...exchanges.Exchange) (exchanges.Crawler, error) {

	crawler := exchanges.Crawler{
		FiatCurrency:   fiatCurrency,
		CryptoCurrency: currency.Bitcoin(),
		Exchanges:      exs,
	}

	for _, e := range crawler.Exchanges {

		if !e.SupportsCryptoCurrency(crawler.CryptoCurrency) {
			return exchanges.Crawler{}, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.CryptoCurrency.Name)
		}

		if !e.SupportsFiatCurrency(crawler.FiatCurrency) {
			return exchanges.Crawler{}, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.FiatCurrency.Name)
		}
	}

	return crawler, nil
}
