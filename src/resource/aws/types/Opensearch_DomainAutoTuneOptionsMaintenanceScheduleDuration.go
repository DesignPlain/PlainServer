package types

type Opensearch_DomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Unit of time specifying the duration of an Auto-Tune maintenance window. Valid values: `HOURS`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// An integer specifying the value of the duration of an Auto-Tune maintenance window.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
