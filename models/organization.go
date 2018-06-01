package models

type Organization struct {
	// OrganizationId bson.ObjectId `json:"OrganizationId" bson:"_id"`
	OrgId         int      `json:"OrgID"`
	OrgTitle      string   `json:"OrgTitle"`
	OrgWebsite    string   `json:"OrgWebsite"`
	StreetAddress string   `json: "StreetAddress"`
	City          string   `json: "City"`
	State         string   `json: "state"`
	ZipCode       string   `json: "ZipCode"`
	Phone         string   `json: "Phon"`
	Email         string   `json: "Email"`
	ActivityDesc  string   `json: "ActivityDesc"`
	Lat           float64  `json: "Lat"`
	Long          float64  `json: "Long"`
	HasShadow     bool     `json: "HasShadow"`
	HasCost       bool     `json: "HasCost"`
	HasTransport  bool     `json: "HasTransport"`
	Under18       bool     `json: "Under18"`
	CareerEmp     []string `json: "CareerEmp"`
	GradeLevels   []int    `json: "GradeLevels"`
}

type UpdateOrganization struct {
	OrgTitle      string   `json:"OrgTitle"`
	OrgWebsite    string   `json:"OrgWebsite"`
	StreetAddress string   `json: "StreetAddress"`
	City          string   `json: "City"`
	State         string   `json: "State"`
	ZipCode       string   `json: "ZipCode"`
	Phone         string   `json: "Phon"`
	Email         string   `json: "Email"`
	ActivityDesc  string   `json: "ActivityDesc"`
	Lat           float64  `json: "Lat"`
	Long          float64  `json: "Long"`
	HasShadow     bool     `json: "HasShadow"`
	HasCost       bool     `json: "HasCost"`
	HasTransport  bool     `json: "HasTransport"`
	Under18       bool     `json: "Under18"`
	CareerEmp     []string `json: "CareerEmp"`
	GradeLevels   []int    `json: "GradeLevels"`
}
