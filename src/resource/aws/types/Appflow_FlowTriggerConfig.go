package types

type Appflow_FlowTriggerConfig struct {
	// Configuration details of a schedule-triggered flow as defined by the user. Currently, these settings only apply to the `Scheduled` trigger type. See Scheduled Trigger Properties for details.
	TriggerProperties Appflow_FlowTriggerConfigTriggerProperties `json:"triggerProperties,omitempty" yaml:"triggerProperties,omitempty"`

	// Type of flow trigger. Valid values are `Scheduled`, `Event`, and `OnDemand`.
	TriggerType string `json:"triggerType,omitempty" yaml:"triggerType,omitempty"`
}
