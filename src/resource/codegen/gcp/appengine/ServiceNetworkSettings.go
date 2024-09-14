package appengine

import types "libds/gcp/types"

type ServiceNetworkSettings struct {
	/*
	   Ingress settings for this service. Will apply to all versions.
	   Structure is documented below.
	*/
	NetworkSettings types.Appengine_ServiceNetworkSettingsNetworkSettings `json:"networkSettings,omitempty" yaml:"networkSettings,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the service these settings apply to.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
