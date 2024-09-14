package lightsail

import types "libds/aws/types"

type Instance struct {
	// The add on configuration for the instance. Detailed below.
	AddOn types.Lightsail_InstanceAddOn `json:"addOn,omitempty" yaml:"addOn,omitempty"`

	/*
	   The ID for a virtual private server image. A list of available
	   blueprint IDs can be obtained using the AWS CLI command:
	   [`aws lightsail get-blueprints`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-blueprints.html).
	*/
	BlueprintId string `json:"blueprintId,omitempty" yaml:"blueprintId,omitempty"`

	/*
	   The bundle of specification information. A list of available
	   bundle IDs can be obtained using the AWS CLI command:
	   [`aws lightsail get-bundles`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-bundles.html).
	*/
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Single lined launch script as a string to configure server with additional user data
	UserData string `json:"userData,omitempty" yaml:"userData,omitempty"`

	/*
	   The Availability Zone in which to create your instance. A
	   list of available zones can be obtained using the AWS CLI command:
	   [`aws lightsail get-regions --include-availability-zones`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-regions.html).
	*/
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	/*
	   The name of your key pair. Created in the
	   Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
	*/
	KeyPairName string `json:"keyPairName,omitempty" yaml:"keyPairName,omitempty"`

	// The name of the Lightsail Instance. Names must be unique within each AWS Region in your Lightsail account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
