package glacier

type VaultLock struct {
	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock bool `json:"completeLock,omitempty" yaml:"completeLock,omitempty"`

	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `complete_lock` being set to `true`.
	IgnoreDeletionError bool `json:"ignoreDeletionError,omitempty" yaml:"ignoreDeletionError,omitempty"`

	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The name of the Glacier Vault.
	VaultName string `json:"vaultName,omitempty" yaml:"vaultName,omitempty"`
}
