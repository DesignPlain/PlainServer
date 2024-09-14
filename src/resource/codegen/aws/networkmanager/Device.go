package networkmanager

import types "libds/aws/types"

type Device struct {
	// A description of the device.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the global network.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The serial number of the device.
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`

	// Key-value tags for the device. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of device.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The vendor of the device.
	Vendor string `json:"vendor,omitempty" yaml:"vendor,omitempty"`

	// The AWS location of the device. Documented below.
	AwsLocation types.Networkmanager_DeviceAwsLocation `json:"awsLocation,omitempty" yaml:"awsLocation,omitempty"`

	// The model of device.
	Model string `json:"model,omitempty" yaml:"model,omitempty"`

	// The ID of the site.
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`

	// The location of the device. Documented below.
	Location types.Networkmanager_DeviceLocation `json:"location,omitempty" yaml:"location,omitempty"`
}
