package types

type Monitoring_AlertPolicyCreationRecord struct {
	/*
	   (Output)
	   The email address of the user making the change.
	*/
	MutatedBy string `json:"mutatedBy,omitempty" yaml:"mutatedBy,omitempty"`

	/*
	   (Output)
	   When the change occurred.
	*/
	MutateTime string `json:"mutateTime,omitempty" yaml:"mutateTime,omitempty"`
}
