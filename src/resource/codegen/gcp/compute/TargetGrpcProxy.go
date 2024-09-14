package compute

type TargetGrpcProxy struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource
	   is created. The name must be 1-63 characters long, and comply
	   with RFC1035. Specifically, the name must be 1-63 characters long
	   and match the regular expression `a-z?` which
	   means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   URL to the UrlMap resource that defines the mapping from URL to
	   the BackendService. The protocol field in the BackendService
	   must be set to GRPC.
	*/
	UrlMap string `json:"urlMap,omitempty" yaml:"urlMap,omitempty"`

	/*
	   If true, indicates that the BackendServices referenced by
	   the urlMap may be accessed by gRPC applications without using
	   a sidecar proxy. This will enable configuration checks on urlMap
	   and its referenced BackendServices to not allow unsupported features.
	   A gRPC application must use "xds:///" scheme in the target URI
	   of the service it is connecting to. If false, indicates that the
	   BackendServices referenced by the urlMap will be accessed by gRPC
	   applications via a sidecar proxy. In this case, a gRPC application
	   must not use "xds:///" scheme in the target URI of the service
	   it is connecting to
	*/
	ValidateForProxyless bool `json:"validateForProxyless,omitempty" yaml:"validateForProxyless,omitempty"`
}
