package orgs

// Organization represents information for a new org
type Organization struct {
	OrgId         int      `bson:"OrgId"`
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

// OrgInfo represents organization information for buiding searching criteria
type OrgInfo struct {
	SearchContent string   `json: "SearchContent"`
	HasShadow     bool     `json: "HasShadow"`
	HasCost       bool     `json: "HasCost"`
	HasTransport  bool     `json: "HasTransport"`
	Under18       bool     `json: "Under18"`
	CareerEmp     []string `json: "CareerEmp"`
	GradeLevels   []int    `json: "GradeLevels"`
}
