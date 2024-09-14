package types

type Ses_ReceiptRuleAddHeaderAction struct {
	// The name of the header to add
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// The value of the header to add
	HeaderValue string `json:"headerValue,omitempty" yaml:"headerValue,omitempty"`

	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`
}
