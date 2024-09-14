package types

type Bedrock_AgentAgentActionGroupActionGroupExecutor struct {
	/*
	   ARN of the Lambda function containing the business logic that is carried out upon invoking the action.
	   Only one of `lambda` or `custom_control` can be specified.
	*/
	Lambda string `json:"lambda,omitempty" yaml:"lambda,omitempty"`

	/*
	   Custom control method for handling the information elicited from the user. Valid values: `RETURN_CONTROL`.
	   To skip using a Lambda function and instead return the predicted action group, in addition to the parameters and information required for it, in the `InvokeAgent` response, specify `RETURN_CONTROL`.
	   Only one of `custom_control` or `lambda` can be specified.
	*/
	CustomControl string `json:"customControl,omitempty" yaml:"customControl,omitempty"`
}
