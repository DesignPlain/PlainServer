package types

type Iot_TopicRuleErrorActionElasticsearch struct {
	// The type of document you are storing.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The endpoint of your Elasticsearch domain.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// The unique identifier for the document you are storing.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The Elasticsearch index where you want to store your data.
	Index string `json:"index,omitempty" yaml:"index,omitempty"`

	// The IAM role ARN that has access to Elasticsearch.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
