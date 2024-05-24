package types

type Cognito_UserPoolLambdaConfig struct {
	// Post-confirmation AWS Lambda trigger.
	PostConfirmation string `json:"postConfirmation,omitempty" yaml:"postConfirmation,omitempty"`

	// Pre-authentication AWS Lambda trigger.
	PreAuthentication string `json:"preAuthentication,omitempty" yaml:"preAuthentication,omitempty"`

	// Allow to customize identity token claims before token generation.
	PreTokenGeneration string `json:"preTokenGeneration,omitempty" yaml:"preTokenGeneration,omitempty"`

	// Verifies the authentication challenge response.
	VerifyAuthChallengeResponse string `json:"verifyAuthChallengeResponse,omitempty" yaml:"verifyAuthChallengeResponse,omitempty"`

	// A custom email sender AWS Lambda trigger. See custom_email_sender Below.
	CustomEmailSender Cognito_UserPoolLambdaConfigCustomEmailSender `json:"customEmailSender,omitempty" yaml:"customEmailSender,omitempty"`

	// Custom Message AWS Lambda trigger.
	CustomMessage string `json:"customMessage,omitempty" yaml:"customMessage,omitempty"`

	// A custom SMS sender AWS Lambda trigger. See custom_sms_sender Below.
	CustomSmsSender Cognito_UserPoolLambdaConfigCustomSmsSender `json:"customSmsSender,omitempty" yaml:"customSmsSender,omitempty"`

	// Post-authentication AWS Lambda trigger.
	PostAuthentication string `json:"postAuthentication,omitempty" yaml:"postAuthentication,omitempty"`

	// Pre-registration AWS Lambda trigger.
	PreSignUp string `json:"preSignUp,omitempty" yaml:"preSignUp,omitempty"`

	// User migration Lambda config type.
	UserMigration string `json:"userMigration,omitempty" yaml:"userMigration,omitempty"`

	// ARN of the lambda creating an authentication challenge.
	CreateAuthChallenge string `json:"createAuthChallenge,omitempty" yaml:"createAuthChallenge,omitempty"`

	// Defines the authentication challenge.
	DefineAuthChallenge string `json:"defineAuthChallenge,omitempty" yaml:"defineAuthChallenge,omitempty"`

	// The Amazon Resource Name of Key Management Service Customer master keys. Amazon Cognito uses the key to encrypt codes and temporary passwords sent to CustomEmailSender and CustomSMSSender.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
