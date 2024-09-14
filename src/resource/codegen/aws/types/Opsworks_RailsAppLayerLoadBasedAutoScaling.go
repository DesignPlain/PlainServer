package types

type Opsworks_RailsAppLayerLoadBasedAutoScaling struct {
	//
	Downscaling Opsworks_RailsAppLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_RailsAppLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`
}
