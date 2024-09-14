package databasemigrationservice

import types "libds/gcp/types"

type PrivateConnection struct {
	// Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Labels.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the location this private connection is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The private connectivity identifier.
	PrivateConnectionId string `json:"privateConnectionId,omitempty" yaml:"privateConnectionId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The VPC Peering configuration is used to create VPC peering
	   between databasemigrationservice and the consumer's VPC.
	   Structure is documented below.
	*/
	VpcPeeringConfig types.Databasemigrationservice_PrivateConnectionVpcPeeringConfig `json:"vpcPeeringConfig,omitempty" yaml:"vpcPeeringConfig,omitempty"`
}
