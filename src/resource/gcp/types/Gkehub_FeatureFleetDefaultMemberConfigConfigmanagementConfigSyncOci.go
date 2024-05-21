package types

type Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci struct {
	// The OCI image repository URL for the package to sync from
	SyncRepo string `json:"syncRepo,omitempty" yaml:"syncRepo,omitempty"`

	// Period in seconds between consecutive syncs. Default: 15
	SyncWaitSecs string `json:"syncWaitSecs,omitempty" yaml:"syncWaitSecs,omitempty"`

	/*
	   (Optional, Deprecated)
	   Version of ACM installed

	   > --Warning:-- The `configmanagement.config_sync.oci.version` field is deprecated and will be removed in a future major release. Please use `configmanagement.version` field to specify the version of ACM installed instead.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount
	GcpServiceAccountEmail string `json:"gcpServiceAccountEmail,omitempty" yaml:"gcpServiceAccountEmail,omitempty"`

	// The absolute path of the directory that contains the local resources. Default: the root directory of the image
	PolicyDir string `json:"policyDir,omitempty" yaml:"policyDir,omitempty"`

	// Type of secret configured for access to the Git repo
	SecretType string `json:"secretType,omitempty" yaml:"secretType,omitempty"`
}
