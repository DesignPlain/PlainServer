package types

type Appstream_getImageApplication struct {
	// The app block ARN of the application.
	AppBlockArn string `json:"appBlockArn,omitempty" yaml:"appBlockArn,omitempty"`

	// Bool based on if the application is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// URL of the application icon. This URL may be time-limited.
	IconUrl string `json:"iconUrl,omitempty" yaml:"iconUrl,omitempty"`

	// Arguments that are passed to the application at it's launch.
	LaunchParameters string `json:"launchParameters,omitempty" yaml:"launchParameters,omitempty"`

	/*
	   String to string map that contains additional attributes used to describe the application.
	   - `Name` - Name of the application.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   Array of strings describing the platforms on which the application can run.
	   Values will be from: WINDOWS | WINDOWS_SERVER_2016 | WINDOWS_SERVER_2019 | WINDOWS_SERVER_2022 | AMAZON_LINUX2
	*/
	Platforms []string `json:"platforms,omitempty" yaml:"platforms,omitempty"`

	// Arn of the image being searched for. Cannot be used with name_regex or name.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Image name to display.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// List of the instance families of the application.
	InstanceFamilies []string `json:"instanceFamilies,omitempty" yaml:"instanceFamilies,omitempty"`

	// Path to the application's excecutable in the instance.
	LaunchPath string `json:"launchPath,omitempty" yaml:"launchPath,omitempty"`

	// Description of image.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Working directory for the application.
	WorkingDirectory string `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`

	// Time at which this image was created.
	CreatedTime string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`

	// A list named icon_s3_location that contains the following:
	IconS3Locations []Appstream_getImageApplicationIconS3Location `json:"iconS3Locations,omitempty" yaml:"iconS3Locations,omitempty"`

	// Name of the image being searched for. Cannot be used with name_regex or arn.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
