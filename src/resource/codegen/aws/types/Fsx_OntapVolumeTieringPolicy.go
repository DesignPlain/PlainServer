package types

type Fsx_OntapVolumeTieringPolicy struct {
	// Specifies the number of days that user data in a volume must remain inactive before it is considered "cold" and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.
	CoolingPeriod int `json:"coolingPeriod,omitempty" yaml:"coolingPeriod,omitempty"`

	// Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
