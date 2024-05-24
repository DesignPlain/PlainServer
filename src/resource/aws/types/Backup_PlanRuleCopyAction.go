package types

type Backup_PlanRuleCopyAction struct {
	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	DestinationVaultArn string `json:"destinationVaultArn,omitempty" yaml:"destinationVaultArn,omitempty"`

	// The lifecycle defines when a protected resource is copied over to a backup vault and when it expires.  Fields documented above.
	Lifecycle Backup_PlanRuleCopyActionLifecycle `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
}
