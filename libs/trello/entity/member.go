package entity

type Member struct {
	ID              string `json:"id"`
	AaID            string `json:"aaId"`
	ActivityBlocked bool   `json:"activityBlocked"`
	AvatarHash      string `json:"avatarHash"`
	AvatarURL       string `json:"avatarUrl"`
	Bio             string `json:"bio"`
	BioData         struct {
		Emoji struct{} `json:"emoji"`
	} `json:"bioData"`
	Confirmed                bool          `json:"confirmed"`
	FullName                 string        `json:"fullName"`
	IDEnterprise             interface{}   `json:"idEnterprise"`
	IDEnterprisesDeactivated []interface{} `json:"idEnterprisesDeactivated"`
	IDMemberReferrer         interface{}   `json:"idMemberReferrer"`
	IDPremOrgsAdmin          []interface{} `json:"idPremOrgsAdmin"`
	Initials                 string        `json:"initials"`
	MemberType               string        `json:"memberType"`
	NonPublic                struct {
		FullName   string `json:"fullName"`
		Initials   string `json:"initials"`
		AvatarURL  string `json:"avatarUrl"`
		AvatarHash string `json:"avatarHash"`
	} `json:"nonPublic"`
	NonPublicAvailable bool          `json:"nonPublicAvailable"`
	Products           []interface{} `json:"products"`
	URL                string        `json:"url"`
	Username           string        `json:"username"`
	Status             string        `json:"status"`
}
