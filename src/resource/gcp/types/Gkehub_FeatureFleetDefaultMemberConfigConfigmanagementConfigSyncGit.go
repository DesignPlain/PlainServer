package types

type Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit struct {
	// The path within the Git repository that represents the top level of the repo to sync
	PolicyDir string `json:"policyDir,omitempty" yaml:"policyDir,omitempty"`

	// Type of secret configured for access to the Git repo
	SecretType string `json:"secretType,omitempty" yaml:"secretType,omitempty"`

	// The branch of the repository to sync from. Default: master
	SyncBranch string `json:"syncBranch,omitempty" yaml:"syncBranch,omitempty"`

	// The URL of the Git repository to use as the source of truth
	SyncRepo string `json:"syncRepo,omitempty" yaml:"syncRepo,omitempty"`

	// Git revision (tag or hash) to check out. Default HEAD
	SyncRev string `json:"syncRev,omitempty" yaml:"syncRev,omitempty"`

	// Period in seconds between consecutive syncs. Default: 15
	SyncWaitSecs string `json:"syncWaitSecs,omitempty" yaml:"syncWaitSecs,omitempty"`

	// The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount
	GcpServiceAccountEmail string `json:"gcpServiceAccountEmail,omitempty" yaml:"gcpServiceAccountEmail,omitempty"`

	// URL for the HTTPS Proxy to be used when communicating with the Git repo
	HttpsProxy string `json:"httpsProxy,omitempty" yaml:"httpsProxy,omitempty"`
}
