package lambda

type LayerVersion struct {
	// Virtual attribute used to trigger replacement when source code changes. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`.
	SourceCodeHash string `json:"sourceCodeHash,omitempty" yaml:"sourceCodeHash,omitempty"`

	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x86_64` and `arm64` can be specified.
	CompatibleArchitectures []string `json:"compatibleArchitectures,omitempty" yaml:"compatibleArchitectures,omitempty"`

	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo string `json:"licenseInfo,omitempty" yaml:"licenseInfo,omitempty"`

	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" yaml:"s3ObjectVersion,omitempty"`

	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatible_architectures`, `compatible_runtimes`, `description`, `filename`, `layer_name`, `license_info`, `s3_bucket`, `s3_key`, `s3_object_version`, or `source_code_hash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 15 runtimes can be specified.
	CompatibleRuntimes []string `json:"compatibleRuntimes,omitempty" yaml:"compatibleRuntimes,omitempty"`

	// Description of what your Lambda Layer does.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Unique name for your Lambda Layer

	   The following arguments are optional:
	*/
	LayerName string `json:"layerName,omitempty" yaml:"layerName,omitempty"`

	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`
}
