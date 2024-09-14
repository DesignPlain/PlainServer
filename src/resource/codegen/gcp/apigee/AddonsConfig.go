package apigee

import types "libds/gcp/types"

type AddonsConfig struct {
	/*
	   Addon configurations of the Apigee organization.
	   Structure is documented below.
	*/
	AddonsConfig types.Apigee_AddonsConfigAddonsConfig `json:"addonsConfig,omitempty" yaml:"addonsConfig,omitempty"`

	/*
	   Name of the Apigee organization.


	   - - -
	*/
	Org string `json:"org,omitempty" yaml:"org,omitempty"`
}
