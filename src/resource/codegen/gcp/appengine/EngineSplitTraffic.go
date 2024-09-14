package appengine

import types "libds/gcp/types"

type EngineSplitTraffic struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic bool `json:"migrateTraffic,omitempty" yaml:"migrateTraffic,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the service these settings apply to.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	   Structure is documented below.
	*/
	Split types.Appengine_EngineSplitTrafficSplit `json:"split,omitempty" yaml:"split,omitempty"`
}
