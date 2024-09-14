package types

type Lambda_getFunctionFileSystemConfig struct {
	// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) ARN identifying your Lambda Function. See also `qualified_arn`.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	//
	LocalMountPath string `json:"localMountPath,omitempty" yaml:"localMountPath,omitempty"`
}
