package types

type Gkeonprem_BareMetalClusterSecurityConfig struct {
	/*
	   Configures user access to the Bare Metal User cluster.
	   Structure is documented below.
	*/
	Authorization Gkeonprem_BareMetalClusterSecurityConfigAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`
}
