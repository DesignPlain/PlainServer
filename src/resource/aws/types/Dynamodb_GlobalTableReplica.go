package types

type Dynamodb_GlobalTableReplica struct {
	// AWS region name of replica DynamoDB TableE.g., `us-east-1`
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`
}
