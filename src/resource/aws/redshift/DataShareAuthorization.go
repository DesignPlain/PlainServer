package redshift

type DataShareAuthorization struct {
	// Whether to allow write operations for a datashare.
	AllowWrites bool `json:"allowWrites,omitempty" yaml:"allowWrites,omitempty"`

	// Identifier of the data consumer that is authorized to access the datashare. This identifier is an AWS account ID or a keyword, such as `ADX`.
	ConsumerIdentifier string `json:"consumerIdentifier,omitempty" yaml:"consumerIdentifier,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the datashare that producers are to authorize sharing for.

	   The following arguments are optional:
	*/
	DataShareArn string `json:"dataShareArn,omitempty" yaml:"dataShareArn,omitempty"`
}
