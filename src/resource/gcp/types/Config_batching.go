package types

type Config_batching struct {
	//
	SendAfter string `json:"sendAfter,omitempty" yaml:"sendAfter,omitempty"`

	//
	EnableBatching bool `json:"enableBatching,omitempty" yaml:"enableBatching,omitempty"`
}
