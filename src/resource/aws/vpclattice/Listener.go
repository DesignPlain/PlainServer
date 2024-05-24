package vpclattice

import types "DesignSphere_Server/src/resource/aws/types"

type Listener struct {
	// Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Protocol for the listener. Supported values are `HTTP` or `HTTPS`
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
	ServiceArn string `json:"serviceArn,omitempty" yaml:"serviceArn,omitempty"`

	/*
	   ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
	   > --NOTE:-- You must specify one of the following arguments: `service_arn` or `service_identifier`.
	*/
	ServiceIdentifier string `json:"serviceIdentifier,omitempty" yaml:"serviceIdentifier,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Default action block for the default listener rule. Default action blocks are defined below.
	DefaultAction types.Vpclattice_ListenerDefaultAction `json:"defaultAction,omitempty" yaml:"defaultAction,omitempty"`

	// Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
