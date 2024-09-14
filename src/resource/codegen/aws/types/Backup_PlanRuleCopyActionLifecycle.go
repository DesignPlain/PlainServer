package types

type Backup_PlanRuleCopyActionLifecycle struct {
	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `cold_storage_after`.
	DeleteAfter int `json:"deleteAfter,omitempty" yaml:"deleteAfter,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	OptInToArchiveForSupportedResources bool `json:"optInToArchiveForSupportedResources,omitempty" yaml:"optInToArchiveForSupportedResources,omitempty"`

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter int `json:"coldStorageAfter,omitempty" yaml:"coldStorageAfter,omitempty"`
}
