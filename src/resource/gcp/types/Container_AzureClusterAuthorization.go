package types

type Container_AzureClusterAuthorization struct {
	// Groups of users that can perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole to the groups. Up to ten admin groups can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminGroups []Container_AzureClusterAuthorizationAdminGroup `json:"adminGroups,omitempty" yaml:"adminGroups,omitempty"`

	// Users that can perform operations as a cluster admin. A new ClusterRoleBinding will be created to grant the cluster-admin ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminUsers []Container_AzureClusterAuthorizationAdminUser `json:"adminUsers,omitempty" yaml:"adminUsers,omitempty"`
}
