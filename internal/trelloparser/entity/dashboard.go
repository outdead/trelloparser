package entity

import "time"

type Dashboard struct {
	ID                string        `json:"id"`
	NodeID            string        `json:"nodeId"`
	Name              string        `json:"name"`
	Desc              string        `json:"desc"`
	DescData          interface{}   `json:"descData"`
	Closed            bool          `json:"closed"`
	DateClosed        interface{}   `json:"dateClosed"`
	IDOrganization    string        `json:"idOrganization"`
	IDEnterprise      interface{}   `json:"idEnterprise"`
	Limits            Limits        `json:"limits"`
	Pinned            bool          `json:"pinned"`
	Starred           bool          `json:"starred"`
	URL               string        `json:"url"`
	Prefs             Prefs         `json:"prefs"`
	ShortLink         string        `json:"shortLink"`
	Subscribed        bool          `json:"subscribed"`
	LabelNames        LabelNames    `json:"labelNames"`
	PowerUps          []interface{} `json:"powerUps"`
	DateLastActivity  time.Time     `json:"dateLastActivity"`
	DateLastView      time.Time     `json:"dateLastView"`
	ShortURL          string        `json:"shortUrl"`
	IDTags            []interface{} `json:"idTags"`
	DatePluginDisable interface{}   `json:"datePluginDisable"`
	CreationMethod    interface{}   `json:"creationMethod"`
	IxUpdate          string        `json:"ixUpdate"`
	TemplateGallery   interface{}   `json:"templateGallery"`
	EnterpriseOwned   bool          `json:"enterpriseOwned"`
	IDBoardSource     interface{}   `json:"idBoardSource"`
	PremiumFeatures   []string      `json:"premiumFeatures"`
	IDMemberCreator   string        `json:"idMemberCreator"`
	Type              interface{}   `json:"type"`
	Actions           []Action      `json:"actions"`
	Cards             []Card        `json:"cards"`
	Labels            []Label       `json:"labels"`
	Lists             []List        `json:"lists"`
	Members           []Member      `json:"members"`
	Checklists        []Checklist   `json:"checklists"`
	CustomFields      []interface{} `json:"customFields"`
	Memberships       []Membership  `json:"memberships"`
	PluginData        []interface{} `json:"pluginData"`
}
