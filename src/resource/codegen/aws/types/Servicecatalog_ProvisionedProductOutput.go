package types

type Servicecatalog_ProvisionedProductOutput struct {
	// The description of the output.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The output key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The output value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
