package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettings struct {
	//
	SegmentationStyle string `json:"segmentationStyle,omitempty" yaml:"segmentationStyle,omitempty"`

	//
	DvbTdtSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettings `json:"dvbTdtSettings,omitempty" yaml:"dvbTdtSettings,omitempty"`

	//
	EbpPlacement string `json:"ebpPlacement,omitempty" yaml:"ebpPlacement,omitempty"`

	//
	Klv string `json:"klv,omitempty" yaml:"klv,omitempty"`

	//
	TimedMetadataBehavior string `json:"timedMetadataBehavior,omitempty" yaml:"timedMetadataBehavior,omitempty"`

	//
	VideoPid string `json:"videoPid,omitempty" yaml:"videoPid,omitempty"`

	//
	AudioStreamType string `json:"audioStreamType,omitempty" yaml:"audioStreamType,omitempty"`

	//
	EcmPid string `json:"ecmPid,omitempty" yaml:"ecmPid,omitempty"`

	//
	KlvDataPids string `json:"klvDataPids,omitempty" yaml:"klvDataPids,omitempty"`

	//
	DvbSdtSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettings `json:"dvbSdtSettings,omitempty" yaml:"dvbSdtSettings,omitempty"`

	//
	Ebif string `json:"ebif,omitempty" yaml:"ebif,omitempty"`

	//
	FragmentTime float64 `json:"fragmentTime,omitempty" yaml:"fragmentTime,omitempty"`

	//
	PmtPid string `json:"pmtPid,omitempty" yaml:"pmtPid,omitempty"`

	//
	ProgramNum int `json:"programNum,omitempty" yaml:"programNum,omitempty"`

	// Average bitrate in bits/second.
	Bitrate int `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	//
	BufferModel string `json:"bufferModel,omitempty" yaml:"bufferModel,omitempty"`

	//
	DvbNitSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbNitSettings `json:"dvbNitSettings,omitempty" yaml:"dvbNitSettings,omitempty"`

	//
	NullPacketBitrate float64 `json:"nullPacketBitrate,omitempty" yaml:"nullPacketBitrate,omitempty"`

	//
	PmtInterval int `json:"pmtInterval,omitempty" yaml:"pmtInterval,omitempty"`

	//
	RateMode string `json:"rateMode,omitempty" yaml:"rateMode,omitempty"`

	//
	TimedMetadataPid string `json:"timedMetadataPid,omitempty" yaml:"timedMetadataPid,omitempty"`

	//
	AbsentInputAudioBehavior string `json:"absentInputAudioBehavior,omitempty" yaml:"absentInputAudioBehavior,omitempty"`

	//
	AribCaptionsPidControl string `json:"aribCaptionsPidControl,omitempty" yaml:"aribCaptionsPidControl,omitempty"`

	//
	EsRateInPes string `json:"esRateInPes,omitempty" yaml:"esRateInPes,omitempty"`

	//
	TransportStreamId int `json:"transportStreamId,omitempty" yaml:"transportStreamId,omitempty"`

	//
	Arib string `json:"arib,omitempty" yaml:"arib,omitempty"`

	//
	AribCaptionsPid string `json:"aribCaptionsPid,omitempty" yaml:"aribCaptionsPid,omitempty"`

	//
	PcrControl string `json:"pcrControl,omitempty" yaml:"pcrControl,omitempty"`

	//
	EtvSignalPid string `json:"etvSignalPid,omitempty" yaml:"etvSignalPid,omitempty"`

	//
	PcrPid string `json:"pcrPid,omitempty" yaml:"pcrPid,omitempty"`

	//
	Scte27Pids string `json:"scte27Pids,omitempty" yaml:"scte27Pids,omitempty"`

	//
	SegmentationMarkers string `json:"segmentationMarkers,omitempty" yaml:"segmentationMarkers,omitempty"`

	//
	DvbSubPids string `json:"dvbSubPids,omitempty" yaml:"dvbSubPids,omitempty"`

	//
	EbpLookaheadMs int `json:"ebpLookaheadMs,omitempty" yaml:"ebpLookaheadMs,omitempty"`

	//
	EtvPlatformPid string `json:"etvPlatformPid,omitempty" yaml:"etvPlatformPid,omitempty"`

	//
	DvbTeletextPid string `json:"dvbTeletextPid,omitempty" yaml:"dvbTeletextPid,omitempty"`

	//
	Scte35Control string `json:"scte35Control,omitempty" yaml:"scte35Control,omitempty"`

	//
	SegmentationTime float64 `json:"segmentationTime,omitempty" yaml:"segmentationTime,omitempty"`

	//
	AudioBufferModel string `json:"audioBufferModel,omitempty" yaml:"audioBufferModel,omitempty"`

	//
	AudioFramesPerPes int `json:"audioFramesPerPes,omitempty" yaml:"audioFramesPerPes,omitempty"`

	//
	CcDescriptor string `json:"ccDescriptor,omitempty" yaml:"ccDescriptor,omitempty"`

	//
	PatInterval int `json:"patInterval,omitempty" yaml:"patInterval,omitempty"`

	//
	PcrPeriod int `json:"pcrPeriod,omitempty" yaml:"pcrPeriod,omitempty"`

	// PID from which to read SCTE-35 messages.
	Scte35Pid string `json:"scte35Pid,omitempty" yaml:"scte35Pid,omitempty"`

	//
	AudioPids string `json:"audioPids,omitempty" yaml:"audioPids,omitempty"`

	//
	EbpAudioInterval string `json:"ebpAudioInterval,omitempty" yaml:"ebpAudioInterval,omitempty"`

	//
	NielsenId3Behavior string `json:"nielsenId3Behavior,omitempty" yaml:"nielsenId3Behavior,omitempty"`
}
