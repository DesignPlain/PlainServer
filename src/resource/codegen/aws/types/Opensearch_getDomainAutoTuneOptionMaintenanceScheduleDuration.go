package types

type Opensearch_getDomainAutoTuneOptionMaintenanceScheduleDuration struct {
	// Unit of time.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Duration of an Auto-Tune maintenance window.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
