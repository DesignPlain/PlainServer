package types

type Container_AttachedClusterProxyConfig struct {
	/*
	   The Kubernetes Secret resource that contains the HTTP(S) proxy configuration.
	   Structure is documented below.
	*/
	KubernetesSecret Container_AttachedClusterProxyConfigKubernetesSecret `json:"kubernetesSecret,omitempty" yaml:"kubernetesSecret,omitempty"`
}
