package types

type S3_BucketCorsRule struct {
	// Specifies expose header in the response.
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" yaml:"maxAgeSeconds,omitempty"`

	// Specifies which headers are allowed.
	AllowedHeaders []string `json:"allowedHeaders,omitempty" yaml:"allowedHeaders,omitempty"`

	// Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
	AllowedMethods []string `json:"allowedMethods,omitempty" yaml:"allowedMethods,omitempty"`

	// Specifies which origins are allowed.
	AllowedOrigins []string `json:"allowedOrigins,omitempty" yaml:"allowedOrigins,omitempty"`
}
