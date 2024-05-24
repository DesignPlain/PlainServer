package rds

type Snapshot struct {
	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier string `json:"dbInstanceIdentifier,omitempty" yaml:"dbInstanceIdentifier,omitempty"`

	// The Identifier for the snapshot.
	DbSnapshotIdentifier string `json:"dbSnapshotIdentifier,omitempty" yaml:"dbSnapshotIdentifier,omitempty"`

	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts []string `json:"sharedAccounts,omitempty" yaml:"sharedAccounts,omitempty"`
}
