package types

type Lambda_CodeSigningConfigPolicies struct {
	// Code signing configuration policy for deployment validation failure. If you set the policy to Enforce, Lambda blocks the deployment request if code-signing validation checks fail. If you set the policy to Warn, Lambda allows the deployment and creates a CloudWatch log. Valid values: `Warn`, `Enforce`. Default value: `Warn`.
	UntrustedArtifactOnDeployment string `json:"untrustedArtifactOnDeployment,omitempty" yaml:"untrustedArtifactOnDeployment,omitempty"`
}
