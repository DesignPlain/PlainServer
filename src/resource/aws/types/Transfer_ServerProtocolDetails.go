package types

type Transfer_ServerProtocolDetails struct {
	// Use to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket. Valid values: `DEFAULT`, `ENABLE_NO_OP`.
	SetStatOption string `json:"setStatOption,omitempty" yaml:"setStatOption,omitempty"`

	// A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: `DISABLED`, `ENABLED`, `ENFORCED`.
	TlsSessionResumptionMode string `json:"tlsSessionResumptionMode,omitempty" yaml:"tlsSessionResumptionMode,omitempty"`

	// Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
	As2Transports []string `json:"as2Transports,omitempty" yaml:"as2Transports,omitempty"`

	// Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
	PassiveIp string `json:"passiveIp,omitempty" yaml:"passiveIp,omitempty"`
}
