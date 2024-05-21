package types

type Cloudrun_getServiceTemplateSpecContainerLivenessProbeHttpGetHttpHeader struct {
	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The header field value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
