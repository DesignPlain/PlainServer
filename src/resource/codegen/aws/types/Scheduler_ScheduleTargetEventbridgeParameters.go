package types

type Scheduler_ScheduleTargetEventbridgeParameters struct {
	// Free-form string used to decide what fields to expect in the event detail. Up to 128 characters.
	DetailType string `json:"detailType,omitempty" yaml:"detailType,omitempty"`

	// Source of the event.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
