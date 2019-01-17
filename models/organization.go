package models

// in DB
type Organization struct {
	OrgId         int      `bson:"_id"`
	OrgTitle      string   `bson:"OrgTitle"`
	OrgWebsite    string   `bson:"OrgWebsite"`
	StreetAddress string   `bson:"StreetAddress"`
	City          string   `bson:"City"`
	State         string   `bson:"State"`
	ZipCode       string   `bson:"ZipCode"`
	Phone         string   `bson:"Phone"`
	Email         string   `bson:"Email"`
	ActivityDesc  string   `bson:"ActivityDesc"`
	Lat           float64  `bson:"Lat"`
	Long          float64  `bson:"Long"`
	HasShadow     bool     `bson:"HasShadow"`
	HasCost       bool     `bson:"HasCost"`
	HasTransport  bool     `bson:"HasTransport"`
	Under18       bool     `bson:"Under18"`
	CareerEmp     []string `bson:"CareerEmp"`
	GradeLevels   []int    `bson:"GradeLevels"`
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
