package types

type Cloudasset_getResourcesSearchAllResult struct {
	// Additional searchable attributes of this resource. Informational only. The exact set of attributes is subject to change. For example: project id, DNS name etc.
	AdditionalAttributes []string `json:"additionalAttributes,omitempty" yaml:"additionalAttributes,omitempty"`

	// One or more paragraphs of text description of this resource. Maximum length could be up to 1M bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The display name of this resource.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Labels associated with this resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Location can be `global`, regional like `us-east1`, or zonal like `us-west1-b`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The project that this resource belongs to, in the form of `projects/{project_number}`.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The type of this resource.
	AssetType string `json:"assetType,omitempty" yaml:"assetType,omitempty"`

	// The full resource name. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Network tags associated with this resource.
	NetworkTags []string `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`
}
