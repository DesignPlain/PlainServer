package types

type Gkeonprem_VMwareClusterLoadBalancerManualLbConfig struct {
	/*
	   NodePort for konnectivity server service running as a sidecar in each
	   kube-apiserver pod (ex. 30564).
	*/
	KonnectivityServerNodePort int `json:"konnectivityServerNodePort,omitempty" yaml:"konnectivityServerNodePort,omitempty"`

	/*
	   NodePort for control plane service. The Kubernetes API server in the admin
	   cluster is implemented as a Service of type NodePort (ex. 30968).
	*/
	ControlPlaneNodePort int `json:"controlPlaneNodePort,omitempty" yaml:"controlPlaneNodePort,omitempty"`

	/*
	   NodePort for ingress service's http. The ingress service in the admin
	   cluster is implemented as a Service of type NodePort (ex. 32527).
	*/
	IngressHttpNodePort int `json:"ingressHttpNodePort,omitempty" yaml:"ingressHttpNodePort,omitempty"`

	/*
	   NodePort for ingress service's https. The ingress service in the admin
	   cluster is implemented as a Service of type NodePort (ex. 30139).
	*/
	IngressHttpsNodePort int `json:"ingressHttpsNodePort,omitempty" yaml:"ingressHttpsNodePort,omitempty"`
}
