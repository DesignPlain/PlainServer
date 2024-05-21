package types

type Clouddomains_RegistrationYearlyPrice struct {
	// The three-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty" yaml:"currencyCode,omitempty"`

	// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
	Units string `json:"units,omitempty" yaml:"units,omitempty"`
}
