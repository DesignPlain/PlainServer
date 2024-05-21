package types

type Looker_InstanceCustomDomain struct {
	// Domain name
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	/*
	   (Output)
	   Status of the custom domain.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
