package medialive

import types "libds/aws/types"

type Input struct {
	// A list of the MediaConnect Flows. See Media Connect Flows for more details.
	MediaConnectFlows []types.Medialive_InputMediaConnectFlow `json:"mediaConnectFlows,omitempty" yaml:"mediaConnectFlows,omitempty"`

	// Name of the input.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The source URLs for a PULL-type input. See Sources for more details.
	Sources []types.Medialive_InputSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// List of input security groups.
	InputSecurityGroups []string `json:"inputSecurityGroups,omitempty" yaml:"inputSecurityGroups,omitempty"`

	// Settings for the devices. See Input Devices for more details.
	InputDevices []types.Medialive_InputInputDevice `json:"inputDevices,omitempty" yaml:"inputDevices,omitempty"`

	// The ARN of the role this input assumes during and after creation.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the Input. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The different types of inputs that AWS Elemental MediaLive supports.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Settings for a private VPC Input. See VPC for more details.
	Vpc types.Medialive_InputVpc `json:"vpc,omitempty" yaml:"vpc,omitempty"`

	// Destination settings for PUSH type inputs. See Destinations for more details.
	Destinations []types.Medialive_InputDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`
}
