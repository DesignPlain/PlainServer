package types

type Recaptcha_EnterpriseKeyWafSettings struct {
	// Supported WAF features. For more information, see https://cloud.google.com/recaptcha-enterprise/docs/usecase#comparison_of_features. Possible values: CHALLENGE_PAGE, SESSION_TOKEN, ACTION_TOKEN, EXPRESS
	WafFeature string `json:"wafFeature,omitempty" yaml:"wafFeature,omitempty"`

	// The WAF service that uses this key. Possible values: CA, FASTLY
	WafService string `json:"wafService,omitempty" yaml:"wafService,omitempty"`
}
