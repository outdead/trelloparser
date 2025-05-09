package entity

import "time"

type Action struct {
	ID              string `json:"id"`
	IDMemberCreator string `json:"idMemberCreator"`
	Data            struct {
		Card struct {
			Due       time.Time `json:"due"`
			ID        string    `json:"id"`
			Name      string    `json:"name"`
			IDShort   int       `json:"idShort"`
			ShortLink string    `json:"shortLink"`
		} `json:"card"`
		Old struct {
			Due interface{} `json:"due"`
		} `json:"old"`
		Board struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
		} `json:"board"`
		List struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"list"`
	} `json:"data"`
	AppCreator    interface{} `json:"appCreator"`
	Type          string      `json:"type"`
	Date          time.Time   `json:"date"`
	Limits        interface{} `json:"limits"`
	MemberCreator struct {
		ID               string      `json:"id"`
		ActivityBlocked  bool        `json:"activityBlocked"`
		AvatarHash       string      `json:"avatarHash"`
		AvatarURL        string      `json:"avatarUrl"`
		FullName         string      `json:"fullName"`
		IDMemberReferrer interface{} `json:"idMemberReferrer"`
		Initials         string      `json:"initials"`
		NonPublic        struct {
			FullName   string `json:"fullName"`
			Initials   string `json:"initials"`
			AvatarURL  string `json:"avatarUrl"`
			AvatarHash string `json:"avatarHash"`
		} `json:"nonPublic"`
		NonPublicAvailable bool   `json:"nonPublicAvailable"`
		Username           string `json:"username"`
	} `json:"memberCreator"`
}
