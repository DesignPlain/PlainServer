package types

type Backup_getPlanRule struct {
	//
	CompletionWindow int `json:"completionWindow,omitempty" yaml:"completionWindow,omitempty"`

	//
	EnableContinuousBackup bool `json:"enableContinuousBackup,omitempty" yaml:"enableContinuousBackup,omitempty"`

	//
	Lifecycles []Backup_getPlanRuleLifecycle `json:"lifecycles,omitempty" yaml:"lifecycles,omitempty"`

	//
	StartWindow int `json:"startWindow,omitempty" yaml:"startWindow,omitempty"`

	//
	CopyActions []Backup_getPlanRuleCopyAction `json:"copyActions,omitempty" yaml:"copyActions,omitempty"`

	//
	RecoveryPointTags map[string]string `json:"recoveryPointTags,omitempty" yaml:"recoveryPointTags,omitempty"`

	//
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	//
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	//
	TargetVaultName string `json:"targetVaultName,omitempty" yaml:"targetVaultName,omitempty"`
}
