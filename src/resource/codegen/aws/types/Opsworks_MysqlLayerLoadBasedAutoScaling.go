package types

type Opsworks_MysqlLayerLoadBasedAutoScaling struct {
	//
	Downscaling Opsworks_MysqlLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`

	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_MysqlLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`
}
