package networkservices

import types "DesignSphere_Server/src/resource/gcp/types"

type EdgeCacheService struct {
	// A human-readable description of the hostRule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Disables HTTP/2.
	   HTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.
	   Some legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.
	*/
	DisableHttp2 bool `json:"disableHttp2,omitempty" yaml:"disableHttp2,omitempty"`

	// HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
	DisableQuic bool `json:"disableQuic,omitempty" yaml:"disableQuic,omitempty"`

	// Resource URL that points at the Cloud Armor edge security policy that is applied on each request against the EdgeCacheService.
	EdgeSecurityPolicy string `json:"edgeSecurityPolicy,omitempty" yaml:"edgeSecurityPolicy,omitempty"`

	/*
	   Set of label tags associated with the EdgeCache resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Specifies the logging options for the traffic served by this service. If logging is enabled, logs will be exported to Cloud Logging.
	   Structure is documented below.
	*/
	LogConfig types.Networkservices_EdgeCacheServiceLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Require TLS (HTTPS) for all clients connecting to this service.
	   Clients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).
	   You must have at least one (1) edgeSslCertificate specified to enable this.
	*/
	RequireTls bool `json:"requireTls,omitempty" yaml:"requireTls,omitempty"`

	/*
	   URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.
	   Note that only "global" certificates with a "scope" of "EDGE_CACHE" can be attached to an EdgeCacheService.
	*/
	EdgeSslCertificates []string `json:"edgeSslCertificates,omitempty" yaml:"edgeSslCertificates,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is created.
	   The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]- which means the first character must be a letter,
	   and all following characters must be a dash, underscore, letter or digit.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Defines how requests are routed, modified, cached and/or which origin content is filled from.
	   Structure is documented below.
	*/
	Routing types.Networkservices_EdgeCacheServiceRouting `json:"routing,omitempty" yaml:"routing,omitempty"`

	/*
	   URL of the SslPolicy resource that will be associated with the EdgeCacheService.
	   If not set, the EdgeCacheService has no SSL policy configured, and will default to the "COMPATIBLE" policy.
	*/
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`
}
