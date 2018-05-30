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

// [{"OrgID": 1, "OrgTitle": "Pre-Health Professions Advising Program",
// "OrgWebsite": "www.uidaho.edu/pre-health", "StreetAddress": "875 Perimeter Drive",
// "City": "Moscow", "State": "Idaho", "ZipCode": "83844-2436", "Phone": "(208) 885-5809",
//  "Email": "pre-health@uidaho.edu",
//  "ActivityDesc": "The Pre-Health Advising Program at the University of Idaho serves as a resource
//  for students and alumni, from all majors, who are exploring graduate and professional programs in healthcare.
//    Services provided include assistance with career exploration, prerequisite course sequencing, advice for
//     resume building and entrance exam preparation, and support with the application and interview process. "
// 	, "Lat": 46.727471200000004, "Long": -117.02390190000001,
// 	"HasShadow": false, "HasCost": false, "HasTransport": false, "Under18": false,
// 	"CareerEmp": ["Medicine", "Dentistry", "Pharmacy", "Public Health", "Generic Health Sciences",
// 	"Allied Health", "STEM"],
// 	"GradeLevels": [9, 10, 11, 12]},
