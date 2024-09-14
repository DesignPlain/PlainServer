package accesscontextmanager

type IngressPolicy struct {
	/*
	   The name of the Service Perimeter to add this resource to.


	   - - -
	*/
	IngressPolicyName string `json:"ingressPolicyName,omitempty" yaml:"ingressPolicyName,omitempty"`

	// A GCP resource that is inside of the service perimeter.
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
