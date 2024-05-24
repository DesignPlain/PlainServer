package types

type Mq_getBrokerMaintenanceWindowStartTime struct {
	//
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	//
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	//
	TimeOfDay string `json:"timeOfDay,omitempty" yaml:"timeOfDay,omitempty"`
}
