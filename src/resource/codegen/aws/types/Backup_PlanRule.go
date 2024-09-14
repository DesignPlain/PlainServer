package types

type Backup_PlanRule struct {
	// A CRON expression specifying when AWS Backup initiates a backup job.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The amount of time in minutes AWS Backup attempts a backup before canceling the job and returning an error.
	CompletionWindow int `json:"completionWindow,omitempty" yaml:"completionWindow,omitempty"`

	// An display name for a backup rule.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle Backup_PlanRuleLifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`

	// Metadata that you can assign to help organize the resources that you create.
	RecoveryPointTags map[string]string `json:"recoveryPointTags,omitempty" yaml:"recoveryPointTags,omitempty"`

	// The amount of time in minutes before beginning a backup.
	StartWindow int `json:"startWindow,omitempty" yaml:"startWindow,omitempty"`

	// The name of a logical container where backups are stored.
	TargetVaultName string `json:"targetVaultName,omitempty" yaml:"targetVaultName,omitempty"`

	// Configuration block(s) with copy operation settings. Detailed below.
	CopyActions []Backup_PlanRuleCopyAction `json:"copyActions,omitempty" yaml:"copyActions,omitempty"`

	// Enable continuous backups for supported resources.
	EnableContinuousBackup bool `json:"enableContinuousBackup,omitempty" yaml:"enableContinuousBackup,omitempty"`
}
