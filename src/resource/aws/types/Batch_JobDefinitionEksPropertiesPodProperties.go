package types

type Batch_JobDefinitionEksPropertiesPodProperties struct {
	// Specifies the volumes for a job definition that uses Amazon EKS resources. AWS Batch supports emptyDir, hostPath, and secret volume
	Volumes []Batch_JobDefinitionEksPropertiesPodPropertiesVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// The properties of the container that's used on the Amazon EKS pod. See containers below.
	Containers Batch_JobDefinitionEksPropertiesPodPropertiesContainers `json:"containers,omitempty" yaml:"containers,omitempty"`

	// The DNS policy for the pod. The default value is `ClusterFirst`. If the `host_network` argument is not specified, the default is `ClusterFirstWithHostNet`. `ClusterFirst` indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node. For more information, see Pod's DNS policy in the Kubernetes documentation.
	DnsPolicy string `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`

	// Indicates if the pod uses the hosts' network IP address. The default value is `true`. Setting this to `false` enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	HostNetwork bool `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`

	// Metadata about the Kubernetes pod.
	Metadata Batch_JobDefinitionEksPropertiesPodPropertiesMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// The name of the service account that's used to run the pod.
	ServiceAccountName string `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
}
