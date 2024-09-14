package types

type Backup_getPlanRuleCopyAction struct {
	//
	DestinationVaultArn string `json:"destinationVaultArn,omitempty" yaml:"destinationVaultArn,omitempty"`

	//
	Lifecycles []Backup_getPlanRuleCopyActionLifecycle `json:"lifecycles,omitempty" yaml:"lifecycles,omitempty"`
}
