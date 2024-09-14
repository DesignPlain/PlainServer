package types

type Cloudrun_ServiceTemplateSpecContainerEnvValueFrom struct {
	/*
	   Selects a key (version) of a secret in Secret Manager.
	   Structure is documented below.
	*/
	SecretKeyRef Cloudrun_ServiceTemplateSpecContainerEnvValueFromSecretKeyRef `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}
