package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265Settings struct {
	// Size of buffer in bits.
	BufSize int `json:"bufSize,omitempty" yaml:"bufSize,omitempty"`

	// Sets the scan type of the output.
	ScanType string `json:"scanType,omitempty" yaml:"scanType,omitempty"`

	// Set the H265 tier in the output.
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	// Controls the target quality for the video encode.
	QvbrQualityLevel int `json:"qvbrQualityLevel,omitempty" yaml:"qvbrQualityLevel,omitempty"`

	// Number of slices per picture.
	Slices int `json:"slices,omitempty" yaml:"slices,omitempty"`

	// H265 level.
	Level string `json:"level,omitempty" yaml:"level,omitempty"`

	// Set the maximum bitrate in order to accommodate expected spikes in the complexity of the video.
	MaxBitrate int `json:"maxBitrate,omitempty" yaml:"maxBitrate,omitempty"`

	// Pixel Aspect Ratio denominator.
	ParDenominator int `json:"parDenominator,omitempty" yaml:"parDenominator,omitempty"`

	// Filters to apply to an encode. See H265 Filter Settings for more details.
	FilterSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsFilterSettings `json:"filterSettings,omitempty" yaml:"filterSettings,omitempty"`

	// Determines how timecodes should be inserted into the video elementary stream.
	TimecodeInsertion string `json:"timecodeInsertion,omitempty" yaml:"timecodeInsertion,omitempty"`

	// Whether or not EML should insert an Alternative Transfer Function SEI message.
	AlternativeTransferFunction string `json:"alternativeTransferFunction,omitempty" yaml:"alternativeTransferFunction,omitempty"`

	// Pixel Aspect Ratio numerator.
	ParNumerator int `json:"parNumerator,omitempty" yaml:"parNumerator,omitempty"`

	// Apply a burned in timecode. See H265 Timecode Burnin Settings for more details.
	TimecodeBurninSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsTimecodeBurninSettings `json:"timecodeBurninSettings,omitempty" yaml:"timecodeBurninSettings,omitempty"`

	// Frequency of closed GOPs.
	GopClosedCadence int `json:"gopClosedCadence,omitempty" yaml:"gopClosedCadence,omitempty"`

	// Indicates that AFD values will be written into the output stream.
	AfdSignaling string `json:"afdSignaling,omitempty" yaml:"afdSignaling,omitempty"`

	// Average bitrate in bits/second.
	Bitrate int `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	// Framerate numerator.
	FramerateNumerator int `json:"framerateNumerator,omitempty" yaml:"framerateNumerator,omitempty"`

	// GOP size in units of either frames of seconds per `gop_size_units`.
	GopSize float64 `json:"gopSize,omitempty" yaml:"gopSize,omitempty"`

	// AAC profile.
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`

	// Includes color space metadata in the output.
	ColorMetadata string `json:"colorMetadata,omitempty" yaml:"colorMetadata,omitempty"`

	// Four bit AFD value to write on all frames of video in the output stream.
	FixedAfd string `json:"fixedAfd,omitempty" yaml:"fixedAfd,omitempty"`

	//
	FlickerAq string `json:"flickerAq,omitempty" yaml:"flickerAq,omitempty"`

	//
	MinIInterval int `json:"minIInterval,omitempty" yaml:"minIInterval,omitempty"`

	// The rate control mode.
	RateControlMode string `json:"rateControlMode,omitempty" yaml:"rateControlMode,omitempty"`

	// Define the color metadata for the output. H265 Color Space Settings for more details.
	ColorSpaceSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettings `json:"colorSpaceSettings,omitempty" yaml:"colorSpaceSettings,omitempty"`

	// Indicates if the `gop_size` is specified in frames or seconds.
	GopSizeUnits string `json:"gopSizeUnits,omitempty" yaml:"gopSizeUnits,omitempty"`

	// Amount of lookahead.
	LookAheadRateControl string `json:"lookAheadRateControl,omitempty" yaml:"lookAheadRateControl,omitempty"`

	// Enables or disables adaptive quantization.
	AdaptiveQuantization string `json:"adaptiveQuantization,omitempty" yaml:"adaptiveQuantization,omitempty"`

	// Framerate denominator.
	FramerateDenominator int `json:"framerateDenominator,omitempty" yaml:"framerateDenominator,omitempty"`

	// Scene change detection.
	SceneChangeDetect string `json:"sceneChangeDetect,omitempty" yaml:"sceneChangeDetect,omitempty"`
}
