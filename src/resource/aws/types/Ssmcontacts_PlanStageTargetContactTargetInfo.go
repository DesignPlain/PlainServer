package types

type Ssmcontacts_PlanStageTargetContactTargetInfo struct {
	// The Amazon Resource Name (ARN) of the contact.
	ContactId string `json:"contactId,omitempty" yaml:"contactId,omitempty"`

	// A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.
	IsEssential bool `json:"isEssential,omitempty" yaml:"isEssential,omitempty"`
}
