package types

type Scheduler_ScheduleTargetEventbridgeParameters struct {
	// Source of the event.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// Free-form string used to decide what fields to expect in the event detail. Up to 128 characters.
	DetailType string `json:"detailType,omitempty" yaml:"detailType,omitempty"`
}
