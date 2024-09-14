package types

type Gkeonprem_BareMetalAdminClusterSecurityConfig struct {
	/*
	   Configures user access to the Bare Metal User cluster.
	   Structure is documented below.
	*/
	Authorization Gkeonprem_BareMetalAdminClusterSecurityConfigAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`
}
