package types

type Gkeonprem_BareMetalAdminClusterSecurityConfigAuthorization struct {
	/*
	   Users that will be granted the cluster-admin role on the cluster, providing full access to the cluster.
	   Structure is documented below.
	*/
	AdminUsers []Gkeonprem_BareMetalAdminClusterSecurityConfigAuthorizationAdminUser `json:"adminUsers,omitempty" yaml:"adminUsers,omitempty"`
}
