package types

type Networkmanager_getDeviceAwsLocation struct {
	// ARN of the subnet that the device is located in.
	SubnetArn string `json:"subnetArn,omitempty" yaml:"subnetArn,omitempty"`

	// Zone that the device is located in.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
