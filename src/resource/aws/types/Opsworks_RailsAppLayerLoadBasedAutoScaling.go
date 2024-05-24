package types

type Opsworks_RailsAppLayerLoadBasedAutoScaling struct {
	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_RailsAppLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`

	//
	Downscaling Opsworks_RailsAppLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`
}
