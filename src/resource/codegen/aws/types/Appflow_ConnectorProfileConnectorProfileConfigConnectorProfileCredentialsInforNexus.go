package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus struct {
	// The Access Key portion of the credentials.
	AccessKeyId string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`

	// Encryption keys used to encrypt data.
	Datakey string `json:"datakey,omitempty" yaml:"datakey,omitempty"`

	// The secret key used to sign requests.
	SecretAccessKey string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`

	// Identifier for the user.
	UserId string `json:"userId,omitempty" yaml:"userId,omitempty"`
}
