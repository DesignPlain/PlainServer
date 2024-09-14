package accesscontextmanager

type EgressPolicy struct {
	/*
	   The name of the Service Perimeter to add this resource to.


	   - - -
	*/
	EgressPolicyName string `json:"egressPolicyName,omitempty" yaml:"egressPolicyName,omitempty"`

	// A GCP resource that is inside of the service perimeter.
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
