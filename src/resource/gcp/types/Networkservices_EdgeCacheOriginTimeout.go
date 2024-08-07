package types

type Networkservices_EdgeCacheOriginTimeout struct {
	/*
	   The maximum duration to wait for a single origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.
	   Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
	   The connectTimeout capped by the deadline set by the request's maxAttemptsTimeout.  The last connection attempt may have a smaller connectTimeout in order to adhere to the overall maxAttemptsTimeout.
	*/
	ConnectTimeout string `json:"connectTimeout,omitempty" yaml:"connectTimeout,omitempty"`

	/*
	   The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 504 will be returned if the timeout is reached before a response is returned.
	   Defaults to 15 seconds. The timeout must be a value between 1s and 30s.
	   If a failoverOrigin is specified, the maxAttemptsTimeout of the first configured origin sets the deadline for all connection attempts across all failoverOrigins.
	*/
	MaxAttemptsTimeout string `json:"maxAttemptsTimeout,omitempty" yaml:"maxAttemptsTimeout,omitempty"`

	/*
	   The maximum duration to wait between reads of a single HTTP connection/stream.
	   Defaults to 15 seconds.  The timeout must be a value between 1s and 30s.
	   The readTimeout is capped by the responseTimeout.  All reads of the HTTP connection/stream must be completed by the deadline set by the responseTimeout.
	   If the response headers have already been written to the connection, the response will be truncated and logged.

	   <a name="nested_aws_v4_authentication"></a>The `aws_v4_authentication` block supports:
	*/
	ReadTimeout string `json:"readTimeout,omitempty" yaml:"readTimeout,omitempty"`

	/*
	   The maximum duration to wait for the last byte of a response to arrive when reading from the HTTP connection/stream.
	   Defaults to 30 seconds. The timeout must be a value between 1s and 120s.
	   The responseTimeout starts after the connection has been established.
	   This also applies to HTTP Chunked Transfer Encoding responses, and/or when an open-ended Range request is made to the origin. Origins that take longer to write additional bytes to the response than the configured responseTimeout will result in an error being returned to the client.
	   If the response headers have already been written to the connection, the response will be truncated and logged.
	*/
	ResponseTimeout string `json:"responseTimeout,omitempty" yaml:"responseTimeout,omitempty"`
}
