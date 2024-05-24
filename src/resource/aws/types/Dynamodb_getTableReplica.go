package types

type Dynamodb_getTableReplica struct {
	//
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	//
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`
}
