package networksecurity

import types "DesignSphere_Server/src/resource/gcp/types"

type AuthorizationPolicy struct {
	/*
	   The action to take when a rule match is found. Possible values are "ALLOW" or "DENY".
	   Possible values are: `ALLOW`, `DENY`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the AuthorizationPolicy resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the authorization policy.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of the AuthorizationPolicy resource.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken.
	   A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
	   Structure is documented below.
	*/
	Rules []types.Networksecurity_AuthorizationPolicyRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
