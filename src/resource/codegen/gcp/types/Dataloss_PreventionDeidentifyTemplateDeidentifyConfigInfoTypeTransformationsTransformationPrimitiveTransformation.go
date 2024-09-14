package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformation struct {
	/*
	   Pseudonymization method that generates deterministic encryption for the given input. Outputs a base64 encoded representation of the encrypted output. Uses AES-SIV based on the RFC [https://tools.ietf.org/html/rfc5297](https://tools.ietf.org/html/rfc5297).
	   Structure is documented below.
	*/
	CryptoDeterministicConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoDeterministicConfig `json:"cryptoDeterministicConfig,omitempty" yaml:"cryptoDeterministicConfig,omitempty"`

	/*
	   Pseudonymization method that generates surrogates via cryptographic hashing. Uses SHA-256. The key size must be either 32 or 64 bytes.
	   Outputs a base64 encoded representation of the hashed output (for example, L7k0BHmF1ha5U3NfGykjro4xWi1MPVQPjhMAZbSV9mM=).
	   Currently, only string and integer values can be hashed.
	   See https://cloud.google.com/dlp/docs/pseudonymization to learn more.
	   Structure is documented below.
	*/
	CryptoHashConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoHashConfig `json:"cryptoHashConfig,omitempty" yaml:"cryptoHashConfig,omitempty"`

	/*
	   Replaces an identifier with a surrogate using Format Preserving Encryption (FPE) with the FFX mode of operation; however when used in the `content.reidentify` API method, it serves the opposite function by reversing the surrogate back into the original identifier. The identifier must be encoded as ASCII. For a given crypto key and context, the same identifier will be replaced with the same surrogate. Identifiers must be at least two characters long. In the case that the identifier is the empty string, it will be skipped. See [https://cloud.google.com/dlp/docs/pseudonymization](https://cloud.google.com/dlp/docs/pseudonymization) to learn more.
	   Note: We recommend using CryptoDeterministicConfig for all use cases which do not require preserving the input alphabet space and size, plus warrant referential integrity.
	   Structure is documented below.
	*/
	CryptoReplaceFfxFpeConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCryptoReplaceFfxFpeConfig `json:"cryptoReplaceFfxFpeConfig,omitempty" yaml:"cryptoReplaceFfxFpeConfig,omitempty"`

	/*
	   Buckets values based on fixed size ranges. The Bucketing transformation can provide all of this functionality, but requires more configuration. This message is provided as a convenience to the user for simple bucketing strategies.
	   The transformed value will be a hyphenated string of {lower_bound}-{upper_bound}. For example, if lower_bound = 10 and upper_bound = 20, all values that are within this bucket will be replaced with "10-20".
	   This can be used on data of type: double, long.
	   If the bound Value type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing.
	   See https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.
	   Structure is documented below.
	*/
	FixedSizeBucketingConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationFixedSizeBucketingConfig `json:"fixedSizeBucketingConfig,omitempty" yaml:"fixedSizeBucketingConfig,omitempty"`

	// Redact a given value. For example, if used with an InfoTypeTransformation transforming PHONE_NUMBER, and input 'My phone number is 206-555-0123', the output would be 'My phone number is '.
	RedactConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationRedactConfig `json:"redactConfig,omitempty" yaml:"redactConfig,omitempty"`

	/*
	   Replace each input value with a given value.
	   Structure is documented below.
	*/
	ReplaceConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfig `json:"replaceConfig,omitempty" yaml:"replaceConfig,omitempty"`

	/*
	   Replace with a value randomly drawn (with replacement) from a dictionary.
	   Structure is documented below.
	*/
	ReplaceDictionaryConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceDictionaryConfig `json:"replaceDictionaryConfig,omitempty" yaml:"replaceDictionaryConfig,omitempty"`

	/*
	   Generalization function that buckets values based on ranges. The ranges and replacement values are dynamically provided by the user for custom behavior, such as 1-30 > LOW 31-65 > MEDIUM 66-100 > HIGH
	   This can be used on data of type: number, long, string, timestamp.
	   If the provided value type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing.
	   See https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.
	   Structure is documented below.
	*/
	BucketingConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationBucketingConfig `json:"bucketingConfig,omitempty" yaml:"bucketingConfig,omitempty"`

	/*
	   For use with Date, Timestamp, and TimeOfDay, extract or preserve a portion of the value.
	   Structure is documented below.
	*/
	TimePartConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationTimePartConfig `json:"timePartConfig,omitempty" yaml:"timePartConfig,omitempty"`

	// Replace each matching finding with the name of the info type.
	ReplaceWithInfoTypeConfig bool `json:"replaceWithInfoTypeConfig,omitempty" yaml:"replaceWithInfoTypeConfig,omitempty"`

	/*
	   Shifts dates by random number of days, with option to be consistent for the same context. See https://cloud.google.com/dlp/docs/concepts-date-shifting to learn more.
	   Structure is documented below.
	*/
	DateShiftConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationDateShiftConfig `json:"dateShiftConfig,omitempty" yaml:"dateShiftConfig,omitempty"`

	/*
	   Partially mask a string by replacing a given number of characters with a fixed character. Masking can start from the beginning or end of the string. This can be used on data of any type (numbers, longs, and so on) and when de-identifying structured data we'll attempt to preserve the original data's type. (This allows you to take a long like 123 and modify it to a string like --3).
	   Structure is documented below.
	*/
	CharacterMaskConfig Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfig `json:"characterMaskConfig,omitempty" yaml:"characterMaskConfig,omitempty"`
}
