package types

type Redshift_getProducerDataSharesDataShare struct {
	/*
	   Amazon Resource Name (ARN) of the producer namespace that returns in the list of datashares.

	   The following arguments are optional:
	*/
	ProducerArn string `json:"producerArn,omitempty" yaml:"producerArn,omitempty"`

	// ARN (Amazon Resource Name) of the data share.
	DataShareArn string `json:"dataShareArn,omitempty" yaml:"dataShareArn,omitempty"`

	// Identifier of a datashare to show its managing entity.
	ManagedBy string `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
}
