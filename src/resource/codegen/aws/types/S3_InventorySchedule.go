package types

type S3_InventorySchedule struct {
	// Specifies how frequently inventory results are produced. Valid values: `Daily`, `Weekly`.
	Frequency string `json:"frequency,omitempty" yaml:"frequency,omitempty"`
}
