package types

type Apigatewayv2_getApiCorsConfiguration struct {
	// Whether credentials are included in the CORS request.
	AllowCredentials bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`

	// Set of allowed HTTP headers.
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`

	// Set of allowed HTTP methods.
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`

	// Set of allowed origins.
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`

	// Set of exposed HTTP headers.
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	// Number of seconds that the browser should cache preflight request results.
	MaxAge int `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
}
