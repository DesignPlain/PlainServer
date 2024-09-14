package backupdisasterrecovery

import types "libds/gcp/types"

type ManagementServer struct {
	// The location for the management server (management console)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of management server (management console)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Network details to create management server (management console).
	   Structure is documented below.
	*/
	Networks []types.Backupdisasterrecovery_ManagementServerNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The type of management server (management console).
	   Default value is `BACKUP_RESTORE`.
	   Possible values are: `BACKUP_RESTORE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
