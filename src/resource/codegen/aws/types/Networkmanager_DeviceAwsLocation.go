package types

type Networkmanager_DeviceAwsLocation struct {
	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	SubnetArn string `json:"subnetArn,omitempty" yaml:"subnetArn,omitempty"`

	// The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
