package types

type Cloudrun_getServiceTemplateSpecContainerEnvFrom struct {
	// The ConfigMap to select from.
	ConfigMapReves []Cloudrun_getServiceTemplateSpecContainerEnvFromConfigMapRef `json:"configMapReves,omitempty" yaml:"configMapReves,omitempty"`

	// An optional identifier to prepend to each key in the ConfigMap.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The Secret to select from.
	SecretReves []Cloudrun_getServiceTemplateSpecContainerEnvFromSecretRef `json:"secretReves,omitempty" yaml:"secretReves,omitempty"`
}
