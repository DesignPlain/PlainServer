package types

type Container_ClusterDatabaseEncryption struct {
	/*
	   the key to use to encrypt/decrypt secrets.  See the [DatabaseEncryption definition](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#Cluster.DatabaseEncryption) for more information.

	   <a name="nested_enable_k8s_beta_apis"></a>The `enable_k8s_beta_apis` block supports:
	*/
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// `ENCRYPTED` or `DECRYPTED`
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
