package types

type Glue_TriggerActionNotificationProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter int `json:"notifyDelayAfter,omitempty" yaml:"notifyDelayAfter,omitempty"`
}
