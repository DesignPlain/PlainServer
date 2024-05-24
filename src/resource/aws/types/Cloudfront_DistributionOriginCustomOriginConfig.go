package types

type Cloudfront_DistributionOriginCustomOriginConfig struct {
	// HTTP port the custom origin listens on.
	HttpPort int `json:"httpPort,omitempty" yaml:"httpPort,omitempty"`

	// HTTPS port the custom origin listens on.
	HttpsPort int `json:"httpsPort,omitempty" yaml:"httpsPort,omitempty"`

	// The Custom KeepAlive timeout, in seconds. By default, AWS enforces an upper limit of `60`. But you can request an [increase](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#request-custom-request-timeout). Defaults to `5`.
	OriginKeepaliveTimeout int `json:"originKeepaliveTimeout,omitempty" yaml:"originKeepaliveTimeout,omitempty"`

	// Origin protocol policy to apply to your origin. One of `http-only`, `https-only`, or `match-viewer`.
	OriginProtocolPolicy string `json:"originProtocolPolicy,omitempty" yaml:"originProtocolPolicy,omitempty"`

	// The Custom Read timeout, in seconds. By default, AWS enforces an upper limit of `60`. But you can request an [increase](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#request-custom-request-timeout). Defaults to `30`.
	OriginReadTimeout int `json:"originReadTimeout,omitempty" yaml:"originReadTimeout,omitempty"`

	// SSL/TLS protocols that you want CloudFront to use when communicating with your origin over HTTPS. A list of one or more of `SSLv3`, `TLSv1`, `TLSv1.1`, and `TLSv1.2`.
	OriginSslProtocols []string `json:"originSslProtocols,omitempty" yaml:"originSslProtocols,omitempty"`
}
