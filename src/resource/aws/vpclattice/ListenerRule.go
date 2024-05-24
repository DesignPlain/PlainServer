package vpclattice

import types "DesignSphere_Server/src/resource/aws/types"

type ListenerRule struct {
	// The ID or Amazon Resource Name (ARN) of the listener.
	ListenerIdentifier string `json:"listenerIdentifier,omitempty" yaml:"listenerIdentifier,omitempty"`

	// The rule match.
	Match types.Vpclattice_ListenerRuleMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.

	   The following arguments are optional:
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The ID or Amazon Resource Identifier (ARN) of the service.
	ServiceIdentifier string `json:"serviceIdentifier,omitempty" yaml:"serviceIdentifier,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The action for the listener rule.
	Action types.Vpclattice_ListenerRuleAction `json:"action,omitempty" yaml:"action,omitempty"`
}
