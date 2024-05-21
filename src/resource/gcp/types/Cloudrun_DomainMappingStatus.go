package types

type Cloudrun_DomainMappingStatus struct {
	/*
	   The resource records required to configure this domain mapping. These
	   records must be added to the domain's DNS configuration in order to
	   serve the application via this domain mapping.
	   Structure is documented below.
	*/
	ResourceRecords []Cloudrun_DomainMappingStatusResourceRecord `json:"resourceRecords,omitempty" yaml:"resourceRecords,omitempty"`

	/*
	   (Output)
	   Array of observed DomainMappingConditions, indicating the current state
	   of the DomainMapping.
	   Structure is documented below.
	*/
	Conditions []Cloudrun_DomainMappingStatusCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	/*
	   (Output)
	   The name of the route that the mapping currently points to.
	*/
	MappedRouteName string `json:"mappedRouteName,omitempty" yaml:"mappedRouteName,omitempty"`

	/*
	   (Output)
	   ObservedGeneration is the 'Generation' of the DomainMapping that
	   was last processed by the controller.
	*/
	ObservedGeneration int `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`
}
