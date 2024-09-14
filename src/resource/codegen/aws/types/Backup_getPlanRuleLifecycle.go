package types

type Backup_getPlanRuleLifecycle struct {
	//
	ColdStorageAfter int `json:"coldStorageAfter,omitempty" yaml:"coldStorageAfter,omitempty"`

	//
	DeleteAfter int `json:"deleteAfter,omitempty" yaml:"deleteAfter,omitempty"`

	//
	OptInToArchiveForSupportedResources bool `json:"optInToArchiveForSupportedResources,omitempty" yaml:"optInToArchiveForSupportedResources,omitempty"`
}
