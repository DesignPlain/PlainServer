package types

type Opsworks_HaproxyLayerLoadBasedAutoScaling struct {
	//
	Downscaling Opsworks_HaproxyLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_HaproxyLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`
}
