package types

type Rds_InstanceBlueGreenUpdate struct {
	/*
	   Enables low-downtime updates when `true`.
	   Default is `false`.

	   [instance-replication]:
	   https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Replication.html
	   [instance-maintenance]:
	   https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html
	   [blue-green]:
	   https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
