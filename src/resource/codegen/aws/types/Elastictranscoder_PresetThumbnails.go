package types

type Elastictranscoder_PresetThumbnails struct {
	// The format of thumbnails, if any. Valid formats are jpg and png.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// The approximate number of seconds between thumbnails. The value must be an integer. The actual interval can vary by several seconds from one thumbnail to the next.
	Interval string `json:"interval,omitempty" yaml:"interval,omitempty"`

	// The maximum height of thumbnails, in pixels. If you specify auto, Elastic Transcoder uses 1080 (Full HD) as the default value. If you specify a numeric value, enter an even integer between 32 and 3072, inclusive.
	MaxHeight string `json:"maxHeight,omitempty" yaml:"maxHeight,omitempty"`

	// The maximum width of thumbnails, in pixels. If you specify auto, Elastic Transcoder uses 1920 (Full HD) as the default value. If you specify a numeric value, enter an even integer between 32 and 4096, inclusive.
	MaxWidth string `json:"maxWidth,omitempty" yaml:"maxWidth,omitempty"`

	// When you set PaddingPolicy to Pad, Elastic Transcoder might add black bars to the top and bottom and/or left and right sides of thumbnails to make the total size of the thumbnails match the values that you specified for thumbnail MaxWidth and MaxHeight settings.
	PaddingPolicy string `json:"paddingPolicy,omitempty" yaml:"paddingPolicy,omitempty"`

	// The width and height of thumbnail files in pixels, in the format WidthxHeight, where both values are even integers. The values cannot exceed the width and height that you specified in the Video:Resolution object. (To better control resolution and aspect ratio of thumbnails, we recommend that you use the thumbnail values `max_width`, `max_height`, `sizing_policy`, and `padding_policy` instead of `resolution` and `aspect_ratio`. The two groups of settings are mutually exclusive. Do not use them together)
	Resolution string `json:"resolution,omitempty" yaml:"resolution,omitempty"`

	// A value that controls scaling of thumbnails. Valid values are: `Fit`, `Fill`, `Stretch`, `Keep`, `ShrinkToFit`, and `ShrinkToFill`.
	SizingPolicy string `json:"sizingPolicy,omitempty" yaml:"sizingPolicy,omitempty"`

	// The aspect ratio of thumbnails. The following values are valid: auto, 1:1, 4:3, 3:2, 16:9
	AspectRatio string `json:"aspectRatio,omitempty" yaml:"aspectRatio,omitempty"`
}
