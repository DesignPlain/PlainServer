package types

type Lambda_FunctionUrlCors struct {
	// Whether to allow cookies or other credentials in requests to the function URL. The default is `false`.
	AllowCredentials bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`

	// The HTTP headers that origins can include in requests to the function URL. For example: `["date", "keep-alive", "x-custom-header"]`.
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`

	// The HTTP methods that are allowed when calling the function URL. For example: `["GET", "POST", "DELETE"]`, or the wildcard character (`["-"]`).
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`

	// The origins that can access the function URL. You can list any number of specific origins (or the wildcard character (`"-"`)), separated by a comma. For example: `["https://www.example.com", "http://localhost:60905"]`.
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`

	// The HTTP headers in your function response that you want to expose to origins that call the function URL.
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	// The maximum amount of time, in seconds, that web browsers can cache results of a preflight request. By default, this is set to `0`, which means that the browser doesn't cache results. The maximum value is `86400`.
	MaxAge int `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
}
