package ssmcontacts

import types "libds/aws/types"

type Plan struct {
	// The Amazon Resource Name (ARN) of the contact or escalation plan.
	ContactId string `json:"contactId,omitempty" yaml:"contactId,omitempty"`

	// One or more configuration blocks for specifying a list of stages that the escalation plan or engagement plan uses to engage contacts and contact methods. See Stage below for more details.
	Stages []types.Ssmcontacts_PlanStage `json:"stages,omitempty" yaml:"stages,omitempty"`
}
