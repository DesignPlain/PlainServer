package types

type Dataproc_ClusterClusterConfigInitializationAction struct {
	/*
	   The script to be executed during initialization of the cluster.
	   The script must be a GCS file with a gs:// prefix.
	*/
	Script string `json:"script,omitempty" yaml:"script,omitempty"`

	/*
	   The maximum duration (in seconds) which `script` is
	   allowed to take to execute its action. GCP will default to a predetermined
	   computed value if not set (currently 300).

	   - - -
	*/
	TimeoutSec int `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`
}
