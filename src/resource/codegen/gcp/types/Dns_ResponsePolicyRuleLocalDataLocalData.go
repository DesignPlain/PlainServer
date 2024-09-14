package types

type Dns_ResponsePolicyRuleLocalDataLocalData struct {
	// For example, www.example.com.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
	Rrdatas []string `json:"rrdatas,omitempty" yaml:"rrdatas,omitempty"`

	/*
	   Number of seconds that this ResourceRecordSet can be cached by
	   resolvers.
	*/
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	/*
	   One of valid DNS resource types.
	   Possible values are: `A`, `AAAA`, `CAA`, `CNAME`, `DNSKEY`, `DS`, `HTTPS`, `IPSECVPNKEY`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, `TXT`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
