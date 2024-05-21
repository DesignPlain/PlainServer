package accesscontextmanager

import types "DesignSphere_Server/src/resource/gcp/types"

type AccessLevels struct {
	/*
	   The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	   Structure is documented below.
	*/
	AccessLevels []types.Accesscontextmanager_AccessLevelsAccessLevel `json:"accessLevels,omitempty" yaml:"accessLevels,omitempty"`

	/*
	   The AccessPolicy this AccessLevel lives in.
	   Format: accessPolicies/{policy_id}


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
