package types



type Cloudbuild_TriggerBuild struct {
	/*
	   Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
	   Structure is documented below.
	*/
	Artifacts Cloudbuild_TriggerBuildArtifacts `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`

	/*
	   A list of images to be pushed upon the successful completion of all build steps.
	   The images are pushed using the builder service account's credentials.
	   The digests of the pushed images will be stored in the Build resource's results field.
	   If any of the images fail to be pushed, the build status is marked FAILURE.
	*/
	Images []string `json:"images,omitempty" yaml:"images,omitempty"`

	/*
	   Secrets to decrypt using Cloud Key Management Service.
	   Structure is documented below.
	*/
	Secrets []Cloudbuild_TriggerBuildSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	/*
	   The location of the source files to build.
	   One of `storageSource` or `repoSource` must be provided.
	   Structure is documented below.
	*/
	Source Cloudbuild_TriggerBuildSource `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   Amount of time that this build should be allowed to run, to second granularity.
	   If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
	   This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
	   The expected format is the number of seconds followed by s.
	   Default time is ten minutes (600s).
	*/
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   Secrets and secret environment variables.
	   Structure is documented below.
	*/
	AvailableSecrets Cloudbuild_TriggerBuildAvailableSecrets `json:"availableSecrets,omitempty" yaml:"availableSecrets,omitempty"`

	/*
	   Google Cloud Storage bucket where logs should be written.
	   Logs file names will be of the format ${logsBucket}/log-${build_id}.txt.
	*/
	LogsBucket string `json:"logsBucket,omitempty" yaml:"logsBucket,omitempty"`

	/*
	   Special options for this build.
	   Structure is documented below.
	*/
	Options Cloudbuild_TriggerBuildOptions `json:"options,omitempty" yaml:"options,omitempty"`

	/*
	   TTL in queue for this build. If provided and the build is enqueued longer than this value,
	   the build will expire and the build status will be EXPIRED.
	   The TTL starts ticking from createTime.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	QueueTtl string `json:"queueTtl,omitempty" yaml:"queueTtl,omitempty"`

	/*
	   The operations to be performed on the workspace.
	   Structure is documented below.
	*/
	Steps []Cloudbuild_TriggerBuildStep `json:"steps,omitempty" yaml:"steps,omitempty"`

	// Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty" yaml:"substitutions,omitempty"`

	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
