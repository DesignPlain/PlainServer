package route53domains

import types "libds/aws/types"

type DelegationSignerRecord struct {
	// The name of the domain that will have its parent DNS zone updated with the Delegation Signer record.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The information about a key, including the algorithm, public key-value, and flags.
	SigningAttributes types.Route53domains_DelegationSignerRecordSigningAttributes `json:"signingAttributes,omitempty" yaml:"signingAttributes,omitempty"`

	//
	Timeouts types.Route53domains_DelegationSignerRecordTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
