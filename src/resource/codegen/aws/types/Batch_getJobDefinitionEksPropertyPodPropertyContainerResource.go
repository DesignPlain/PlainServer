package types

type Batch_getJobDefinitionEksPropertyPodPropertyContainerResource struct {
	// The type and quantity of the resources to reserve for the container.
	Limits map[string]string `json:"limits,omitempty" yaml:"limits,omitempty"`

	// The type and quantity of the resources to request for the container.
	Requests map[string]string `json:"requests,omitempty" yaml:"requests,omitempty"`
}
