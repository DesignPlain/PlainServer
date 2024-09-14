package types

type Looker_InstanceUserMetadata struct {
	// Number of additional Developer Users to allocate to the Looker Instance.
	AdditionalDeveloperUserCount int `json:"additionalDeveloperUserCount,omitempty" yaml:"additionalDeveloperUserCount,omitempty"`

	// Number of additional Standard Users to allocate to the Looker Instance.
	AdditionalStandardUserCount int `json:"additionalStandardUserCount,omitempty" yaml:"additionalStandardUserCount,omitempty"`

	// Number of additional Viewer Users to allocate to the Looker Instance.
	AdditionalViewerUserCount int `json:"additionalViewerUserCount,omitempty" yaml:"additionalViewerUserCount,omitempty"`
}
