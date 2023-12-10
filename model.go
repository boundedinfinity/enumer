package enumer

type EnumData struct {
	Type        string        `json:"type,omitempty" yaml:"type,omitempty"`
	Struct      string        `json:"struct,omitempty" yaml:"struct,omitempty"`
	Package     string        `json:"package,omitempty" yaml:"package,omitempty"`
	InputPath   string        `json:"input-path,omitempty" yaml:"input-path,omitempty"`
	OutputPath  string        `json:"output-path,omitempty" yaml:"output-path,omitempty"`
	Desc        string        `json:"desc" yaml,omitempty:"desc,omitempty"`
	Header      string        `json:"header,omitempty" yaml:"header,omitempty"`
	HeaderFrom  string        `json:"header-from,omitempty" yaml:"header-from,omitempty"`
	HeaderLines []string      `json:"header-lines,omitempty" yaml:"header-lines,omitempty"`
	SkipFormat  bool          `json:"skip-format,omitempty" yaml:"skip-format,omitempty"`
	Debug       bool          `json:"debug,omitempty" yaml:"debug,omitempty"`
	Overwrite   bool          `json:"overwrite,omitempty" yaml:"overwrite,omitempty"`
	Serialize   EnumSerialize `json:"serialize,omitempty" yaml:"serialize,omitempty"`
	Values      []EnumValue   `json:"values,omitempty" yaml:"values,omitempty"`
}

type EnumSerialize struct {
	Type  string `json:"type,omitempty" yaml:"type,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type EnumValue struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Serialized string `json:"serialized,omitempty" yaml:"serialized,omitempty"`
}
