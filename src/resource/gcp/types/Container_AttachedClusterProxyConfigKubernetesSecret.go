package types

type Container_AttachedClusterProxyConfigKubernetesSecret struct {
	// Name of the kubernetes secret containing the proxy config.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Namespace of the kubernetes secret containing the proxy config.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
