package types

type Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenCbetSettings struct {
	//
	CbetCheckDigitString string `json:"cbetCheckDigitString,omitempty" yaml:"cbetCheckDigitString,omitempty"`

	// Determines the method of CBET insertion mode when prior encoding is detected on the same layer.
	CbetStepaside string `json:"cbetStepaside,omitempty" yaml:"cbetStepaside,omitempty"`

	// CBET source ID to use in the watermark.
	Csid string `json:"csid,omitempty" yaml:"csid,omitempty"`
}
