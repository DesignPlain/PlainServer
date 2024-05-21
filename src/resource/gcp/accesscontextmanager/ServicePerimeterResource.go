package accesscontextmanager

type ServicePerimeterResource struct {
	/*
	   The name of the Service Perimeter to add this resource to.


	   - - -
	*/
	PerimeterName string `json:"perimeterName,omitempty" yaml:"perimeterName,omitempty"`

	/*
	   A GCP resource that is inside of the service perimeter.
	   Currently only projects are allowed.
	   Format: projects/{project_number}
	*/
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
