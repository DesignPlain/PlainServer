package opsworks

type RdsDbInstance struct {
	// A db password
	DbPassword string `json:"dbPassword,omitempty" yaml:"dbPassword,omitempty"`

	// A db username
	DbUser string `json:"dbUser,omitempty" yaml:"dbUser,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn string `json:"rdsDbInstanceArn,omitempty" yaml:"rdsDbInstanceArn,omitempty"`

	// The stack to register a db instance for. Changing this will force a new resource.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`
}
