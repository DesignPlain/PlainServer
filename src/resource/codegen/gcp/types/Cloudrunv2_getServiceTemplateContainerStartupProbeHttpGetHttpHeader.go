package types

type Cloudrunv2_getServiceTemplateContainerStartupProbeHttpGetHttpHeader struct {
	// The name of the Cloud Run v2 Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The header field value
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
