package backup

import types "DesignSphere_Server/src/resource/aws/types"

type Plan struct {
	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings []types.Backup_PlanAdvancedBackupSetting `json:"advancedBackupSettings,omitempty" yaml:"advancedBackupSettings,omitempty"`

	// The display name of a backup plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules []types.Backup_PlanRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Metadata that you can assign to help organize the plans you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
