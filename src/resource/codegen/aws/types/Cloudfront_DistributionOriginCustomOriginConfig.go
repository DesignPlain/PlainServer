package types

type Cloudfront_DistributionOriginCustomOriginConfig struct {
	// Origin protocol policy to apply to your origin. One of `http-only`, `https-only`, or `match-viewer`.
	OriginProtocolPolicy string `json:"originProtocolPolicy,omitempty" yaml:"originProtocolPolicy,omitempty"`

	// The Custom Read timeout, in seconds. By default, AWS enforces an upper limit of `60`. But you can request an [increase](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#request-custom-request-timeout). Defaults to `30`.
	OriginReadTimeout int `json:"originReadTimeout,omitempty" yaml:"originReadTimeout,omitempty"`

	// List of SSL/TLS protocols that CloudFront can use when connecting to your origin over HTTPS. Valid values: `SSLv3`, `TLSv1`, `TLSv1.1`, `TLSv1.2`. For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the Amazon CloudFront Developer Guide.
	OriginSslProtocols []string `json:"originSslProtocols,omitempty" yaml:"originSslProtocols,omitempty"`

	// HTTP port the custom origin listens on.
	HttpPort int `json:"httpPort,omitempty" yaml:"httpPort,omitempty"`

	// HTTPS port the custom origin listens on.
	HttpsPort int `json:"httpsPort,omitempty" yaml:"httpsPort,omitempty"`

	// The Custom KeepAlive timeout, in seconds. By default, AWS enforces an upper limit of `60`. But you can request an [increase](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#request-custom-request-timeout). Defaults to `5`.
	OriginKeepaliveTimeout int `json:"originKeepaliveTimeout,omitempty" yaml:"originKeepaliveTimeout,omitempty"`
}
