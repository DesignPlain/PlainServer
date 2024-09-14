package types

type Certificateauthority_getAuthorityConfigSubjectConfig struct {
	// The subject alternative name fields.
	SubjectAltNames []Certificateauthority_getAuthorityConfigSubjectConfigSubjectAltName `json:"subjectAltNames,omitempty" yaml:"subjectAltNames,omitempty"`

	// Contains distinguished name fields such as the location and organization.
	Subjects []Certificateauthority_getAuthorityConfigSubjectConfigSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
}
