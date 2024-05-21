package types

type Cloudrun_ServiceTemplateSpecContainerEnvFrom struct {
	/*
	   The ConfigMap to select from.
	   Structure is documented below.
	*/
	ConfigMapRef Cloudrun_ServiceTemplateSpecContainerEnvFromConfigMapRef `json:"configMapRef,omitempty" yaml:"configMapRef,omitempty"`

	// An optional identifier to prepend to each key in the ConfigMap.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	/*
	   The Secret to select from.
	   Structure is documented below.
	*/
	SecretRef Cloudrun_ServiceTemplateSpecContainerEnvFromSecretRef `json:"secretRef,omitempty" yaml:"secretRef,omitempty"`
}
