package types

type Autoscaling_GroupInstanceRefreshPreferencesAlarmSpecification struct {
	// List of Cloudwatch alarms. If any of these alarms goes into ALARM state, Instance Refresh is failed.
	Alarms []string `json:"alarms,omitempty" yaml:"alarms,omitempty"`
}
