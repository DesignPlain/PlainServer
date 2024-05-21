package compute

type TargetHttpsProxy struct {
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
	   A URL referring to a networksecurity.ServerTlsPolicy
	   resource that describes how the proxy should authenticate inbound
	   traffic. serverTlsPolicy only applies to a global TargetHttpsProxy
	   attached to globalForwardingRules with the loadBalancingScheme
	   set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.
	   For details which ServerTlsPolicy resources are accepted with
	   INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED
	   loadBalancingScheme consult ServerTlsPolicy documentation.
	   If left blank, communications are not encrypted.
	*/
	ServerTlsPolicy string `json:"serverTlsPolicy,omitempty" yaml:"serverTlsPolicy,omitempty"`

	/*
	   URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.
	   Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	   sslCertificates and certificateManagerCertificates can not be defined together.
	*/
	SslCertificates []string `json:"sslCertificates,omitempty" yaml:"sslCertificates,omitempty"`

	/*
	   A reference to the SslPolicy resource that will be associated with
	   the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	   resource will not have any SSL policy configured.
	*/
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`

	/*
	   A reference to the UrlMap resource that defines the mapping from URL
	   to the BackendService.


	   - - -
	*/
	UrlMap string `json:"urlMap,omitempty" yaml:"urlMap,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies how long to keep a connection open, after completing a response,
	   while there is no matching traffic (in seconds). If an HTTP keepalive is
	   not specified, a default value (610 seconds) will be used. For Global
	   external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	   the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	   load balancer (classic), this option is not available publicly.
	*/
	HttpKeepAliveTimeoutSec int `json:"httpKeepAliveTimeoutSec,omitempty" yaml:"httpKeepAliveTimeoutSec,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   This field only applies when the forwarding rule that references
	   this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	*/
	ProxyBind bool `json:"proxyBind,omitempty" yaml:"proxyBind,omitempty"`

	/*
	   Specifies the QUIC override policy for this resource. This determines
	   whether the load balancer will attempt to negotiate QUIC with clients
	   or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	   specified, Google manages whether QUIC is used.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	*/
	QuicOverride string `json:"quicOverride,omitempty" yaml:"quicOverride,omitempty"`

	/*
	   URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
	   Currently, you may specify up to 15 certificates. Certificate manager certificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	   sslCertificates and certificateManagerCertificates fields can not be defined together.
	   Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName}` or just the self_link `projects/{project}/locations/{location}/certificates/{resourceName}`
	*/
	CertificateManagerCertificates []string `json:"certificateManagerCertificates,omitempty" yaml:"certificateManagerCertificates,omitempty"`

	/*
	   A reference to the CertificateMap resource uri that identifies a certificate map
	   associated with the given target proxy. This field can only be set for global target proxies.
	   Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	*/
	CertificateMap string `json:"certificateMap,omitempty" yaml:"certificateMap,omitempty"`
}
