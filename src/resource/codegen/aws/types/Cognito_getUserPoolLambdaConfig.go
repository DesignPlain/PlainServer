package types

type Cognito_getUserPoolLambdaConfig struct {
	//
	CustomEmailSenders []Cognito_getUserPoolLambdaConfigCustomEmailSender `json:"customEmailSenders,omitempty" yaml:"customEmailSenders,omitempty"`

	//
	PreTokenGeneration string `json:"preTokenGeneration,omitempty" yaml:"preTokenGeneration,omitempty"`

	//
	VerifyAuthChallengeResponse string `json:"verifyAuthChallengeResponse,omitempty" yaml:"verifyAuthChallengeResponse,omitempty"`

	//
	PostConfirmation string `json:"postConfirmation,omitempty" yaml:"postConfirmation,omitempty"`

	//
	PreSignUp string `json:"preSignUp,omitempty" yaml:"preSignUp,omitempty"`

	//
	CustomSmsSenders []Cognito_getUserPoolLambdaConfigCustomSmsSender `json:"customSmsSenders,omitempty" yaml:"customSmsSenders,omitempty"`

	//
	DefineAuthChallenge string `json:"defineAuthChallenge,omitempty" yaml:"defineAuthChallenge,omitempty"`

	//
	PostAuthentication string `json:"postAuthentication,omitempty" yaml:"postAuthentication,omitempty"`

	//
	UserMigration string `json:"userMigration,omitempty" yaml:"userMigration,omitempty"`

	//
	CreateAuthChallenge string `json:"createAuthChallenge,omitempty" yaml:"createAuthChallenge,omitempty"`

	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	//
	PreTokenGenerationConfigs []Cognito_getUserPoolLambdaConfigPreTokenGenerationConfig `json:"preTokenGenerationConfigs,omitempty" yaml:"preTokenGenerationConfigs,omitempty"`

	//
	CustomMessage string `json:"customMessage,omitempty" yaml:"customMessage,omitempty"`

	//
	PreAuthentication string `json:"preAuthentication,omitempty" yaml:"preAuthentication,omitempty"`
}
