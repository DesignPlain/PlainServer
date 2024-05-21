package types

type Gkeonprem_VMwareClusterAuthorization struct {
	/*
	   Users that will be granted the cluster-admin role on the cluster, providing
	   full access to the cluster.
	   Structure is documented below.
	*/
	AdminUsers []Gkeonprem_VMwareClusterAuthorizationAdminUser `json:"adminUsers,omitempty" yaml:"adminUsers,omitempty"`
}
