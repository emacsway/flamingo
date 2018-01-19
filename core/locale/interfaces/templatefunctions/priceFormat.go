package templatefunctions

import (
	"strconv"

	"github.com/leekchan/accounting"
	"go.aoe.com/flamingo/core/locale/application"
	"go.aoe.com/flamingo/core/pugtemplate/pugjs"
	"go.aoe.com/flamingo/framework/config"
)

type (
	// PriceFormatFunc for formatting prices
	PriceFormatFunc struct {
		Config             config.Map                     `inject:"config:locale.accounting"`
		TranslationService application.TranslationService `inject:""`
	}
)

// Name alias for use in template
func (pff PriceFormatFunc) Name() string {
	return "priceFormat"
}

// Func as implementation of debug method
func (pff PriceFormatFunc) Func() interface{} {
	return func(value interface{}, currency string) string {
		currency = pff.TranslationService.Translate(currency, "", "", 1, nil)
		ac := accounting.Accounting{
			Symbol:    currency,
			Precision: 2,
		}
		decimal, ok := pff.Config["decimal"].(string)
		if ok {
			ac.Decimal = decimal
		}
		thousand, ok := pff.Config["thousand"].(string)
		if ok {
			ac.Thousand = thousand
		}
		formatZero, ok := pff.Config["formatZero"].(string)
		if ok {
			ac.FormatZero = formatZero
		}
		format, ok := pff.Config["format"].(string)
		if ok {
			ac.Format = format
		}
		if valueNumber, ok := value.(pugjs.Number); ok {
			return ac.FormatMoney(float64(valueNumber))
		} else if valueString, ok := value.(pugjs.String); ok {
			float, err := strconv.ParseFloat(string(valueString), 64)
			if err != nil {
				float = 0.0
			}
			return ac.FormatMoney(float)
		} else {
			return ac.FormatMoney(0)
		}
	}
}