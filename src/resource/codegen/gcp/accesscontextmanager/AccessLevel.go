package accesscontextmanager

import types "libds/gcp/types"

type AccessLevel struct {
	/*
	   A set of predefined conditions for the access level and a combining function.
	   Structure is documented below.
	*/
	Basic types.Accesscontextmanager_AccessLevelBasic `json:"basic,omitempty" yaml:"basic,omitempty"`

	/*
	   Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	   See CEL spec at: https://github.com/google/cel-spec.
	   Structure is documented below.
	*/
	Custom types.Accesscontextmanager_AccessLevelCustom `json:"custom,omitempty" yaml:"custom,omitempty"`

	// Description of the AccessLevel and its use. Does not affect behavior.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Resource name for the Access Level. The short_name component must begin
	   with a letter and only include alphanumeric and '_'.
	   Format: accessPolicies/{policy_id}/accessLevels/{short_name}


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The AccessPolicy this AccessLevel lives in.
	   Format: accessPolicies/{policy_id}
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// Human readable title. Must be unique within the Policy.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
