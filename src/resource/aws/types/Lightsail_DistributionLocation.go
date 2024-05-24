package types

type Lightsail_DistributionLocation struct {
	// The Availability Zone. Follows the format us-east-2a (case-sensitive).
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The AWS Region name of the origin resource.
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`
}
