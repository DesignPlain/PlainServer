package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettings struct {
	//
	ProgramDateTime string `json:"programDateTime,omitempty" yaml:"programDateTime,omitempty"`

	//
	SegmentLength int `json:"segmentLength,omitempty" yaml:"segmentLength,omitempty"`

	//
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	//
	IframeOnlyPlaylists string `json:"iframeOnlyPlaylists,omitempty" yaml:"iframeOnlyPlaylists,omitempty"`

	//
	KeyFormatVersions string `json:"keyFormatVersions,omitempty" yaml:"keyFormatVersions,omitempty"`

	//
	KeyProviderSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsKeyProviderSettings `json:"keyProviderSettings,omitempty" yaml:"keyProviderSettings,omitempty"`

	//
	ManifestDurationFormat string `json:"manifestDurationFormat,omitempty" yaml:"manifestDurationFormat,omitempty"`

	//
	OutputSelection string `json:"outputSelection,omitempty" yaml:"outputSelection,omitempty"`

	//
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	HlsId3SegmentTagging string `json:"hlsId3SegmentTagging,omitempty" yaml:"hlsId3SegmentTagging,omitempty"`

	//
	ProgramDateTimePeriod int `json:"programDateTimePeriod,omitempty" yaml:"programDateTimePeriod,omitempty"`

	//
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	// The ad marker type for this output group.
	AdMarkers []string `json:"adMarkers,omitempty" yaml:"adMarkers,omitempty"`

	//
	BaseUrlContent1 string `json:"baseUrlContent1,omitempty" yaml:"baseUrlContent1,omitempty"`

	//
	ConstantIv string `json:"constantIv,omitempty" yaml:"constantIv,omitempty"`

	//
	HlsCdnSettings []Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSetting `json:"hlsCdnSettings,omitempty" yaml:"hlsCdnSettings,omitempty"`

	//
	IncompleteSegmentBehavior string `json:"incompleteSegmentBehavior,omitempty" yaml:"incompleteSegmentBehavior,omitempty"`

	//
	ManifestCompression string `json:"manifestCompression,omitempty" yaml:"manifestCompression,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	BaseUrlManifest1 string `json:"baseUrlManifest1,omitempty" yaml:"baseUrlManifest1,omitempty"`

	//
	ClientCache string `json:"clientCache,omitempty" yaml:"clientCache,omitempty"`

	//
	DirectoryStructure string `json:"directoryStructure,omitempty" yaml:"directoryStructure,omitempty"`

	//
	IvInManifest string `json:"ivInManifest,omitempty" yaml:"ivInManifest,omitempty"`

	//
	KeepSegments int `json:"keepSegments,omitempty" yaml:"keepSegments,omitempty"`

	//
	KeyFormat string `json:"keyFormat,omitempty" yaml:"keyFormat,omitempty"`

	// Indicates ID3 frame that has the timecode.
	TimedMetadataId3Frame string `json:"timedMetadataId3Frame,omitempty" yaml:"timedMetadataId3Frame,omitempty"`

	//
	TimedMetadataId3Period int `json:"timedMetadataId3Period,omitempty" yaml:"timedMetadataId3Period,omitempty"`

	//
	TsFileMode string `json:"tsFileMode,omitempty" yaml:"tsFileMode,omitempty"`

	//
	CaptionLanguageMappings []Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsCaptionLanguageMapping `json:"captionLanguageMappings,omitempty" yaml:"captionLanguageMappings,omitempty"`

	//
	DiscontinuityTags string `json:"discontinuityTags,omitempty" yaml:"discontinuityTags,omitempty"`

	//
	StreamInfResolution string `json:"streamInfResolution,omitempty" yaml:"streamInfResolution,omitempty"`

	//
	TimestampDeltaMilliseconds int `json:"timestampDeltaMilliseconds,omitempty" yaml:"timestampDeltaMilliseconds,omitempty"`

	//
	BaseUrlContent string `json:"baseUrlContent,omitempty" yaml:"baseUrlContent,omitempty"`

	//
	BaseUrlManifest string `json:"baseUrlManifest,omitempty" yaml:"baseUrlManifest,omitempty"`

	//
	CodecSpecification string `json:"codecSpecification,omitempty" yaml:"codecSpecification,omitempty"`

	//
	IndexNSegments int `json:"indexNSegments,omitempty" yaml:"indexNSegments,omitempty"`

	//
	ProgramDateTimeClock string `json:"programDateTimeClock,omitempty" yaml:"programDateTimeClock,omitempty"`

	//
	CaptionLanguageSetting string `json:"captionLanguageSetting,omitempty" yaml:"captionLanguageSetting,omitempty"`

	//
	IvSource string `json:"ivSource,omitempty" yaml:"ivSource,omitempty"`

	//
	MinSegmentLength int `json:"minSegmentLength,omitempty" yaml:"minSegmentLength,omitempty"`

	//
	RedundantManifest string `json:"redundantManifest,omitempty" yaml:"redundantManifest,omitempty"`

	//
	SegmentsPerSubdirectory int `json:"segmentsPerSubdirectory,omitempty" yaml:"segmentsPerSubdirectory,omitempty"`
}
