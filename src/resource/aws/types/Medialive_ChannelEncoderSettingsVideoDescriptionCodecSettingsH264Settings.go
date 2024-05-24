package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH264Settings struct {
	//
	BufFillPct int `json:"bufFillPct,omitempty" yaml:"bufFillPct,omitempty"`

	// GOP-B reference.
	GopBReference string `json:"gopBReference,omitempty" yaml:"gopBReference,omitempty"`

	// AAC profile.
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`

	// Subgop length.
	SubgopLength string `json:"subgopLength,omitempty" yaml:"subgopLength,omitempty"`

	// Size of buffer in bits.
	BufSize int `json:"bufSize,omitempty" yaml:"bufSize,omitempty"`

	// Controls the target quality for the video encode.
	QvbrQualityLevel int `json:"qvbrQualityLevel,omitempty" yaml:"qvbrQualityLevel,omitempty"`

	// Indicates how the output video frame rate is specified.
	FramerateControl string `json:"framerateControl,omitempty" yaml:"framerateControl,omitempty"`

	// Set the maximum bitrate in order to accommodate expected spikes in the complexity of the video.
	MaxBitrate int `json:"maxBitrate,omitempty" yaml:"maxBitrate,omitempty"`

	//
	MinIInterval int `json:"minIInterval,omitempty" yaml:"minIInterval,omitempty"`

	// Quality level.
	QualityLevel string `json:"qualityLevel,omitempty" yaml:"qualityLevel,omitempty"`

	// The rate control mode.
	RateControlMode string `json:"rateControlMode,omitempty" yaml:"rateControlMode,omitempty"`

	// Sets the scan type of the output.
	ScanType string `json:"scanType,omitempty" yaml:"scanType,omitempty"`

	// Softness.
	Softness int `json:"softness,omitempty" yaml:"softness,omitempty"`

	// Controls whether coding is performed on a field basis or on a frame basis.
	ForceFieldPictures string `json:"forceFieldPictures,omitempty" yaml:"forceFieldPictures,omitempty"`

	// Determines how timecodes should be inserted into the video elementary stream.
	TimecodeInsertion string `json:"timecodeInsertion,omitempty" yaml:"timecodeInsertion,omitempty"`

	// GOP size in units of either frames of seconds per `gop_size_units`.
	GopSize float64 `json:"gopSize,omitempty" yaml:"gopSize,omitempty"`

	// Indicates how the output pixel aspect ratio is specified.
	ParControl string `json:"parControl,omitempty" yaml:"parControl,omitempty"`

	// Makes adjustments within each frame based on spatial variation of content complexity.
	SpatialAq string `json:"spatialAq,omitempty" yaml:"spatialAq,omitempty"`

	// Produces a bitstream compliant with SMPTE RP-2027.
	Syntax string `json:"syntax,omitempty" yaml:"syntax,omitempty"`

	// Makes adjustments within each frame based on temporal variation of content complexity.
	TemporalAq string `json:"temporalAq,omitempty" yaml:"temporalAq,omitempty"`

	// Framerate numerator.
	FramerateNumerator int `json:"framerateNumerator,omitempty" yaml:"framerateNumerator,omitempty"`

	// Number of B-frames between reference frames.
	GopNumBFrames int `json:"gopNumBFrames,omitempty" yaml:"gopNumBFrames,omitempty"`

	// Number of reference frames to use.
	NumRefFrames int `json:"numRefFrames,omitempty" yaml:"numRefFrames,omitempty"`

	// Scene change detection.
	SceneChangeDetect string `json:"sceneChangeDetect,omitempty" yaml:"sceneChangeDetect,omitempty"`

	// Entropy encoding mode.
	EntropyEncoding string `json:"entropyEncoding,omitempty" yaml:"entropyEncoding,omitempty"`

	// Indicates that AFD values will be written into the output stream.
	AfdSignaling string `json:"afdSignaling,omitempty" yaml:"afdSignaling,omitempty"`

	// Average bitrate in bits/second.
	Bitrate int `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	// Includes color space metadata in the output.
	ColorMetadata string `json:"colorMetadata,omitempty" yaml:"colorMetadata,omitempty"`

	// Four bit AFD value to write on all frames of video in the output stream.
	FixedAfd string `json:"fixedAfd,omitempty" yaml:"fixedAfd,omitempty"`

	//
	FlickerAq string `json:"flickerAq,omitempty" yaml:"flickerAq,omitempty"`

	// Frequency of closed GOPs.
	GopClosedCadence int `json:"gopClosedCadence,omitempty" yaml:"gopClosedCadence,omitempty"`

	// Indicates if the `gop_size` is specified in frames or seconds.
	GopSizeUnits string `json:"gopSizeUnits,omitempty" yaml:"gopSizeUnits,omitempty"`

	// Enables or disables adaptive quantization.
	AdaptiveQuantization string `json:"adaptiveQuantization,omitempty" yaml:"adaptiveQuantization,omitempty"`

	// Amount of lookahead.
	LookAheadRateControl string `json:"lookAheadRateControl,omitempty" yaml:"lookAheadRateControl,omitempty"`

	// Pixel Aspect Ratio denominator.
	ParDenominator int `json:"parDenominator,omitempty" yaml:"parDenominator,omitempty"`

	// H265 level.
	Level string `json:"level,omitempty" yaml:"level,omitempty"`

	// Framerate denominator.
	FramerateDenominator int `json:"framerateDenominator,omitempty" yaml:"framerateDenominator,omitempty"`

	// Pixel Aspect Ratio numerator.
	ParNumerator int `json:"parNumerator,omitempty" yaml:"parNumerator,omitempty"`

	// Number of slices per picture.
	Slices int `json:"slices,omitempty" yaml:"slices,omitempty"`

	// Filters to apply to an encode. See H265 Filter Settings for more details.
	FilterSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH264SettingsFilterSettings `json:"filterSettings,omitempty" yaml:"filterSettings,omitempty"`
}
