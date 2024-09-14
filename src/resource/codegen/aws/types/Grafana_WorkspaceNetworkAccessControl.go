package types

type Grafana_WorkspaceNetworkAccessControl struct {
	// An array of prefix list IDs.
	PrefixListIds []string `json:"prefixListIds,omitempty" yaml:"prefixListIds,omitempty"`

	// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	VpceIds []string `json:"vpceIds,omitempty" yaml:"vpceIds,omitempty"`
}
