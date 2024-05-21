package types

type Accesscontextmanager_ServicePerimetersServicePerimeterStatusVpcAccessibleServices struct {
	/*
	   The list of APIs usable within the Service Perimeter.
	   Must be empty unless `enableRestriction` is True.
	*/
	AllowedServices []string `json:"allowedServices,omitempty" yaml:"allowedServices,omitempty"`

	/*
	   Whether to restrict API calls within the Service Perimeter to the
	   list of APIs specified in 'allowedServices'.
	*/
	EnableRestriction bool `json:"enableRestriction,omitempty" yaml:"enableRestriction,omitempty"`
}
