package lightsail

type LbHttpsRedirectionPolicy struct {
	// The Https Redirection state of the load balancer. `true` to activate http to https redirection or `false` to deactivate http to https redirection.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the load balancer to which you want to enable http to https redirection.
	LbName string `json:"lbName,omitempty" yaml:"lbName,omitempty"`
}
