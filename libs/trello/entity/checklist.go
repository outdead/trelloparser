package entity

type Checklist struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	IDBoard string  `json:"idBoard"`
	IDCard  string  `json:"idCard"`
	Pos     float32 `json:"pos"`
	Limits  struct {
		CheckItems struct {
			PerChecklist struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perChecklist"`
		} `json:"checkItems"`
	} `json:"limits"`
	CheckItems []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		NameData    interface{} `json:"nameData"`
		Pos         float32     `json:"pos"`
		State       string      `json:"state"`
		Due         interface{} `json:"due"`
		DueReminder interface{} `json:"dueReminder"`
		IDMember    interface{} `json:"idMember"`
		IDChecklist string      `json:"idChecklist"`
	} `json:"checkItems"`
	CreationMethod interface{} `json:"creationMethod"`
}
