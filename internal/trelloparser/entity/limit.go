package entity

type Limits struct {
	Attachments struct {
		PerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perBoard"`
		PerCard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perCard"`
	} `json:"attachments"`
	Boards struct {
		TotalMembersPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"totalMembersPerBoard"`
		TotalAccessRequestsPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"totalAccessRequestsPerBoard"`
	} `json:"boards"`
	Cards struct {
		OpenPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"openPerBoard"`
		OpenPerList struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"openPerList"`
		TotalPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"totalPerBoard"`
		TotalPerList struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"totalPerList"`
	} `json:"cards"`
	Checklists struct {
		PerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perBoard"`
		PerCard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perCard"`
	} `json:"checklists"`
	CheckItems struct {
		PerChecklist struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perChecklist"`
	} `json:"checkItems"`
	CustomFields struct {
		PerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perBoard"`
	} `json:"customFields"`
	CustomFieldOptions struct {
		PerField struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perField"`
	} `json:"customFieldOptions"`
	Labels struct {
		PerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perBoard"`
	} `json:"labels"`
	Lists struct {
		OpenPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"openPerBoard"`
		TotalPerBoard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"totalPerBoard"`
	} `json:"lists"`
	Stickers struct {
		PerCard struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perCard"`
	} `json:"stickers"`
	Reactions struct {
		PerAction struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"perAction"`
		UniquePerAction struct {
			Status    string `json:"status"`
			DisableAt int    `json:"disableAt"`
			WarnAt    int    `json:"warnAt"`
		} `json:"uniquePerAction"`
	} `json:"reactions"`
}
