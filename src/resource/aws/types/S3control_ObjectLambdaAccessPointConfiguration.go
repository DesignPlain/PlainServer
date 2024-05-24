package types

type S3control_ObjectLambdaAccessPointConfiguration struct {
	// Allowed features. Valid values: `GetObject-Range`, `GetObject-PartNumber`.
	AllowedFeatures []string `json:"allowedFeatures,omitempty" yaml:"allowedFeatures,omitempty"`

	// Whether or not the CloudWatch metrics configuration is enabled.
	CloudWatchMetricsEnabled bool `json:"cloudWatchMetricsEnabled,omitempty" yaml:"cloudWatchMetricsEnabled,omitempty"`

	// Standard access point associated with the Object Lambda Access Point.
	SupportingAccessPoint string `json:"supportingAccessPoint,omitempty" yaml:"supportingAccessPoint,omitempty"`

	// List of transformation configurations for the Object Lambda Access Point. See Transformation Configuration below for more details.
	TransformationConfigurations []S3control_ObjectLambdaAccessPointConfigurationTransformationConfiguration `json:"transformationConfigurations,omitempty" yaml:"transformationConfigurations,omitempty"`
}
