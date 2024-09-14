package vmwareengine

type Subnet struct {
	// The IP address range of the subnet in CIDR format.
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	/*
	   The ID of the subnet. For userDefined subnets, this name should be in the format of "service-n",
	   where n ranges from 1 to 5.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The resource name of the private cloud to create a new subnet in.
	   Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	   For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
