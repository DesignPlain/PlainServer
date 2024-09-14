package types

type Compute_getResourcePolicySnapshotSchedulePolicySnapshotProperty struct {
	/*
	   Creates the new snapshot in the snapshot chain labeled with the
	   specified name. The chain name must be 1-63 characters long and comply
	   with RFC1035.
	*/
	ChainName string `json:"chainName,omitempty" yaml:"chainName,omitempty"`

	// Whether to perform a 'guest aware' snapshot.
	GuestFlush bool `json:"guestFlush,omitempty" yaml:"guestFlush,omitempty"`

	// A set of key-value pairs.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Cloud Storage bucket location to store the auto snapshot
	   (regional or multi-regional)
	*/
	StorageLocations []string `json:"storageLocations,omitempty" yaml:"storageLocations,omitempty"`
}
