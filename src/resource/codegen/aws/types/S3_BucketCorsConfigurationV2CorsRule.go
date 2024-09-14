package types

type S3_BucketCorsConfigurationV2CorsRule struct {
	// Time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" yaml:"maxAgeSeconds,omitempty"`

	// Set of Headers that are specified in the `Access-Control-Request-Headers` header.
	AllowedHeaders []string `json:"allowedHeaders,omitempty" yaml:"allowedHeaders,omitempty"`

	// Set of HTTP methods that you allow the origin to execute. Valid values are `GET`, `PUT`, `HEAD`, `POST`, and `DELETE`.
	AllowedMethods []string `json:"allowedMethods,omitempty" yaml:"allowedMethods,omitempty"`

	// Set of origins you want customers to be able to access the bucket from.
	AllowedOrigins []string `json:"allowedOrigins,omitempty" yaml:"allowedOrigins,omitempty"`

	// Set of headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
