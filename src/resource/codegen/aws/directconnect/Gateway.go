package directconnect

type Gateway struct {
	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn string `json:"amazonSideAsn,omitempty" yaml:"amazonSideAsn,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
