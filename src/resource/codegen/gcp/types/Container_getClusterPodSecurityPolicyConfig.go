package types

type Container_getClusterPodSecurityPolicyConfig struct {
	// Enable the PodSecurityPolicy controller for this cluster. If enabled, pods must be valid under a PodSecurityPolicy to be created.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
