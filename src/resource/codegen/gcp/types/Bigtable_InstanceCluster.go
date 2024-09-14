package types

type Bigtable_InstanceCluster struct {
	// [Autoscaling](https://cloud.google.com/bigtable/docs/autoscaling#parameters) config for the cluster, contains the following arguments:
	AutoscalingConfig Bigtable_InstanceClusterAutoscalingConfig `json:"autoscalingConfig,omitempty" yaml:"autoscalingConfig,omitempty"`

	// The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	/*
	   Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the `cloudkms.cryptoKeyEncrypterDecrypter` role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster.

	   > --Note--: Removing the field entirely from the config will cause the provider to default to the backend value.

	   !> --Warning--: Modifying this field will cause the provider to delete/recreate the entire resource.

	   !> --Warning:-- Modifying the `storage_type`, `zone` or `kms_key_name` of an existing cluster (by
	   `cluster_id`) will cause the provider to delete/recreate the entire
	   `gcp.bigtable.Instance` resource. If these values are changing, use a new
	   `cluster_id`.
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   The number of nodes in the cluster.
	   If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50%!!(MISSING)s(MISSING)torage utilization.
	*/
	NumNodes int `json:"numNodes,omitempty" yaml:"numNodes,omitempty"`

	// The state of the cluster
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   The storage type to use. One of `"SSD"` or
	   `"HDD"`. Defaults to `"SSD"`.
	*/
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	/*
	   The zone to create the Cloud Bigtable cluster in. If it not
	   specified, the provider zone is used. Each cluster must have a different zone in the same region. Zones that support
	   Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
