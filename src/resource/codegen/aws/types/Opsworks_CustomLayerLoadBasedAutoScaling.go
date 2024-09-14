package types

type Opsworks_CustomLayerLoadBasedAutoScaling struct {
	// The downscaling settings, as defined below, used for load-based autoscaling
	Downscaling Opsworks_CustomLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	// Whether load-based auto scaling is enabled for the layer.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// The upscaling settings, as defined below, used for load-based autoscaling
	Upscaling Opsworks_CustomLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`
}
