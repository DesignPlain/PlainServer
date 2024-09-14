package vmwareengine

type ExternalAddress struct {
	// User-provided description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The internal IP address of a workload VM.
	InternalIp string `json:"internalIp,omitempty" yaml:"internalIp,omitempty"`

	/*
	   The ID of the external IP Address.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The resource name of the private cloud to create a new external address in.
	   Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	   For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
