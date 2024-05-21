package tpu

import types "DesignSphere_Server/src/resource/gcp/types"

type Node struct {
	// The immutable name of the TPU.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of a network to peer the TPU node to. It must be a
	   preexisting Compute Engine network inside of the project on which
	   this API has been activated. If none is provided, "default" will be
	   used.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The version of Tensorflow running in the Node.


	   - - -
	*/
	TensorflowVersion string `json:"tensorflowVersion,omitempty" yaml:"tensorflowVersion,omitempty"`

	// The type of hardware accelerators associated with this node.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`

	/*
	   Resource labels to represent user provided metadata.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Sets the scheduling options for this TPU instance.
	   Structure is documented below.
	*/
	SchedulingConfig types.Tpu_NodeSchedulingConfig `json:"schedulingConfig,omitempty" yaml:"schedulingConfig,omitempty"`

	/*
	   Whether the VPC peering for the node is set up through Service Networking API.
	   The VPC Peering should be set up before provisioning the node. If this field is set,
	   cidr_block field should not be specified. If the network that you want to peer the
	   TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	*/
	UseServiceNetworking bool `json:"useServiceNetworking,omitempty" yaml:"useServiceNetworking,omitempty"`

	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   The CIDR block that the TPU node will use when selecting an IP
	   address. This CIDR block must be a /29 block; the Compute Engine
	   networks API forbids a smaller block, and using a larger block would
	   be wasteful (a node can only consume one IP address).
	   Errors will occur if the CIDR block has already been used for a
	   currently existing TPU node, the CIDR block conflicts with any
	   subnetworks in the user's provided network, or the provided network
	   is peered with another network that is using that CIDR block.
	*/
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
