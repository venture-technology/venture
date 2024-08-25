package entity

type Child struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name"`
	RG          string      `json:"rg"`
	Responsible Responsible `json:"responsible"`
	Shift       string      `json:"shift" validate:"oneof=matutino vespertino noturno"`
}
