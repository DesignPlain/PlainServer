package types

type Batch_getJobDefinitionEksPropertyPodProperty struct {
	// Metadata about the Kubernetes pod.
	Metadatas []Batch_getJobDefinitionEksPropertyPodPropertyMetadata `json:"metadatas,omitempty" yaml:"metadatas,omitempty"`

	// The name of the service account that's used to run the pod.
	ServiceAccountName bool `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`

	// A list of data volumes used in a job.
	Volumes []Batch_getJobDefinitionEksPropertyPodPropertyVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// The properties of the container that's used on the Amazon EKS pod. Array of EksContainer objects.
	Containers []Batch_getJobDefinitionEksPropertyPodPropertyContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	// The DNS policy for the pod. The default value is ClusterFirst. If the hostNetwork parameter is not specified, the default is ClusterFirstWithHostNet. ClusterFirst indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node.
	DnsPolicy string `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`

	// Indicates if the pod uses the hosts' network IP address. The default value is true. Setting this to false enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	HostNetwork bool `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`
}
