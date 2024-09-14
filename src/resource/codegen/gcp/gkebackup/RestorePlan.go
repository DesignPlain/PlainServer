package gkebackup

import types "libds/gcp/types"

type RestorePlan struct {
	/*
	   Description: A set of custom labels supplied by the user.
	   A list of key->value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The region of the Restore Plan.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The full name of the BackupPlan Resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Defines the configuration of Restores created via this RestorePlan.
	   Structure is documented below.
	*/
	RestoreConfig types.Gkebackup_RestorePlanRestoreConfig `json:"restoreConfig,omitempty" yaml:"restoreConfig,omitempty"`

	/*
	   A reference to the BackupPlan from which Backups may be used
	   as the source for Restores created via this RestorePlan.
	*/
	BackupPlan string `json:"backupPlan,omitempty" yaml:"backupPlan,omitempty"`

	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// User specified descriptive string for this RestorePlan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
