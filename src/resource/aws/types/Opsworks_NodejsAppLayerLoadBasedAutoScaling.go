package types

type Opsworks_NodejsAppLayerLoadBasedAutoScaling struct {
	//
	Downscaling Opsworks_NodejsAppLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_NodejsAppLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`
}
