package asset

import "encoding/json"

type Asset struct {
	FileAsset   string `yaml:"fn::fileAsset,omitempty"`
	StringAsset string `yaml:"fn::stringAsset,omitempty"`
	RemoteAsset string `yaml:"fn::remoteAsset,omitempty"`
}

/* Custom Marshaller for FileAsset */
func (s Asset) MarshalJSON() ([]byte, error) {
	if s.FileAsset != "" {
		return json.Marshal(s.FileAsset)
	} else if s.RemoteAsset != "" {
		return json.Marshal(s.RemoteAsset)
	} else {
		return json.Marshal(s.StringAsset)
	}

}

/* Custom Unmarshaller for FileAsset */
func (s *Asset) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	s.FileAsset = str
	return nil
}
