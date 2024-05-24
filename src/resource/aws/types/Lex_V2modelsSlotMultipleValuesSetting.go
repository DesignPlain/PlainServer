package types

type Lex_V2modelsSlotMultipleValuesSetting struct {
	// Whether a slot can return multiple values. When `true`, the slot may return more than one value in a response. When `false`, the slot returns only a single value. Multi-value slots are only available in the `en-US` locale.
	AllowMultipleValues bool `json:"allowMultipleValues,omitempty" yaml:"allowMultipleValues,omitempty"`
}
