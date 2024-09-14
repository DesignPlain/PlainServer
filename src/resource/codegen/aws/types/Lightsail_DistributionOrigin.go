package types

type Lightsail_DistributionOrigin struct {
	// The AWS Region name of the origin resource.
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`

	// The resource type of the origin resource (e.g., Instance).
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// The name of the origin resource. Your origin can be an instance with an attached static IP, a bucket, or a load balancer that has at least one instance attached to it.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	ProtocolPolicy string `json:"protocolPolicy,omitempty" yaml:"protocolPolicy,omitempty"`
}
