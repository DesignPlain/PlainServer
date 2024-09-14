package types

type Sql_getDatabaseInstancesInstanceSettingAdvancedMachineFeature struct {
	// The number of threads per physical core. Can be 1 or 2.
	ThreadsPerCore int `json:"threadsPerCore,omitempty" yaml:"threadsPerCore,omitempty"`
}
