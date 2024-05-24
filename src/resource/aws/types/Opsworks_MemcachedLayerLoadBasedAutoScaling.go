package types

type Opsworks_MemcachedLayerLoadBasedAutoScaling struct {
	//
	Upscaling Opsworks_MemcachedLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`

	//
	Downscaling Opsworks_MemcachedLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}
