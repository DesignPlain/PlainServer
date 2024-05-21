package networksecurity

import types "DesignSphere_Server/src/resource/gcp/types"

type SecurityProfile struct {
	// An optional description of the security profile. The Max length is 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A map of key/value label pairs to assign to the resource.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the security profile.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name of the security profile resource.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of the parent this security profile belongs to.
	   Format: organizations/{organization_id}.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The threat prevention configuration for the security profile.
	   Structure is documented below.
	*/
	ThreatPreventionProfile types.Networksecurity_SecurityProfileThreatPreventionProfile `json:"threatPreventionProfile,omitempty" yaml:"threatPreventionProfile,omitempty"`

	/*
	   The type of security profile.
	   Possible values are: `THREAT_PREVENTION`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
