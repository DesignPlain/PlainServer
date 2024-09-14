package sfn

import types "libds/aws/types"

type StateMachine struct {
	// The name of the state machine. The name should only contain `0`-`9`, `A`-`Z`, `a`-`z`, `-` and `_`. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Determines whether a Standard or Express state machine is created. The default is `STANDARD`. You cannot update the type of a state machine once it has been created. Valid values: `STANDARD`, `EXPRESS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration types.Sfn_StateMachineTracingConfiguration `json:"tracingConfiguration,omitempty" yaml:"tracingConfiguration,omitempty"`

	// The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) definition of the state machine.
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`

	// Defines what encryption configuration is used to encrypt data in the State Machine. For more information see [TBD] in the AWS Step Functions User Guide.
	EncryptionConfiguration types.Sfn_StateMachineEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration types.Sfn_StateMachineLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`

	// Set to true to publish a version of the state machine during creation. Default: false.
	Publish bool `json:"publish,omitempty" yaml:"publish,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
