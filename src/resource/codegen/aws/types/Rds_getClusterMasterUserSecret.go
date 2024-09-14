package types

type Rds_getClusterMasterUserSecret struct {
	//
	SecretStatus string `json:"secretStatus,omitempty" yaml:"secretStatus,omitempty"`

	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	//
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`
}
