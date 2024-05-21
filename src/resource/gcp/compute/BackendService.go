package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type BackendService struct {
	/*
	   Headers that the HTTP/S load balancer should add to proxied
	   requests.
	*/
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty" yaml:"customRequestHeaders,omitempty"`

	// If true, enable Cloud CDN for this BackendService.
	EnableCdn bool `json:"enableCdn,omitempty" yaml:"enableCdn,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   How many seconds to wait for the backend before considering it a
	   failed request. Default is 30 seconds. Valid range is [1, 86400].
	*/
	TimeoutSec int `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`

	/*
	   Consistent Hash-based load balancing can be used to provide soft session
	   affinity based on HTTP headers, cookies or other properties. This load balancing
	   policy is applicable only for HTTP connections. The affinity to a particular
	   destination host will be lost when one or more hosts are added/removed from the
	   destination service. This field specifies parameters that control consistent
	   hashing. This field only applies if the load_balancing_scheme is set to
	   INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is
	   set to MAGLEV or RING_HASH.
	   Structure is documented below.
	*/
	ConsistentHash types.Compute_BackendServiceConsistentHash `json:"consistentHash,omitempty" yaml:"consistentHash,omitempty"`

	/*
	   The security settings that apply to this backend service. This field is applicable to either
	   a regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and
	   load_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the
	   load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	   Structure is documented below.
	*/
	SecuritySettings types.Compute_BackendServiceSecuritySettings `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`

	/*
	   Cloud CDN configuration for this BackendService.
	   Structure is documented below.
	*/
	CdnPolicy types.Compute_BackendServiceCdnPolicy `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`

	/*
	   Headers that the HTTP/S load balancer should add to proxied
	   responses.
	*/
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty"`

	/*
	   A list of locality load balancing policies to be used in order of
	   preference. Either the policy or the customPolicy field should be set.
	   Overrides any value set in the localityLbPolicy field.
	   localityLbPolicies is only supported when the BackendService is referenced
	   by a URL Map that is referenced by a target gRPC proxy that has the
	   validateForProxyless field set to true.
	   Structure is documented below.
	*/
	LocalityLbPolicies []types.Compute_BackendServiceLocalityLbPolicy `json:"localityLbPolicies,omitempty" yaml:"localityLbPolicies,omitempty"`

	/*
	   Time for which instance will be drained (not accept new
	   connections, but still work to finish started).
	*/
	ConnectionDrainingTimeoutSec int `json:"connectionDrainingTimeoutSec,omitempty" yaml:"connectionDrainingTimeoutSec,omitempty"`

	// The resource URL for the edge security policy associated with this backend service.
	EdgeSecurityPolicy string `json:"edgeSecurityPolicy,omitempty" yaml:"edgeSecurityPolicy,omitempty"`

	/*
	   The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
	   for health checking this BackendService. Currently at most one health
	   check can be specified.
	   A health check must be specified unless the backend service uses an internet
	   or serverless NEG as a backend.
	   For internal load balancing, a URL to a HealthCheck resource must be specified instead.
	*/
	HealthChecks string `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`

	/*
	   The load balancing algorithm used within the scope of the locality.
	   The possible values are:
	*/
	LocalityLbPolicy string `json:"localityLbPolicy,omitempty" yaml:"localityLbPolicy,omitempty"`

	/*
	   Lifetime of cookies in seconds if session_affinity is
	   GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	   only until the end of the browser session (or equivalent). The
	   maximum allowed value for TTL is one day.
	   When the load balancing scheme is INTERNAL, this field is not used.
	*/
	AffinityCookieTtlSec int `json:"affinityCookieTtlSec,omitempty" yaml:"affinityCookieTtlSec,omitempty"`

	/*
	   Settings controlling eviction of unhealthy hosts from the load balancing pool.
	   Applicable backend service types can be a global backend service with the
	   loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.
	   Structure is documented below.
	*/
	OutlierDetection types.Compute_BackendServiceOutlierDetection `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`

	/*
	   Name of backend port. The same name should appear in the instance
	   groups referenced by this service. Required when the load balancing
	   scheme is EXTERNAL.
	*/
	PortName string `json:"portName,omitempty" yaml:"portName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The security policy associated with this backend service.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	/*
	   The protocol this BackendService uses to communicate with backends.
	   The default is HTTP. --NOTE--: HTTP2 is only valid for beta HTTP/2 load balancer
	   types and may result in errors if used with the GA API. --NOTE--: With protocol “UNSPECIFIED”,
	   the backend service can be used by Layer 4 Internal Load Balancing or Network Load Balancing
	   with TCP/UDP/L3_DEFAULT Forwarding Rule protocol.
	   Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `TCP`, `SSL`, `GRPC`, `UNSPECIFIED`.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	   Possible values are: `AUTOMATIC`, `DISABLED`.
	*/
	CompressionMode string `json:"compressionMode,omitempty" yaml:"compressionMode,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Settings for enabling Cloud Identity Aware Proxy
	   Structure is documented below.
	*/
	Iap types.Compute_BackendServiceIap `json:"iap,omitempty" yaml:"iap,omitempty"`

	/*
	   This field denotes the logging options for the load balancer traffic served by this backend service.
	   If logging is enabled, logs will be exported to Stackdriver.
	   Structure is documented below.
	*/
	LogConfig types.Compute_BackendServiceLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	/*
	   Type of session affinity to use. The default is NONE. Session affinity is
	   not applicable if the protocol is UDP.
	   Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`.
	*/
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`

	/*
	   The set of backends that serve this BackendService.
	   Structure is documented below.
	*/
	Backends []types.Compute_BackendServiceBackend `json:"backends,omitempty" yaml:"backends,omitempty"`

	/*
	   Settings controlling the volume of connections to a backend service. This field
	   is applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
	   Structure is documented below.
	*/
	CircuitBreakers types.Compute_BackendServiceCircuitBreakers `json:"circuitBreakers,omitempty" yaml:"circuitBreakers,omitempty"`

	/*
	   Indicates whether the backend service will be used with internal or
	   external load balancing. A backend service created for one type of
	   load balancing cannot be used with the other. For more information, refer to
	   [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
	   Default value is `EXTERNAL`.
	   Possible values are: `EXTERNAL`, `INTERNAL_SELF_MANAGED`, `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.
	*/
	LoadBalancingScheme string `json:"loadBalancingScheme,omitempty" yaml:"loadBalancingScheme,omitempty"`
}
