package types

type Config_batching struct {
	//
	EnableBatching bool `json:"enableBatching,omitempty" yaml:"enableBatching,omitempty"`

	//
	SendAfter string `json:"sendAfter,omitempty" yaml:"sendAfter,omitempty"`
}
