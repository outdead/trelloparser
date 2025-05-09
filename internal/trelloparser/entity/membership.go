package entity

type Membership struct {
	ID          string `json:"id"`
	IDMember    string `json:"idMember"`
	MemberType  string `json:"memberType"`
	Unconfirmed bool   `json:"unconfirmed"`
	Deactivated bool   `json:"deactivated"`
}
