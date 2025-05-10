package entity

import "time"

type Card struct {
	ID      string      `json:"id"`
	Address interface{} `json:"address"`
	Badges  struct {
		Fogbugz               string      `json:"fogbugz"`
		CheckItems            int         `json:"checkItems"`
		CheckItemsChecked     int         `json:"checkItemsChecked"`
		CheckItemsEarliestDue interface{} `json:"checkItemsEarliestDue"`
		Comments              int         `json:"comments"`
		Attachments           int         `json:"attachments"`
		Description           bool        `json:"description"`
		Due                   interface{} `json:"due"`
		DueComplete           bool        `json:"dueComplete"`
		LastUpdatedByAi       bool        `json:"lastUpdatedByAi"`
		Start                 interface{} `json:"start"`
		AttachmentsByType     struct {
			Trello struct {
				Board int `json:"board"`
				Card  int `json:"card"`
			} `json:"trello"`
		} `json:"attachmentsByType"`
		ExternalSource       interface{} `json:"externalSource"`
		Location             bool        `json:"location"`
		Votes                int         `json:"votes"`
		MaliciousAttachments int         `json:"maliciousAttachments"`
		ViewingMemberVoted   bool        `json:"viewingMemberVoted"`
		Subscribed           bool        `json:"subscribed"`
	} `json:"badges"`
	CheckItemStates []struct {
		IDCheckItem string `json:"idCheckItem"`
		State       string `json:"state"`
	} `json:"checkItemStates"`
	Closed                         bool          `json:"closed"`
	Coordinates                    interface{}   `json:"coordinates"`
	CreationMethod                 interface{}   `json:"creationMethod"`
	CreationMethodError            interface{}   `json:"creationMethodError"`
	CreationMethodLoadingStartedAt interface{}   `json:"creationMethodLoadingStartedAt"`
	DueComplete                    bool          `json:"dueComplete"`
	DateClosed                     interface{}   `json:"dateClosed"`
	DateLastActivity               time.Time     `json:"dateLastActivity"`
	DateCompleted                  time.Time     `json:"dateCompleted"`
	DateViewedByCreator            interface{}   `json:"dateViewedByCreator"`
	Desc                           string        `json:"desc"`
	DescData                       interface{}   `json:"descData"`
	Due                            time.Time     `json:"due"`
	DueReminder                    interface{}   `json:"dueReminder"`
	Email                          string        `json:"email"`
	ExternalSource                 interface{}   `json:"externalSource"`
	IDBoard                        string        `json:"idBoard"`
	IDChecklists                   []string      `json:"idChecklists"`
	IDLabels                       []interface{} `json:"idLabels"`
	IDList                         string        `json:"idList"`
	IDMemberCreator                interface{}   `json:"idMemberCreator"`
	IDMembers                      []interface{} `json:"idMembers"`
	IDMembersVoted                 []interface{} `json:"idMembersVoted"`
	IDOrganization                 string        `json:"idOrganization"`
	IDShort                        int           `json:"idShort"`
	IDAttachmentCover              interface{}   `json:"idAttachmentCover"`
	Labels                         []interface{} `json:"labels"`
	Limits                         struct {
		Attachments struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"attachments"`
		Checklists struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"checklists"`
		Stickers struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"stickers"`
	} `json:"limits"`
	LocationName          interface{} `json:"locationName"`
	ManualCoverAttachment bool        `json:"manualCoverAttachment"`
	Name                  string      `json:"name"`
	NodeID                string      `json:"nodeId"`
	Pinned                bool        `json:"pinned"`
	Pos                   float32     `json:"pos"`
	ShortLink             string      `json:"shortLink"`
	ShortURL              string      `json:"shortUrl"`
	SourceEmail           interface{} `json:"sourceEmail"`
	StaticMapURL          interface{} `json:"staticMapUrl"`
	Start                 interface{} `json:"start"`
	Subscribed            bool        `json:"subscribed"`
	URL                   string      `json:"url"`
	Cover                 struct {
		IDAttachment         interface{} `json:"idAttachment"`
		Color                interface{} `json:"color"`
		IDUploadedBackground interface{} `json:"idUploadedBackground"`
		Size                 string      `json:"size"`
		Brightness           string      `json:"brightness"`
		IDPlugin             interface{} `json:"idPlugin"`
	} `json:"cover"`
	IsTemplate         bool          `json:"isTemplate"`
	CardRole           interface{}   `json:"cardRole"`
	MirrorSourceID     interface{}   `json:"mirrorSourceId"`
	MirrorSourceNodeID interface{}   `json:"mirrorSourceNodeId"`
	Attachments        []interface{} `json:"attachments"`
	PluginData         []interface{} `json:"pluginData"`
	CustomFieldItems   []interface{} `json:"customFieldItems"`

	Checklists []Checklist `json:"checklists"`
}
