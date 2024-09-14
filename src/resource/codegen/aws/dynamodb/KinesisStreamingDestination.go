package dynamodb

type KinesisStreamingDestination struct {
	// The ARN for a Kinesis data stream. This must exist in the same account and region as the DynamoDB table.
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`

	/*
	   The name of the DynamoDB table. There
	   can only be one Kinesis streaming destination for a given DynamoDB table.
	*/
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
