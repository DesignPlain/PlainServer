package projects

type Service struct {
	// The service to enable.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   If `true`, services that are enabled
	   and which depend on this service should also be disabled when this service is
	   destroyed. If `false` or unset, an error will be generated if any enabled
	   services depend on this service when destroying it.
	*/
	DisableDependentServices bool `json:"disableDependentServices,omitempty" yaml:"disableDependentServices,omitempty"`

	// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy bool `json:"disableOnDestroy,omitempty" yaml:"disableOnDestroy,omitempty"`

	/*
	   The project ID. If not provided, the provider project
	   is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
