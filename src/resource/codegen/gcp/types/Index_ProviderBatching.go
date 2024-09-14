package types

type Index_ProviderBatching struct {
	//
	EnableBatching bool `json:"enableBatching,omitempty" yaml:"enableBatching,omitempty"`

	//
	SendAfter string `json:"sendAfter,omitempty" yaml:"sendAfter,omitempty"`
}
