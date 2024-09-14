package types

type Notebooks_RuntimeSoftwareConfig struct {
	// Verifies core internal services are running. Default: True.
	EnableHealthMonitoring bool `json:"enableHealthMonitoring,omitempty" yaml:"enableHealthMonitoring,omitempty"`

	/*
	   Time in minutes to wait before shuting down runtime.
	   Default: 180 minutes
	*/
	IdleShutdownTimeout int `json:"idleShutdownTimeout,omitempty" yaml:"idleShutdownTimeout,omitempty"`

	// Install Nvidia Driver automatically.
	InstallGpuDriver bool `json:"installGpuDriver,omitempty" yaml:"installGpuDriver,omitempty"`

	/*
	   Use a list of container images to use as Kernels in the notebook instance.
	   Structure is documented below.
	*/
	Kernels []Notebooks_RuntimeSoftwareConfigKernel `json:"kernels,omitempty" yaml:"kernels,omitempty"`

	/*
	   (Output)
	   Bool indicating whether an newer image is available in an image family.
	*/
	Upgradeable bool `json:"upgradeable,omitempty" yaml:"upgradeable,omitempty"`

	/*
	   Specify a custom Cloud Storage path where the GPU driver is stored.
	   If not specified, we'll automatically choose from official GPU drivers.
	*/
	CustomGpuDriverPath string `json:"customGpuDriverPath,omitempty" yaml:"customGpuDriverPath,omitempty"`

	/*
	   Runtime will automatically shutdown after idle_shutdown_time.
	   Default: True
	*/
	IdleShutdown bool `json:"idleShutdown,omitempty" yaml:"idleShutdown,omitempty"`

	/*
	   Cron expression in UTC timezone for schedule instance auto upgrade.
	   Please follow the [cron format](https://en.wikipedia.org/wiki/Cron).
	*/
	NotebookUpgradeSchedule string `json:"notebookUpgradeSchedule,omitempty" yaml:"notebookUpgradeSchedule,omitempty"`

	/*
	   Path to a Bash script that automatically runs after a notebook instance
	   fully boots up. The path must be a URL or
	   Cloud Storage path (gs://path-to-file/file-name).
	*/
	PostStartupScript string `json:"postStartupScript,omitempty" yaml:"postStartupScript,omitempty"`

	/*
	   Behavior for the post startup script.
	   Possible values are: `POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED`, `RUN_EVERY_START`, `DOWNLOAD_AND_RUN_EVERY_START`.
	*/
	PostStartupScriptBehavior string `json:"postStartupScriptBehavior,omitempty" yaml:"postStartupScriptBehavior,omitempty"`
}
