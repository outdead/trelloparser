package entity

type LabelNames struct {
	Green       string `json:"green"`
	Yellow      string `json:"yellow"`
	Orange      string `json:"orange"`
	Red         string `json:"red"`
	Purple      string `json:"purple"`
	Blue        string `json:"blue"`
	Sky         string `json:"sky"`
	Lime        string `json:"lime"`
	Pink        string `json:"pink"`
	Black       string `json:"black"`
	GreenDark   string `json:"green_dark"`
	YellowDark  string `json:"yellow_dark"`
	OrangeDark  string `json:"orange_dark"`
	RedDark     string `json:"red_dark"`
	PurpleDark  string `json:"purple_dark"`
	BlueDark    string `json:"blue_dark"`
	SkyDark     string `json:"sky_dark"`
	LimeDark    string `json:"lime_dark"`
	PinkDark    string `json:"pink_dark"`
	BlackDark   string `json:"black_dark"`
	GreenLight  string `json:"green_light"`
	YellowLight string `json:"yellow_light"`
	OrangeLight string `json:"orange_light"`
	RedLight    string `json:"red_light"`
	PurpleLight string `json:"purple_light"`
	BlueLight   string `json:"blue_light"`
	SkyLight    string `json:"sky_light"`
	LimeLight   string `json:"lime_light"`
	PinkLight   string `json:"pink_light"`
	BlackLight  string `json:"black_light"`
}

type Label struct {
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}
