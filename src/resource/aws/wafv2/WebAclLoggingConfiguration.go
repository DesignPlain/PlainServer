package wafv2

import types "DesignSphere_Server/src/resource/aws/types"

type WebAclLoggingConfiguration struct {
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter types.Wafv2_WebAclLoggingConfigurationLoggingFilter `json:"loggingFilter,omitempty" yaml:"loggingFilter,omitempty"`

	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redacted_fields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields []types.Wafv2_WebAclLoggingConfigurationRedactedField `json:"redactedFields,omitempty" yaml:"redactedFields,omitempty"`

	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `log_destination_configs`.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. --Note:-- data firehose, log group, or bucket name --must-- be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs []string `json:"logDestinationConfigs,omitempty" yaml:"logDestinationConfigs,omitempty"`
}
