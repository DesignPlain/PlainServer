package types

type Fsx_FileCacheLustreConfiguration struct {
	// The configuration for a Lustre MDT (Metadata Target) storage volume. See the `metadata_configuration` block.
	MetadataConfigurations []Fsx_FileCacheLustreConfigurationMetadataConfiguration `json:"metadataConfigurations,omitempty" yaml:"metadataConfigurations,omitempty"`

	//
	MountName string `json:"mountName,omitempty" yaml:"mountName,omitempty"`

	// Provisions the amount of read and write throughput for each 1 tebibyte (TiB) of cache storage capacity, in MB/s/TiB. The only supported value is `1000`.
	PerUnitStorageThroughput int `json:"perUnitStorageThroughput,omitempty" yaml:"perUnitStorageThroughput,omitempty"`

	// A recurring weekly time, in the format `D:HH:MM`. `D` is the day of the week, for which `1` represents Monday and `7` represents Sunday. `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, 1:05:00 specifies maintenance at 5 AM Monday. See the [ISO week date](https://en.wikipedia.org/wiki/ISO_week_date) for more information.
	WeeklyMaintenanceStartTime string `json:"weeklyMaintenanceStartTime,omitempty" yaml:"weeklyMaintenanceStartTime,omitempty"`

	// Specifies the cache deployment type. The only supported value is `CACHE_1`.
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	//
	LogConfigurations []Fsx_FileCacheLustreConfigurationLogConfiguration `json:"logConfigurations,omitempty" yaml:"logConfigurations,omitempty"`
}
