package appengine

import types "DesignSphere_Server/src/resource/gcp/types"

type ApplicationUrlDispatchRules struct {
	/*
	   Rules to match an HTTP request and dispatch that request to a service.
	   Structure is documented below.
	*/
	DispatchRules []types.Appengine_ApplicationUrlDispatchRulesDispatchRule `json:"dispatchRules,omitempty" yaml:"dispatchRules,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
