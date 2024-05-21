package types

type Cloudrunv2_ServiceTemplateContainerEnvValueSource struct {
	/*
	   Selects a secret and a specific version from Cloud Secret Manager.
	   Structure is documented below.
	*/
	SecretKeyRef Cloudrunv2_ServiceTemplateContainerEnvValueSourceSecretKeyRef `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}
