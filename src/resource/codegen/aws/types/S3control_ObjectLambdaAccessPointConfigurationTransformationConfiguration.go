package types

type S3control_ObjectLambdaAccessPointConfigurationTransformationConfiguration struct {
	// The content transformation of an Object Lambda Access Point configuration. See Content Transformation below for more details.
	ContentTransformation S3control_ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformation `json:"contentTransformation,omitempty" yaml:"contentTransformation,omitempty"`

	// The actions of an Object Lambda Access Point configuration. Valid values: `GetObject`.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`
}
