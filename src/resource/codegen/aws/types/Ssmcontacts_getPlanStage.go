package types

type Ssmcontacts_getPlanStage struct {
	//
	DurationInMinutes int `json:"durationInMinutes,omitempty" yaml:"durationInMinutes,omitempty"`

	//
	Targets []Ssmcontacts_getPlanStageTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}
