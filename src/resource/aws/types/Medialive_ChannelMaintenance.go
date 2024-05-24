package types

type Medialive_ChannelMaintenance struct {
	// The day of the week to use for maintenance.
	MaintenanceDay string `json:"maintenanceDay,omitempty" yaml:"maintenanceDay,omitempty"`

	// The hour maintenance will start.
	MaintenanceStartTime string `json:"maintenanceStartTime,omitempty" yaml:"maintenanceStartTime,omitempty"`
}
