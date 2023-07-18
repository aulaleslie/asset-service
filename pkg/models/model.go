package models

type AssetsGroup struct {
	AgrID            int    `json:"agr_id" gorm:"primaryKey"`
	AgrGroupName     string `json:"agr_group_name"`
	AgrBusinessGroup string `json:"agr_business_group"`
	AgrOrganization  int    `json:"agr_organization"`
}

type AssetSubGroup struct {
	AsgID           int    `json:"asg_id" gorm:"primaryKey"`
	AsgName         string `json:"asg_name"`
	AsgParentGroup  string `json:"asg_parent_group"`
	AsgOrganization int    `json:"asg_organization"`
}

type BusinessGroup struct {
	BgpID           int    `json:"bgp_id" gorm:"primaryKey"`
	BgpLabel        string `json:"bgp_label"`
	BgpName         string `json:"bgp_name" gorm:"not null"`
	BgpOrganization int    `json:"bgp_organization" gorm:"not null"`
}

type User struct {
	UsrID              int    `json:"usr_id" gorm:"primaryKey"`
	UsrName            string `json:"usr_name"`
	UsrEmail           string `json:"usr_email"`
	UsrPhoneNumber     string `json:"usr_phone_number"`
	UsrMobileNumber    string `json:"usr_mobile_number"`
	UsrOrganization    int    `json:"usr_organization"`
	UsrUsername        string `json:"usr_username"`
	UsrPassword        string `json:"usr_password"`
	UsrJWTAccessToken  string `json:"usr_jwt_access_token"`
	UsrAPIOrganization int    `json:"usr_api_organization"`
}

func (AssetsGroup) TableName() string {
	return "assets_group"
}
