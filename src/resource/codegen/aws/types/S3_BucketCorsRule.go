package types

type S3_BucketCorsRule struct {
	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" yaml:"maxAgeSeconds,omitempty"`

	// List of headers allowed.
	AllowedHeaders []string `json:"allowedHeaders,omitempty" yaml:"allowedHeaders,omitempty"`

	// One or more HTTP methods that you allow the origin to execute. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
	AllowedMethods []string `json:"allowedMethods,omitempty" yaml:"allowedMethods,omitempty"`

	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins []string `json:"allowedOrigins,omitempty" yaml:"allowedOrigins,omitempty"`

	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`
}
