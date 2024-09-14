package types

type Gkeonprem_BareMetalClusterSecurityConfigAuthorization struct {
	/*
	   Users that will be granted the cluster-admin role on the cluster, providing full access to the cluster.
	   Structure is documented below.
	*/
	AdminUsers []Gkeonprem_BareMetalClusterSecurityConfigAuthorizationAdminUser `json:"adminUsers,omitempty" yaml:"adminUsers,omitempty"`
}
