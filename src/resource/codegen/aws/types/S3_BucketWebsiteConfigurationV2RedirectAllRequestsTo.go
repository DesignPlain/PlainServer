package types

type S3_BucketWebsiteConfigurationV2RedirectAllRequestsTo struct {
	// Name of the host where requests are redirected.
	HostName string `json:"hostName,omitempty" yaml:"hostName,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: `http`, `https`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
