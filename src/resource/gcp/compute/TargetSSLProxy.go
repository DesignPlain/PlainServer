package compute

type TargetSSLProxy struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies the type of proxy header to append before sending data to
	   the backend.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `PROXY_V1`.
	*/
	ProxyHeader string `json:"proxyHeader,omitempty" yaml:"proxyHeader,omitempty"`

	/*
	   A list of SslCertificate resources that are used to authenticate
	   connections between users and the load balancer. At least one
	   SSL certificate must be specified.
	*/
	SslCertificates []string `json:"sslCertificates,omitempty" yaml:"sslCertificates,omitempty"`

	/*
	   A reference to the SslPolicy resource that will be associated with
	   the TargetSslProxy resource. If not set, the TargetSslProxy
	   resource will not have any SSL policy configured.
	*/
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`

	/*
	   A reference to the BackendService resource.


	   - - -
	*/
	BackendService string `json:"backendService,omitempty" yaml:"backendService,omitempty"`

	/*
	   A reference to the CertificateMap resource uri that identifies a certificate map
	   associated with the given target proxy. This field can only be set for global target proxies.
	   Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	*/
	CertificateMap string `json:"certificateMap,omitempty" yaml:"certificateMap,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
