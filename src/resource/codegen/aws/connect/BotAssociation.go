package connect

import types "libds/aws/types"

type BotAssociation struct {
	// The identifier of the Amazon Connect instance. You can find the instanceId in the ARN of the instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Configuration information of an Amazon Lex (V1) bot. Detailed below.
	LexBot types.Connect_BotAssociationLexBot `json:"lexBot,omitempty" yaml:"lexBot,omitempty"`
}
