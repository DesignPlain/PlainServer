package tpu

import types "libds/gcp/types"

type V2Vm struct {
	/*
	   Resource labels to represent user-provided metadata.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The immutable name of the TPU.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   TPU accelerator type for the TPU. `accelerator_type` cannot be used at the same time as
	   `accelerator_config`. If neither is specified, `accelerator_type` defaults to 'v2-8'.
	*/
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`

	/*
	   The additional data disks for the Node.
	   Structure is documented below.
	*/
	DataDisks []types.Tpu_V2VmDataDisk `json:"dataDisks,omitempty" yaml:"dataDisks,omitempty"`

	/*
	   Network configurations for the TPU node.
	   Structure is documented below.
	*/
	NetworkConfig types.Tpu_V2VmNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is
	   specified, the default compute service account will be used.
	   Structure is documented below.
	*/
	ServiceAccount types.Tpu_V2VmServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   The scheduling options for this node.
	   Structure is documented below.
	*/
	SchedulingConfig types.Tpu_V2VmSchedulingConfig `json:"schedulingConfig,omitempty" yaml:"schedulingConfig,omitempty"`

	/*
	   The AccleratorConfig for the TPU Node. `accelerator_config` cannot be used at the same time
	   as `accelerator_type`. If neither is specified, `accelerator_type` defaults to 'v2-8'.
	   Structure is documented below.
	*/
	AcceleratorConfig types.Tpu_V2VmAcceleratorConfig `json:"acceleratorConfig,omitempty" yaml:"acceleratorConfig,omitempty"`

	/*
	   The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must
	   be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger
	   block would be wasteful (a node can only consume one IP address). Errors will occur if the
	   CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts
	   with any subnetworks in the user's provided network, or the provided network is peered with
	   another network that is using that CIDR block.
	*/
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// Text description of the TPU.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   Runtime version for the TPU.


	   - - -
	*/
	RuntimeVersion string `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`

	/*
	   Shielded Instance options.
	   Structure is documented below.
	*/
	ShieldedInstanceConfig types.Tpu_V2VmShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
