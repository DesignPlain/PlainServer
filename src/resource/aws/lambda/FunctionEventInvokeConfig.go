package lambda

import types "DesignSphere_Server/src/resource/aws/types"

type FunctionEventInvokeConfig struct {
	// Configuration block with destination configuration. See below for details.
	DestinationConfig types.Lambda_FunctionEventInvokeConfigDestinationConfig `json:"destinationConfig,omitempty" yaml:"destinationConfig,omitempty"`

	/*
	   Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.

	   The following arguments are optional:
	*/
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds int `json:"maximumEventAgeInSeconds,omitempty" yaml:"maximumEventAgeInSeconds,omitempty"`

	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts int `json:"maximumRetryAttempts,omitempty" yaml:"maximumRetryAttempts,omitempty"`

	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`
}
