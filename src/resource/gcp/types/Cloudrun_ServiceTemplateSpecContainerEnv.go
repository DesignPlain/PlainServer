package types

type Cloudrun_ServiceTemplateSpecContainerEnv struct {
	// Name of the environment variable.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Defaults to "".
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   Source for the environment variable's value. Only supports secret_key_ref.
	   Structure is documented below.
	*/
	ValueFrom Cloudrun_ServiceTemplateSpecContainerEnvValueFrom `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}
