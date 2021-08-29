# NWCHP test server

live at [https://test.nwhcp.tk](https://test.nwhcp.tk)

![workflow](https://github.com/wwami-pipeline/NWHCP-server/actions/workflows/mail.yml/badge.svg)

## About

Database runs on MongoDB allows for getting by ID/name, inserting, updating, and deleting of schools and organizations.

## Setup

Clone the repository.

To run locally:

- Install Go
- Install Docker
- Run ```sh build-local.sh```
- A MongoDB database is required, with test data loaded from ```org.json``` *Free at MongoDB Atlas*
- A MySQL database is required, with schema in ```userdb_schema.sql``` *Free at Heroku*
- A Redis database is required *Free at Heroku*
- A ```docker.env``` file is required, based on ```docker.env.sample``` with updated values

To deploy to the server:

- Code auto deploys on changes to master branch

## API:

The test server is hosted at: [https://test.nwhcp.tk](https://test.nwhcp.tk)

-------------

### *Organization Functions*

### POST /api/v3/search

Sample request body:

```
{
    "searchContent": "seattle",
    "CareerEmp": ["Nursing"],
    "HasCost": false,
    "Under18":false, 
    "HasTransport":false, 
    "HasShadow":false, 
    "GradeLevels":[]
}
```

You can inspect the website with Chrome developer tool to get a sample response, basically it's and array of orgs.

### GET /api/v3/orgs

Returns an array of all the organizations in the database

Sample response:

```
[
  {
    "OrgId": 8,
    "OrgTitle": "Pre-Health Advising",
    "OrgWebsite": "https://www.washington.edu/uaa/advising/at-the-uw/pre-health-advising/overview/",
    "StreetAddress": "University of Washington",
    "City": "Seattle",
    "State": "Washington",
    "ZipCode": "98195-2805",
    "Phone": "(206) 685-6582",
    "Email": "danpoux@uw.edu",
    "ActivityDesc": "Empowering UW students as they identify a health career path and prepare to apply to professional schools.\r\n\r\nThrough working with UW Pre-Health Advisers, students will be able to: \r\n\r\n* Develop academic and experiential pathways to their desired health profession\r\n* Participate in meaningful clinical experiences\r\n* Find campus and community resources for academic support, leadership development\r\nand research\r\n* Connect to the larger UW pre-health community\r\n\r\n",
    "Lat": 47.6499333,
    "Long": -122.3084015,
    "HasShadow": false,
    "HasCost": false,
    "HasTransport": false,
    "Under18": false,
    "CareerEmp": [
      "Medicine",
      "Nursing",
      "Dentistry",
      "Pharmacy",
      "Public Health"
    ],
    "GradeLevels": []
  },
  {
    "OrgId": 19,
    "OrgTitle": "Future Nurses Club",
    "OrgWebsite": "",
    "StreetAddress": "1959 NE Pacific Street",
    "City": "Seattle",
    "State": "Washington",
    "ZipCode": "98195",
    "Phone": "(206) 543-8736",
    "Email": "sukras@uw.edu",
    "ActivityDesc": "We will serve as a bridge from current BSN students to undergraduate students interested in learning more about the field of nursing, getting involved in service learning projects, and connecting with other students with similar interests.\r\n\r\nWe will support and advocate for future nurses regardless of race, ethnicity, citizenship status, religion, educational and socioeconomic status, disability, age, sex, gender identity, and sexual orientation.",
    "Lat": 47.6498722,
    "Long": -122.3082296,
    "HasShadow": false,
    "HasCost": false,
    "HasTransport": false,
    "Under18": false,
    "CareerEmp": [
      "Nursing"
    ],
    "GradeLevels": []
  }
]
```

### GET /api/v3/orgs/{orgID}

Get by org ID.

Sample response:

```
{
    "OrgId": 3,
    "OrgTitle": "Community Health Professions Academy",
    "OrgWebsite": "https://dental.washington.edu/oepd/our-programs/community-health-professions-academy/",
    "StreetAddress": "1959 NE PACIFIC ST",
    "City": "Seattle",
    "State": "Washington",
    "ZipCode": "98195-0001",
    "Phone": "(206) 221-1816",
    "Email": "UWOEPD@uw.edu",
    "ActivityDesc": "Community Health Professions Academy (CHPA) is a program aimed at students between the ages of 13 and 18, in which dental students, faculty, and pre-health undergraduate students, and AmeriCorps members volunteer their time to provide career guidance and enrichment for underrepresented high school students in Washington State. Over the course of seven sessions, the program focused on topics such as 'Leadership in Healthcare', 'Cardiac and Oral Health', 'Interprofessional care', and 'Traditional Medicine' alongside facilitating hands-on activities that broadened their perspective of the medical and dental field. The scholars, with aid from their health professional student mentors, also presented a project in front of family and friends for their final session.",
    "Lat": 47.6498722,
    "Long": -122.3082296,
    "HasShadow": false,
    "HasCost": false,
    "HasTransport": false,
    "Under18": true,
    "CareerEmp": [
        "Medicine",
        "Nursing",
        "Dentistry",
        "Public Health",
        "Generic Health Sciences"
    ],
    "GradeLevels": [
        0,
        1
    ]
}
```

-------------

### *User Functions*

### POST /api/v1/users

Creates a new user and inserts into the database Sample Request Body:

```
{
    “Email”: “anna@test.com”,
    “Password”: “password”,
    “PasswordConf”: “password”,
    “FirstName”: “Anna”,
    “LastName”: “Zhou
}
```

Returns Status `201`

### POST /api/v1/sessions

Logs the user into the website Sample Request Body:

```
{
    "Email": "anna2@test.com",
    "Password": "password"
}
```

Returns status code `201` and json of the user information:
Also in the returned response headers, there will be an authorization token used to access the user information Returns
session ID in Authorization header as Bearer: ________.

```
{
    "id": 1,
    "firstName": "Anna",
    "lastName": "Zhou",
    "joinDate": "05-19-2021"
}
```

### DELETE /api/v1/sessions/mine

Logs the user out of the website Returns status code `200`

### POST /api/v2/orgs/{id}

Adds organizations/programs to the user's account User must be authorized!
Returns status code `201`

### DELETE /api/v2/orgs/{id}

Deletes organizations/programs from the user's account User must be authorized Returns status code `200`

### GET /api/v2/getuser/

Returns all the organizations in the user's account Returns an array of organizations with basic information. User must
be authorized Sample response body:

```
[
    {
        "UserID": 1,
        "OrgID": 41,
        "OrgTitle": "Shoreline Community  College Post Baccalaureate Studies for the Health Professions",
        "OrgWebsite": "www.shoreline.edu",
        "OrgCity": "Shoreline",
        "OrgState": "Washington",
        "OrgZipcode": "98133",
        "OrgPhone": "(206) 564-6984"
    },
    {
        "UserID": 1,
        "OrgID": 44,
        "OrgTitle": "Tacoma Community College",
        "OrgWebsite": "Www.tacomacc.edu",
        "OrgCity": "Tacoma",
        "OrgState": "Washington",
        "OrgZipcode": "98466",
        "OrgPhone": "(253) 566-5131"
    }
]
```

----------

### *Data Management Functions*

### On port 90, which is only accessible from inisde docker network:

#### POST /pipeline-db/truncate

Clear the database. Will be called by NWHCP-data_cleaning microservice.

#### /pipeline-db/poporgs

Post an array of orgs to insert into database. Will be called by NWHCP-data_cleaning microservice.
