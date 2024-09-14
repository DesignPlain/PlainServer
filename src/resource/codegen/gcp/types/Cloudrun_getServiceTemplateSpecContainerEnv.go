package types

type Cloudrun_getServiceTemplateSpecContainerEnv struct {
	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Defaults to "".
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Source for the environment variable's value. Only supports secret_key_ref.
	ValueFroms []Cloudrun_getServiceTemplateSpecContainerEnvValueFrom `json:"valueFroms,omitempty" yaml:"valueFroms,omitempty"`
}
