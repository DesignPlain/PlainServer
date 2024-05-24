package types

type Amplify_AppProductionBranch struct {
	// Last deploy time of the production branch.
	LastDeployTime string `json:"lastDeployTime,omitempty" yaml:"lastDeployTime,omitempty"`

	// Status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Thumbnail URL for the production branch.
	ThumbnailUrl string `json:"thumbnailUrl,omitempty" yaml:"thumbnailUrl,omitempty"`

	// Branch name for the production branch.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`
}
