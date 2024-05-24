package cloudhsmv2

type Hsm struct {
	// The IDs of AZ in which HSM module will be located. Conflicts with `subnet_id`.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The ID of Cloud HSM v2 cluster to which HSM will be added.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// The IP address of HSM module. Must be within the CIDR of selected subnet.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// The ID of subnet in which HSM module will be located. Conflicts with `availability_zone`.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
