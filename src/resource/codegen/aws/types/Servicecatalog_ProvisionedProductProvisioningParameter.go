package types

type Servicecatalog_ProvisionedProductProvisioningParameter struct {
	// Parameter value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Parameter key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Whether to ignore `value` and keep the previous parameter value. Ignored when initially provisioning a product.
	UsePreviousValue bool `json:"usePreviousValue,omitempty" yaml:"usePreviousValue,omitempty"`
}
