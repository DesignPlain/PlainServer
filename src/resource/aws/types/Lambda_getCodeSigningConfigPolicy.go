package types

type Lambda_getCodeSigningConfigPolicy struct {
	// Code signing configuration policy for deployment validation failure.
	UntrustedArtifactOnDeployment string `json:"untrustedArtifactOnDeployment,omitempty" yaml:"untrustedArtifactOnDeployment,omitempty"`
}
