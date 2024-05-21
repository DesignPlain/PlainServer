package types



type Folder_getOrganizationPolicyListPolicy struct {
	// One or the other must be set.
	Denies []Folder_getOrganizationPolicyListPolicyDeny `json:"denies,omitempty" yaml:"denies,omitempty"`

	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	InheritFromParent bool `json:"inheritFromParent,omitempty" yaml:"inheritFromParent,omitempty"`

	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	SuggestedValue string `json:"suggestedValue,omitempty" yaml:"suggestedValue,omitempty"`

	// One or the other must be set.
	Allows []Folder_getOrganizationPolicyListPolicyAllow `json:"allows,omitempty" yaml:"allows,omitempty"`
}
