package types

type Index_ProviderBatching struct {
	//
	SendAfter string `json:"sendAfter,omitempty" yaml:"sendAfter,omitempty"`

	//
	EnableBatching bool `json:"enableBatching,omitempty" yaml:"enableBatching,omitempty"`
}
