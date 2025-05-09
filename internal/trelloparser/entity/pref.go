package entity

type Prefs struct {
	PermissionLevel          string        `json:"permissionLevel"`
	HideVotes                bool          `json:"hideVotes"`
	Voting                   string        `json:"voting"`
	Comments                 string        `json:"comments"`
	Invitations              string        `json:"invitations"`
	SelfJoin                 bool          `json:"selfJoin"`
	CardCovers               bool          `json:"cardCovers"`
	ShowCompleteStatus       bool          `json:"showCompleteStatus"`
	CardCounts               bool          `json:"cardCounts"`
	IsTemplate               bool          `json:"isTemplate"`
	CardAging                string        `json:"cardAging"`
	CalendarFeedEnabled      bool          `json:"calendarFeedEnabled"`
	HiddenPluginBoardButtons []interface{} `json:"hiddenPluginBoardButtons"`
	SwitcherViews            []struct {
		ViewType string `json:"viewType"`
		Enabled  bool   `json:"enabled"`
	} `json:"switcherViews"`
	AutoArchive           interface{} `json:"autoArchive"`
	Background            string      `json:"background"`
	BackgroundColor       interface{} `json:"backgroundColor"`
	BackgroundDarkColor   interface{} `json:"backgroundDarkColor"`
	BackgroundImage       string      `json:"backgroundImage"`
	BackgroundDarkImage   interface{} `json:"backgroundDarkImage"`
	BackgroundImageScaled []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"backgroundImageScaled"`
	BackgroundTile        bool   `json:"backgroundTile"`
	BackgroundBrightness  string `json:"backgroundBrightness"`
	SharedSourceURL       string `json:"sharedSourceUrl"`
	BackgroundBottomColor string `json:"backgroundBottomColor"`
	BackgroundTopColor    string `json:"backgroundTopColor"`
	CanBePublic           bool   `json:"canBePublic"`
	CanBeEnterprise       bool   `json:"canBeEnterprise"`
	CanBeOrg              bool   `json:"canBeOrg"`
	CanBePrivate          bool   `json:"canBePrivate"`
	CanInvite             bool   `json:"canInvite"`
}
