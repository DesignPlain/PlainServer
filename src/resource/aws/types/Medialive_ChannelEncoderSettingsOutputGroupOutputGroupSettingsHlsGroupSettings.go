package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettings struct {
	//
	BaseUrlContent string `json:"baseUrlContent,omitempty" yaml:"baseUrlContent,omitempty"`

	//
	CaptionLanguageMappings []Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsCaptionLanguageMapping `json:"captionLanguageMappings,omitempty" yaml:"captionLanguageMappings,omitempty"`

	//
	IvSource string `json:"ivSource,omitempty" yaml:"ivSource,omitempty"`

	//
	ManifestCompression string `json:"manifestCompression,omitempty" yaml:"manifestCompression,omitempty"`

	//
	ProgramDateTimeClock string `json:"programDateTimeClock,omitempty" yaml:"programDateTimeClock,omitempty"`

	//
	BaseUrlContent1 string `json:"baseUrlContent1,omitempty" yaml:"baseUrlContent1,omitempty"`

	//
	DirectoryStructure string `json:"directoryStructure,omitempty" yaml:"directoryStructure,omitempty"`

	//
	IvInManifest string `json:"ivInManifest,omitempty" yaml:"ivInManifest,omitempty"`

	//
	KeyFormat string `json:"keyFormat,omitempty" yaml:"keyFormat,omitempty"`

	//
	SegmentsPerSubdirectory int `json:"segmentsPerSubdirectory,omitempty" yaml:"segmentsPerSubdirectory,omitempty"`

	//
	CodecSpecification string `json:"codecSpecification,omitempty" yaml:"codecSpecification,omitempty"`

	// A director and base filename where archive files should be written. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	HlsCdnSettings []Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSetting `json:"hlsCdnSettings,omitempty" yaml:"hlsCdnSettings,omitempty"`

	//
	MinSegmentLength int `json:"minSegmentLength,omitempty" yaml:"minSegmentLength,omitempty"`

	//
	ProgramDateTime string `json:"programDateTime,omitempty" yaml:"programDateTime,omitempty"`

	// The ad marker type for this output group.
	AdMarkers []string `json:"adMarkers,omitempty" yaml:"adMarkers,omitempty"`

	//
	BaseUrlManifest1 string `json:"baseUrlManifest1,omitempty" yaml:"baseUrlManifest1,omitempty"`

	//
	HlsId3SegmentTagging string `json:"hlsId3SegmentTagging,omitempty" yaml:"hlsId3SegmentTagging,omitempty"`

	//
	IncompleteSegmentBehavior string `json:"incompleteSegmentBehavior,omitempty" yaml:"incompleteSegmentBehavior,omitempty"`

	//
	KeepSegments int `json:"keepSegments,omitempty" yaml:"keepSegments,omitempty"`

	//
	ManifestDurationFormat string `json:"manifestDurationFormat,omitempty" yaml:"manifestDurationFormat,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	ProgramDateTimePeriod int `json:"programDateTimePeriod,omitempty" yaml:"programDateTimePeriod,omitempty"`

	//
	CaptionLanguageSetting string `json:"captionLanguageSetting,omitempty" yaml:"captionLanguageSetting,omitempty"`

	//
	ClientCache string `json:"clientCache,omitempty" yaml:"clientCache,omitempty"`

	//
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// Controls the behavior of the RTMP group if input becomes unavailable.
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	//
	RedundantManifest string `json:"redundantManifest,omitempty" yaml:"redundantManifest,omitempty"`

	//
	TimestampDeltaMilliseconds int `json:"timestampDeltaMilliseconds,omitempty" yaml:"timestampDeltaMilliseconds,omitempty"`

	//
	TsFileMode string `json:"tsFileMode,omitempty" yaml:"tsFileMode,omitempty"`

	//
	ConstantIv string `json:"constantIv,omitempty" yaml:"constantIv,omitempty"`

	//
	DiscontinuityTags string `json:"discontinuityTags,omitempty" yaml:"discontinuityTags,omitempty"`

	//
	IndexNSegments int `json:"indexNSegments,omitempty" yaml:"indexNSegments,omitempty"`

	//
	KeyProviderSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsKeyProviderSettings `json:"keyProviderSettings,omitempty" yaml:"keyProviderSettings,omitempty"`

	// Indicates ID3 frame that has the timecode.
	TimedMetadataId3Frame string `json:"timedMetadataId3Frame,omitempty" yaml:"timedMetadataId3Frame,omitempty"`

	//
	BaseUrlManifest string `json:"baseUrlManifest,omitempty" yaml:"baseUrlManifest,omitempty"`

	//
	IframeOnlyPlaylists string `json:"iframeOnlyPlaylists,omitempty" yaml:"iframeOnlyPlaylists,omitempty"`

	//
	KeyFormatVersions string `json:"keyFormatVersions,omitempty" yaml:"keyFormatVersions,omitempty"`

	//
	StreamInfResolution string `json:"streamInfResolution,omitempty" yaml:"streamInfResolution,omitempty"`

	//
	OutputSelection string `json:"outputSelection,omitempty" yaml:"outputSelection,omitempty"`

	//
	SegmentLength int `json:"segmentLength,omitempty" yaml:"segmentLength,omitempty"`

	//
	TimedMetadataId3Period int `json:"timedMetadataId3Period,omitempty" yaml:"timedMetadataId3Period,omitempty"`
}
