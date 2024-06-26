package types

type Edgecontainer_ClusterAuthorization struct {
	/*
	   User that will be granted the cluster-admin role on the cluster, providing
	   full access to the cluster. Currently, this is a singular field, but will
	   be expanded to allow multiple admins in the future.
	   Structure is documented below.
	*/
	AdminUsers Edgecontainer_ClusterAuthorizationAdminUsers `json:"adminUsers,omitempty" yaml:"adminUsers,omitempty"`
}
