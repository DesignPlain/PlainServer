package types

type Notebooks_InstanceAcceleratorConfig struct {
	/*
	   Type of this accelerator.
	   Possible values are: `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, `TPU_V3`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Count of cores of this accelerator.
	CoreCount int `json:"coreCount,omitempty" yaml:"coreCount,omitempty"`
}
