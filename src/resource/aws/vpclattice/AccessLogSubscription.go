package vpclattice

type AccessLogSubscription struct {
	// Amazon Resource Name (ARN) of the log destination.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// The ID or Amazon Resource Identifier (ARN) of the service network or service. You must use the ARN if the resources specified in the operation are in different accounts.
	ResourceIdentifier string `json:"resourceIdentifier,omitempty" yaml:"resourceIdentifier,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
