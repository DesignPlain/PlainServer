package types

type Medialive_ChannelEncoderSettingsNielsenConfiguration struct {
	// Enter the Distributor ID assigned to your organization by Nielsen.
	DistributorId string `json:"distributorId,omitempty" yaml:"distributorId,omitempty"`

	// Enables Nielsen PCM to ID3 tagging.
	NielsenPcmToId3Tagging string `json:"nielsenPcmToId3Tagging,omitempty" yaml:"nielsenPcmToId3Tagging,omitempty"`
}
