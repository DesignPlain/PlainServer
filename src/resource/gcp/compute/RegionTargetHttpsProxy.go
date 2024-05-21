package compute

type RegionTargetHttpsProxy struct {
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

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The Region in which the created target https proxy should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.
	   At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
	   sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	*/
	SslCertificates []string `json:"sslCertificates,omitempty" yaml:"sslCertificates,omitempty"`

	/*
	   A reference to the Region SslPolicy resource that will be associated with
	   the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	   resource will not have any SSL policy configured.
	*/
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`

	/*
	   A reference to the RegionUrlMap resource that defines the mapping from URL
	   to the RegionBackendService.


	   - - -
	*/
	UrlMap string `json:"urlMap,omitempty" yaml:"urlMap,omitempty"`

	/*
	   URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
	   Currently, you may specify up to 15 certificates. Certificate manager certificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	   sslCertificates and certificateManagerCertificates fields can not be defined together.
	   Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName}` or just the self_link `projects/{project}/locations/{location}/certificates/{resourceName}`
	*/
	CertificateManagerCertificates []string `json:"certificateManagerCertificates,omitempty" yaml:"certificateManagerCertificates,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
