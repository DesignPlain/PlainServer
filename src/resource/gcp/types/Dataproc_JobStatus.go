package types

type Dataproc_JobStatus struct {
	// Output-only. Optional job state details, such as an error description if the state is ERROR
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	// Output-only. A state message specifying the overall job state
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Output-only. The time when this state was entered
	StateStartTime string `json:"stateStartTime,omitempty" yaml:"stateStartTime,omitempty"`

	// Output-only. Additional state information, which includes status reported by the agent
	Substate string `json:"substate,omitempty" yaml:"substate,omitempty"`
}
