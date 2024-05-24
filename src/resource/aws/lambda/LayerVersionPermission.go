package lambda

type LayerVersionPermission struct {
	// The name or ARN of the Lambda Layer, which you want to grant access to.
	LayerName string `json:"layerName,omitempty" yaml:"layerName,omitempty"`

	// An identifier of AWS Organization, which should be able to use your Lambda Layer. `principal` should be equal to `-` if `organization_id` provided.
	OrganizationId string `json:"organizationId,omitempty" yaml:"organizationId,omitempty"`

	// AWS account ID which should be able to use your Lambda Layer. `-` can be used here, if you want to share your Lambda Layer widely.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatible_architectures`, `compatible_runtimes`, `description`, `filename`, `layer_name`, `license_info`, `s3_bucket`, `s3_key`, `s3_object_version`, or `source_code_hash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// The name of Lambda Layer Permission, for example `dev-account` - human readable note about what is this permission for.
	StatementId string `json:"statementId,omitempty" yaml:"statementId,omitempty"`

	// Version of Lambda Layer, which you want to grant access to. Note: permissions only apply to a single version of a layer.
	VersionNumber int `json:"versionNumber,omitempty" yaml:"versionNumber,omitempty"`

	// Action, which will be allowed. `lambda:GetLayerVersion` value is suggested by AWS documantation.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
}
