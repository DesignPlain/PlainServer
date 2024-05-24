package types

type Lex_V2modelsIntentKendraConfiguration struct {
	// ARN of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search. The index must be in the same account and Region as the Amazon Lex bot.
	KendraIndex string `json:"kendraIndex,omitempty" yaml:"kendraIndex,omitempty"`

	// Query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query. The filter is in the format defined by Amazon Kendra. For more information, see [Filtering queries](https://docs.aws.amazon.com/kendra/latest/dg/filtering.html).
	QueryFilterString string `json:"queryFilterString,omitempty" yaml:"queryFilterString,omitempty"`

	// Whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.
	QueryFilterStringEnabled bool `json:"queryFilterStringEnabled,omitempty" yaml:"queryFilterStringEnabled,omitempty"`
}
