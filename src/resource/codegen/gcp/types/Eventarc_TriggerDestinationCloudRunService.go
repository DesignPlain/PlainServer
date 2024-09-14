package types

type Eventarc_TriggerDestinationCloudRunService struct {
	// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Required. The region the Cloud Run service is deployed in.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Required. The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project of the trigger object can be addressed.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
