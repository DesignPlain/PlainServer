package types

type Cloudrun_getServiceTemplateSpecContainerEnvValueFrom struct {
	// Selects a key (version) of a secret in Secret Manager.
	SecretKeyReves []Cloudrun_getServiceTemplateSpecContainerEnvValueFromSecretKeyRef `json:"secretKeyReves,omitempty" yaml:"secretKeyReves,omitempty"`
}
