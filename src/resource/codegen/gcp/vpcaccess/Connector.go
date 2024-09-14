package vpcaccess

import types "libds/gcp/types"

type Connector struct {
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	// Machine type of VM Instance underlying connector. Default is e2-micro
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances int `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput int `json:"minThroughput,omitempty" yaml:"minThroughput,omitempty"`

	// Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The subnet in which to house the connector
	   Structure is documented below.
	*/
	Subnet types.Vpcaccess_ConnectorSubnet `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	// Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300.
	MaxThroughput int `json:"maxThroughput,omitempty" yaml:"maxThroughput,omitempty"`

	// Minimum value of instances in autoscaling group underlying the connector.
	MinInstances int `json:"minInstances,omitempty" yaml:"minInstances,omitempty"`

	/*
	   The name of the resource (Max 25 characters).


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
