package types

type Lightsail_DistributionCacheBehaviorSettings struct {
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	MaximumTtl int `json:"maximumTtl,omitempty" yaml:"maximumTtl,omitempty"`

	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	MinimumTtl int `json:"minimumTtl,omitempty" yaml:"minimumTtl,omitempty"`

	// The HTTP methods that are processed and forwarded to the distribution's origin.
	AllowedHttpMethods string `json:"allowedHttpMethods,omitempty" yaml:"allowedHttpMethods,omitempty"`

	// The HTTP method responses that are cached by your distribution.
	CachedHttpMethods string `json:"cachedHttpMethods,omitempty" yaml:"cachedHttpMethods,omitempty"`

	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	DefaultTtl int `json:"defaultTtl,omitempty" yaml:"defaultTtl,omitempty"`

	// An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded. Detailed below
	ForwardedCookies Lightsail_DistributionCacheBehaviorSettingsForwardedCookies `json:"forwardedCookies,omitempty" yaml:"forwardedCookies,omitempty"`

	// An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded. Detailed below
	ForwardedHeaders Lightsail_DistributionCacheBehaviorSettingsForwardedHeaders `json:"forwardedHeaders,omitempty" yaml:"forwardedHeaders,omitempty"`

	// An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded. Detailed below
	ForwardedQueryStrings Lightsail_DistributionCacheBehaviorSettingsForwardedQueryStrings `json:"forwardedQueryStrings,omitempty" yaml:"forwardedQueryStrings,omitempty"`
}
