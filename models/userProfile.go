package models

import "gorm.io/gorm"

type UserProfile struct {
	Id                int `json:"id"`
	AccountState      string
	Dni               string
	Uid               string
	Cn                string
	GivenName         string
	Sn                string
	Title             string
	Initials          string
	Description       string
	UserType          string
	Roll              string
	MailEnabled       string
	ProxyEnabled      string
	RemoteEnabled     string
	MailAccessIn      string
	MailAccessOut     string
	MailQuota         string
	ProxyAccess       string
	ProxyQuota        string
	RemotePhoneNumber string
	WifiEnabled       string
	DeviceNumber      string
	NextcloudEnabled  string
	NextcloudQuota    string
	Telegramuid       string
	DhcpHWAddress     []string
	//--Date ASSETS
	AssetsArea             string
	AssetsDepartmentNumber string
	AssetsCategory         string
	AssetsPosition         string
	AssetsCategoryName     string
	AssetsSubcategoryName  string
	AssetsProfession       string

	//--Date Sigenu
	Country                string
	StudentType            string
	Carrer                 string
	Faculty                string
	CourseType             string
	ScholasticOrigin       string
	TownUniversity         string
	MatriculationEndDate   string
	ReMatriculationEndDate string
	StudentStatus          string
	StudentYear            string

	CreateUser      string
	CreateDate      string
	ModifyUser      string
	ModifyData      string
	UserPasswordSet string

	PassValid string
	PassSet   string

	ExpireDate  string
	HashModLdap string

	HashModApi   string
	SyncRequired bool

	SecurityCode string
	Ou           string
}

type PersonalInformation struct {
	Dni           string `json:"dni"`
	Cn            string `json:"cn"`
	GivenName     string `json:"given_name"`
	Sn            string `json:"sn"`
	PersonalPhoto string `json:"personal_photo"`
	Overlapping   string `json:"overlapping"`
}

type AccountInfo struct {
	UserType             string `json:"user_type"`
	CreateUser           string `json:"create_user"`
	CreateDate           string `json:"create_date"`
	ModifyUser           string `json:"modify_user"`
	ModifyData           string `json:"modify_data"`
	AcceptSystemPolicies bool   `json:"accept_system_policies"`
	Password             struct {
		UserPasswordSet string `json:"user_password_set"`
		PassValid       string `json:"pass_valid"`
		PassSet         string `json:"pass_set"`
	} `json:"password"`
}

type ActiveUser struct {
	Status              int                 `json:"status"`
	AccountState        string              `json:"account_state"`
	Uid                 string              `json:"uid"`
	PersonalInformation PersonalInformation `json:"personal_information"`
	AccountInfo         AccountInfo         `json:"account_info"`
}

/*
Model that stores all the data related to the user account
*/
type User struct {
	gorm.Model
	Username             string `gorm:"uniqueIndex:idx_username,length:255" json:"username"`
	AcceptSystemPolicies bool   `json:"accept_system_policies"`
}
