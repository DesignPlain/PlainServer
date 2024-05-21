package dataform

import types "DesignSphere_Server/src/resource/gcp/types"

type RepositoryReleaseConfig struct {
	// A reference to the Dataform repository
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	/*
	   Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	   Structure is documented below.
	*/
	CodeCompilationConfig types.Dataform_RepositoryReleaseConfigCodeCompilationConfig `json:"codeCompilationConfig,omitempty" yaml:"codeCompilationConfig,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule string `json:"cronSchedule,omitempty" yaml:"cronSchedule,omitempty"`

	/*
	   Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.


	   - - -
	*/
	GitCommitish string `json:"gitCommitish,omitempty" yaml:"gitCommitish,omitempty"`

	// The release's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A reference to the region
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
