package types

type Gkehub_FeatureMembershipConfigmanagementConfigSyncOci struct {
	// The GCP Service Account Email used for auth when secret_type is gcpserviceaccount.
	GcpServiceAccountEmail string `json:"gcpServiceAccountEmail,omitempty" yaml:"gcpServiceAccountEmail,omitempty"`

	// The absolute path of the directory that contains the local resources. Default: the root directory of the image.
	PolicyDir string `json:"policyDir,omitempty" yaml:"policyDir,omitempty"`

	// Type of secret configured for access to the OCI Image. Must be one of gcenode, gcpserviceaccount or none.
	SecretType string `json:"secretType,omitempty" yaml:"secretType,omitempty"`

	// The OCI image repository URL for the package to sync from. e.g. LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME.
	SyncRepo string `json:"syncRepo,omitempty" yaml:"syncRepo,omitempty"`

	// Period in seconds(int64 format) between consecutive syncs. Default: 15.
	SyncWaitSecs string `json:"syncWaitSecs,omitempty" yaml:"syncWaitSecs,omitempty"`
}
