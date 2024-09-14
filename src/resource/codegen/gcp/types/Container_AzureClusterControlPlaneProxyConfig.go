package types

type Container_AzureClusterControlPlaneProxyConfig struct {
	// The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`
	ResourceGroupId string `json:"resourceGroupId,omitempty" yaml:"resourceGroupId,omitempty"`

	// The URL the of the proxy setting secret with its version. Secret ids are formatted as `https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>`.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`
}
