package types

type Cloudrunv2_JobTemplateTemplateContainerEnvValueSource struct {
	/*
	   Selects a secret and a specific version from Cloud Secret Manager.
	   Structure is documented below.
	*/
	SecretKeyRef Cloudrunv2_JobTemplateTemplateContainerEnvValueSourceSecretKeyRef `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}
