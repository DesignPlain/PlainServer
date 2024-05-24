package types

type Servicecatalog_ProvisionedProductProvisioningParameter struct {
	// Parameter key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Whether to ignore `value` and keep the previous parameter value. Ignored when initially provisioning a product.
	UsePreviousValue bool `json:"usePreviousValue,omitempty" yaml:"usePreviousValue,omitempty"`

	// Parameter value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
