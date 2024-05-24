package evidently

import types "DesignSphere_Server/src/resource/aws/types"

type Project struct {
	// A name for the project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the project. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A block that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so. If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view. See below.
	DataDelivery types.Evidently_ProjectDataDelivery `json:"dataDelivery,omitempty" yaml:"dataDelivery,omitempty"`

	// Specifies the description of the project.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
