package types

type Datafusion_InstanceAccelerator struct {
	/*
	   The type of an accelator for a CDF instance.
	   Possible values are: `CDC`, `HEALTHCARE`, `CCAI_INSIGHTS`.
	*/
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`

	/*
	   The type of an accelator for a CDF instance.
	   Possible values are: `ENABLED`, `DISABLED`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
