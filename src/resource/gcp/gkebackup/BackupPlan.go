package gkebackup

import types "DesignSphere_Server/src/resource/gcp/types"

type BackupPlan struct {
	/*
	   The region of the Backup Plan.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   RetentionPolicy governs lifecycle of Backups created under this plan.
	   Structure is documented below.
	*/
	RetentionPolicy types.Gkebackup_BackupPlanRetentionPolicy `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`

	/*
	   Defines the configuration of Backups created via this BackupPlan.
	   Structure is documented below.
	*/
	BackupConfig types.Gkebackup_BackupPlanBackupConfig `json:"backupConfig,omitempty" yaml:"backupConfig,omitempty"`

	/*
	   Defines a schedule for automatic Backup creation via this BackupPlan.
	   Structure is documented below.
	*/
	BackupSchedule types.Gkebackup_BackupPlanBackupSchedule `json:"backupSchedule,omitempty" yaml:"backupSchedule,omitempty"`

	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	/*
	   This flag indicates whether this BackupPlan has been deactivated.
	   Setting this field to True locks the BackupPlan such that no further updates will be allowed
	   (except deletes), including the deactivated field itself. It also prevents any new Backups
	   from being created via this BackupPlan (including scheduled Backups).
	*/
	Deactivated bool `json:"deactivated,omitempty" yaml:"deactivated,omitempty"`

	/*
	   Description: A set of custom labels supplied by the user.
	   A list of key->value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// User specified descriptive string for this BackupPlan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The full name of the BackupPlan Resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
