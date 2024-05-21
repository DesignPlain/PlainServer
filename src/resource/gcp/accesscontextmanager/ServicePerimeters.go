package accesscontextmanager

import types "DesignSphere_Server/src/resource/gcp/types"

type ServicePerimeters struct {
	/*
	   The AccessPolicy this ServicePerimeter lives in.
	   Format: accessPolicies/{policy_id}


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	   Structure is documented below.
	*/
	ServicePerimeters []types.Accesscontextmanager_ServicePerimetersServicePerimeter `json:"servicePerimeters,omitempty" yaml:"servicePerimeters,omitempty"`
}
