package redshift

type DataShareConsumerAssociation struct {
	// Whether to allow write operations for a datashare.
	AllowWrites bool `json:"allowWrites,omitempty" yaml:"allowWrites,omitempty"`

	// Whether the datashare is associated with the entire account. Conflicts with `consumer_arn` and `consumer_region`.
	AssociateEntireAccount bool `json:"associateEntireAccount,omitempty" yaml:"associateEntireAccount,omitempty"`

	// Amazon Resource Name (ARN) of the consumer that is associated with the datashare. Conflicts with `associate_entire_account` and `consumer_region`.
	ConsumerArn string `json:"consumerArn,omitempty" yaml:"consumerArn,omitempty"`

	// From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified AWS Region. Conflicts with `associate_entire_account` and `consumer_arn`.
	ConsumerRegion string `json:"consumerRegion,omitempty" yaml:"consumerRegion,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the datashare that the consumer is to use with the account or the namespace.

	   The following arguments are optional:
	*/
	DataShareArn string `json:"dataShareArn,omitempty" yaml:"dataShareArn,omitempty"`
}
