package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type EventPermission struct {
	/*
	   The name of the event bus to set the permissions on.
	   If you omit this, the permissions are set on the `default` event bus.
	*/
	EventBusName string `json:"eventBusName,omitempty" yaml:"eventBusName,omitempty"`

	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `-` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// An identifier string for the external account that you are granting permissions to.
	StatementId string `json:"statementId,omitempty" yaml:"statementId,omitempty"`

	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition types.Cloudwatch_EventPermissionCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
