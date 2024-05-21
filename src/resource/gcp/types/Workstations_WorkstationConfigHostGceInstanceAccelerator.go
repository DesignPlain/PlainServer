package types

type Workstations_WorkstationConfigHostGceInstanceAccelerator struct {
	// Number of accelerator cards exposed to the instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// Type of accelerator resource to attach to the instance, for example, "nvidia-tesla-p100".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
