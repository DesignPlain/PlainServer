package types

type Redshift_getDataSharesDataShare struct {
	// Identifier of a datashare to show its managing entity.
	ManagedBy string `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`

	// ARN (Amazon Resource Name) of the producer.
	ProducerArn string `json:"producerArn,omitempty" yaml:"producerArn,omitempty"`

	// ARN (Amazon Resource Name) of the data share.
	DataShareArn string `json:"dataShareArn,omitempty" yaml:"dataShareArn,omitempty"`
}
