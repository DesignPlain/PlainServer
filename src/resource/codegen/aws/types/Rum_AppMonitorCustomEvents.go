package types

type Rum_AppMonitorCustomEvents struct {
	// Specifies whether this app monitor allows the web client to define and send custom events. The default is for custom events to be `DISABLED`. Valid values are `DISABLED` and `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
