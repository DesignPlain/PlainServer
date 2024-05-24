package ram

type PrincipalAssociation struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn string `json:"resourceShareArn,omitempty" yaml:"resourceShareArn,omitempty"`
}
