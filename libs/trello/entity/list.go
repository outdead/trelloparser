package entity

type List struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Closed     bool        `json:"closed"`
	Color      interface{} `json:"color"`
	IDBoard    string      `json:"idBoard"`
	Pos        float32     `json:"pos"`
	Subscribed bool        `json:"subscribed"`
	SoftLimit  interface{} `json:"softLimit"`
	Type       interface{} `json:"type"`
	Datasource struct {
		Filter bool `json:"filter"`
	} `json:"datasource"`
	CreationMethod interface{} `json:"creationMethod"`
	IDOrganization string      `json:"idOrganization"`
	Limits         struct {
		Cards struct {
			OpenPerList struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"openPerList"`
			TotalPerList struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerList"`
		} `json:"cards"`
	} `json:"limits"`
	NodeID string `json:"nodeId"`

	Cards []Card `json:"cards"`
}
