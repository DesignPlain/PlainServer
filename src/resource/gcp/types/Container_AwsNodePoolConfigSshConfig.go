package types

type Container_AwsNodePoolConfigSshConfig struct {
	// The name of the EC2 key pair used to login into cluster machines.
	Ec2KeyPair string `json:"ec2KeyPair,omitempty" yaml:"ec2KeyPair,omitempty"`
}
