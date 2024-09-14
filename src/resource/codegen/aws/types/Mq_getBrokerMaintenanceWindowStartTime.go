package types

type Mq_getBrokerMaintenanceWindowStartTime struct {
	//
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	//
	TimeOfDay string `json:"timeOfDay,omitempty" yaml:"timeOfDay,omitempty"`

	//
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}
