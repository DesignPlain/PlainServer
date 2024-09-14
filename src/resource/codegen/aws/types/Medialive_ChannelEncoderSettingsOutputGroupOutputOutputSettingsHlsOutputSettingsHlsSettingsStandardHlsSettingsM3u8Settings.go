package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettingsStandardHlsSettingsM3u8Settings struct {
	//
	PmtPid string `json:"pmtPid,omitempty" yaml:"pmtPid,omitempty"`

	//
	Scte35Behavior string `json:"scte35Behavior,omitempty" yaml:"scte35Behavior,omitempty"`

	//
	NielsenId3Behavior string `json:"nielsenId3Behavior,omitempty" yaml:"nielsenId3Behavior,omitempty"`

	//
	PcrPeriod int `json:"pcrPeriod,omitempty" yaml:"pcrPeriod,omitempty"`

	//
	PatInterval int `json:"patInterval,omitempty" yaml:"patInterval,omitempty"`

	//
	PmtInterval int `json:"pmtInterval,omitempty" yaml:"pmtInterval,omitempty"`

	//
	ProgramNum int `json:"programNum,omitempty" yaml:"programNum,omitempty"`

	// PID from which to read SCTE-35 messages.
	Scte35Pid string `json:"scte35Pid,omitempty" yaml:"scte35Pid,omitempty"`

	//
	TimedMetadataBehavior string `json:"timedMetadataBehavior,omitempty" yaml:"timedMetadataBehavior,omitempty"`

	//
	TimedMetadataPid string `json:"timedMetadataPid,omitempty" yaml:"timedMetadataPid,omitempty"`

	//
	AudioFramesPerPes int `json:"audioFramesPerPes,omitempty" yaml:"audioFramesPerPes,omitempty"`

	//
	AudioPids string `json:"audioPids,omitempty" yaml:"audioPids,omitempty"`

	//
	TransportStreamId int `json:"transportStreamId,omitempty" yaml:"transportStreamId,omitempty"`

	//
	EcmPid string `json:"ecmPid,omitempty" yaml:"ecmPid,omitempty"`

	//
	PcrPid string `json:"pcrPid,omitempty" yaml:"pcrPid,omitempty"`

	//
	PcrControl string `json:"pcrControl,omitempty" yaml:"pcrControl,omitempty"`

	//
	VideoPid string `json:"videoPid,omitempty" yaml:"videoPid,omitempty"`
}
