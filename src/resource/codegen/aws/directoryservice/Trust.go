package directoryservice

type Trust struct {
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName string `json:"remoteDomainName,omitempty" yaml:"remoteDomainName,omitempty"`

	/*
	   Whether to enable selective authentication.
	   Valid values are `Enabled` and `Disabled`.
	   Default value is `Disabled`.
	*/
	SelectiveAuth string `json:"selectiveAuth,omitempty" yaml:"selectiveAuth,omitempty"`

	/*
	   The direction of the Trust relationship.
	   Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	*/
	TrustDirection string `json:"trustDirection,omitempty" yaml:"trustDirection,omitempty"`

	/*
	   Password for the Trust.
	   Does not need to match the passwords for either Directory.
	   Can contain upper- and lower-case letters, numbers, and punctuation characters.
	   May be up to 128 characters long.
	*/
	TrustPassword string `json:"trustPassword,omitempty" yaml:"trustPassword,omitempty"`

	/*
	   Type of the Trust relationship.
	   Valid values are `Forest` and `External`.
	   Default value is `Forest`.
	*/
	TrustType string `json:"trustType,omitempty" yaml:"trustType,omitempty"`

	/*
	   Set of IPv4 addresses for the DNS server associated with the remote Directory.
	   Can contain between 1 and 4 values.
	*/
	ConditionalForwarderIpAddrs []string `json:"conditionalForwarderIpAddrs,omitempty" yaml:"conditionalForwarderIpAddrs,omitempty"`

	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder bool `json:"deleteAssociatedConditionalForwarder,omitempty" yaml:"deleteAssociatedConditionalForwarder,omitempty"`

	// ID of the Directory.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`
}
