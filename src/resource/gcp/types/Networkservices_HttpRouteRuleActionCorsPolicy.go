package types

type Networkservices_HttpRouteRuleActionCorsPolicy struct {
	// Specifies the content for Access-Control-Allow-Methods header.
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`

	// Specifies the regular expression patterns that match allowed origins.
	AllowOriginRegexes []string `json:"allowOriginRegexes,omitempty" yaml:"allowOriginRegexes,omitempty"`

	// Specifies the list of origins that will be allowed to do CORS requests.
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`

	/*
	   If true, the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.

	   - - -
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Specifies the content for Access-Control-Expose-Headers header.
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	// Specifies how long result of a preflight request can be cached in seconds.
	MaxAge string `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`

	// In response to a preflight request, setting this to true indicates that the actual request can include user credentials.
	AllowCredentials bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`

	// Specifies the content for Access-Control-Allow-Headers header.
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`
}
