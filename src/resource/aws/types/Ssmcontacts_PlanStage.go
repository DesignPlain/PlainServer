package types

type Ssmcontacts_PlanStage struct {
	// The time to wait until beginning the next stage. The duration can only be set to 0 if a target is specified.
	DurationInMinutes int `json:"durationInMinutes,omitempty" yaml:"durationInMinutes,omitempty"`

	// One or more configuration blocks for specifying the contacts or contact methods that the escalation plan or engagement plan is engaging. See Target below for more details.
	Targets []Ssmcontacts_PlanStageTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}
