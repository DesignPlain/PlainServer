package types

type Edgecontainer_ClusterSystemAddonsConfig struct {
	/*
	   Config for the Ingress add-on which allows customers to create an Ingress
	   object to manage external access to the servers in a cluster. The add-on
	   consists of istiod and istio-ingress.
	   Structure is documented below.
	*/
	Ingress Edgecontainer_ClusterSystemAddonsConfigIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`
}
