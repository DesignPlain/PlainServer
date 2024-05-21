package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type RegionBackendService struct {
	/*
	   The load balancing algorithm used within the scope of the locality.
	   The possible values are:
	*/
	LocalityLbPolicy string `json:"localityLbPolicy,omitempty" yaml:"localityLbPolicy,omitempty"`

	/*
	   The protocol this RegionBackendService uses to communicate with backends.
	   The default is HTTP. --NOTE--: HTTP2 is only valid for beta HTTP/2 load balancer
	   types and may result in errors if used with the GA API.
	   Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, `UNSPECIFIED`.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   Type of session affinity to use. The default is NONE. Session affinity is
	   not applicable if the protocol is UDP.
	   Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`, `CLIENT_IP_NO_DESTINATION`.
	*/
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`

	/*
	   Connection Tracking configuration for this BackendService.
	   This is available only for Layer 4 Internal Load Balancing and
	   Network Load Balancing.
	   Structure is documented below.
	*/
	ConnectionTrackingPolicy types.Compute_RegionBackendServiceConnectionTrackingPolicy `json:"connectionTrackingPolicy,omitempty" yaml:"connectionTrackingPolicy,omitempty"`

	// If true, enable Cloud CDN for this RegionBackendService.
	EnableCdn bool `json:"enableCdn,omitempty" yaml:"enableCdn,omitempty"`

	/*
	   The set of URLs to HealthCheck resources for health checking
	   this RegionBackendService. Currently at most one health
	   check can be specified.
	   A health check must be specified unless the backend service uses an internet
	   or serverless NEG as a backend.
	*/
	HealthChecks string `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`

	/*
	   Settings for enabling Cloud Identity Aware Proxy
	   Structure is documented below.
	*/
	Iap types.Compute_RegionBackendServiceIap `json:"iap,omitempty" yaml:"iap,omitempty"`

	/*
	   Indicates what kind of load balancing this regional backend service
	   will be used for. A backend service created for one type of load
	   balancing cannot be used with the other(s). For more information, refer to
	   [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
	   Default value is `INTERNAL`.
	   Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL`, `INTERNAL_MANAGED`.
	*/
	LoadBalancingScheme string `json:"loadBalancingScheme,omitempty" yaml:"loadBalancingScheme,omitempty"`

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
	   Settings controlling eviction of unhealthy hosts from the load balancing pool.
	   This field is applicable only when the `load_balancing_scheme` is set
	   to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
	   Structure is documented below.
	*/
	OutlierDetection types.Compute_RegionBackendServiceOutlierDetection `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`

	/*
	   A named port on a backend instance group representing the port for
	   communication to the backend VMs in that group. Required when the
	   loadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
	   and the backends are instance groups. The named port must be defined on each
	   backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
	   default of "http" if not given.
	   Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
	*/
	PortName string `json:"portName,omitempty" yaml:"portName,omitempty"`

	/*
	   Lifetime of cookies in seconds if session_affinity is
	   GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	   only until the end of the browser session (or equivalent). The
	   maximum allowed value for TTL is one day.
	   When the load balancing scheme is INTERNAL, this field is not used.
	*/
	AffinityCookieTtlSec int `json:"affinityCookieTtlSec,omitempty" yaml:"affinityCookieTtlSec,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The security policy associated with this backend service.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	/*
	   Policy for failovers.
	   Structure is documented below.
	*/
	FailoverPolicy types.Compute_RegionBackendServiceFailoverPolicy `json:"failoverPolicy,omitempty" yaml:"failoverPolicy,omitempty"`

	/*
	   The Region in which the created backend service should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   How many seconds to wait for the backend before considering it a
	   failed request. Default is 30 seconds. Valid range is [1, 86400].
	*/
	TimeoutSec int `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`

	/*
	   The set of backends that serve this RegionBackendService.
	   Structure is documented below.
	*/
	Backends []types.Compute_RegionBackendServiceBackend `json:"backends,omitempty" yaml:"backends,omitempty"`

	/*
	   Consistent Hash-based load balancing can be used to provide soft session
	   affinity based on HTTP headers, cookies or other properties. This load balancing
	   policy is applicable only for HTTP connections. The affinity to a particular
	   destination host will be lost when one or more hosts are added/removed from the
	   destination service. This field specifies parameters that control consistent
	   hashing.
	   This field only applies when all of the following are true -
	*/
	ConsistentHash types.Compute_RegionBackendServiceConsistentHash `json:"consistentHash,omitempty" yaml:"consistentHash,omitempty"`

	/*
	   Time for which instance will be drained (not accept new
	   connections, but still work to finish started).
	*/
	ConnectionDrainingTimeoutSec int `json:"connectionDrainingTimeoutSec,omitempty" yaml:"connectionDrainingTimeoutSec,omitempty"`

	/*
	   This field denotes the logging options for the load balancer traffic served by this backend service.
	   If logging is enabled, logs will be exported to Stackdriver.
	   Structure is documented below.
	*/
	LogConfig types.Compute_RegionBackendServiceLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	/*
	   The URL of the network to which this backend service belongs.
	   This field can only be specified when the load balancing scheme is set to INTERNAL.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S) load balancing.
	   Structure is documented below.
	*/
	Subsetting types.Compute_RegionBackendServiceSubsetting `json:"subsetting,omitempty" yaml:"subsetting,omitempty"`

	/*
	   Cloud CDN configuration for this BackendService.
	   Structure is documented below.
	*/
	CdnPolicy types.Compute_RegionBackendServiceCdnPolicy `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`

	/*
	   Settings controlling the volume of connections to a backend service. This field
	   is applicable only when the `load_balancing_scheme` is set to INTERNAL_MANAGED
	   and the `protocol` is set to HTTP, HTTPS, or HTTP2.
	   Structure is documented below.
	*/
	CircuitBreakers types.Compute_RegionBackendServiceCircuitBreakers `json:"circuitBreakers,omitempty" yaml:"circuitBreakers,omitempty"`
}
