package backup

type VaultPolicy struct {
	// Name of the backup vault to add policy for.
	BackupVaultName string `json:"backupVaultName,omitempty" yaml:"backupVaultName,omitempty"`

	// The backup vault access policy document in JSON format.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
