package types

type Cloudrunv2_getJobTemplateTemplateContainerEnvValueSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	SecretKeyReves []Cloudrunv2_getJobTemplateTemplateContainerEnvValueSourceSecretKeyRef `json:"secretKeyReves,omitempty" yaml:"secretKeyReves,omitempty"`
}
