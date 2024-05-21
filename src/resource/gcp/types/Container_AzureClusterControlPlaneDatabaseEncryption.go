package types

type Container_AzureClusterControlPlaneDatabaseEncryption struct {
	// The ARM ID of the Azure Key Vault key to encrypt / decrypt data. For example: `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>` Encryption will always take the latest version of the key and hence specific version is not supported.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}
