package types

type Compute_URLMapPathMatcherPathRuleRouteActionCorsPolicy struct {
	// If true, specifies the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Specifies the content for the Access-Control-Expose-Headers header.
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	/*
	   Specifies how long results of a preflight request can be cached in seconds.
	   This translates to the Access-Control-Max-Age header.
	*/
	MaxAge int `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`

	/*
	   In response to a preflight request, setting this to true indicates that the actual request can include user credentials.
	   This translates to the Access-Control-Allow-Credentials header.
	*/
	AllowCredentials bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`

	// Specifies the content for the Access-Control-Allow-Headers header.
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`

	// Specifies the content for the Access-Control-Allow-Methods header.
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`

	/*
	   Specifies the regular expression patterns that match allowed origins. For regular expression grammar
	   please see en.cppreference.com/w/cpp/regex/ecmascript
	   An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.
	*/
	AllowOriginRegexes []string `json:"allowOriginRegexes,omitempty" yaml:"allowOriginRegexes,omitempty"`

	/*
	   Specifies the list of origins that will be allowed to do CORS requests.
	   An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.
	*/
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`
}
