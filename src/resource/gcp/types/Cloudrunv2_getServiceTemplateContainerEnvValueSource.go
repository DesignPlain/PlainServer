package types

type Cloudrunv2_getServiceTemplateContainerEnvValueSource struct {
	// Selects a secret and a specific version from Cloud Secret Manager.
	SecretKeyReves []Cloudrunv2_getServiceTemplateContainerEnvValueSourceSecretKeyRef `json:"secretKeyReves,omitempty" yaml:"secretKeyReves,omitempty"`
}
