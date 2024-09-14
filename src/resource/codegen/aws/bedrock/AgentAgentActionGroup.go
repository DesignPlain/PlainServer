package bedrock

import types "libds/aws/types"

type AgentAgentActionGroup struct {
	// Name of the action group.
	ActionGroupName string `json:"actionGroupName,omitempty" yaml:"actionGroupName,omitempty"`

	// Whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request. Valid values: `ENABLED`, `DISABLED`.
	ActionGroupState string `json:"actionGroupState,omitempty" yaml:"actionGroupState,omitempty"`

	// The unique identifier of the agent for which to create the action group.
	AgentId string `json:"agentId,omitempty" yaml:"agentId,omitempty"`

	// Version of the agent for which to create the action group. Valid values: `DRAFT`.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`

	/*
	   Describes the function schema for the action group.
	   Each function represents an action in an action group.
	   See `function_schema` Block for details.
	*/
	FunctionSchema types.Bedrock_AgentAgentActionGroupFunctionSchema `json:"functionSchema,omitempty" yaml:"functionSchema,omitempty"`

	// Whether the in-use check is skipped when deleting the action group.
	SkipResourceInUseCheck bool `json:"skipResourceInUseCheck,omitempty" yaml:"skipResourceInUseCheck,omitempty"`

	/*
	   ARN of the Lambda function containing the business logic that is carried out upon invoking the action or custom control method for handling the information elicited from the user. See `action_group_executor` Block for details.

	   The following arguments are optional:
	*/
	ActionGroupExecutor types.Bedrock_AgentAgentActionGroupActionGroupExecutor `json:"actionGroupExecutor,omitempty" yaml:"actionGroupExecutor,omitempty"`

	// Either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema. For more information, see [Action group OpenAPI schemas](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-api-schema.html). See `api_schema` Block for details.
	ApiSchema types.Bedrock_AgentAgentActionGroupApiSchema `json:"apiSchema,omitempty" yaml:"apiSchema,omitempty"`

	// Description of the action group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// To allow your agent to request the user for additional information when trying to complete a task, set this argument to `AMAZON.UserInput`. You must leave the `description`, `api_schema`, and `action_group_executor` arguments blank for this action group. Valid values: `AMAZON.UserInput`.
	ParentActionGroupSignature string `json:"parentActionGroupSignature,omitempty" yaml:"parentActionGroupSignature,omitempty"`
}
