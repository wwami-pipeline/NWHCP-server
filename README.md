# Waypoint Database


## About
Database runs on MongoDB allows for getting by ID/name, inserting, updating, and deleting of schools and organizations.

## Setup
Clone down the repository. 

To run locally:
  - Install go 1.11+ (because it's using go modules, go.mod is like package.json in node.js)
  - Run ```go mod vendor``` to install dependencies. You should run it when adding new dependencies.
  - Start mongoDB on port 27017
  - Run ```go run main.go```

Or:
  - Refer to NWHCP-docker/readme.md to use docker-compose to start a mongoDB

## API:

### /POST /api/v1/search

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

### GET /api/v1/org/{orgID}

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


### On port 90, which is only accessible from inisde docker network:

####  POST /pipeline-db/truncate

Clear the database. Will be called by NWHCP-data_cleaning microservice.

####   /pipeline-db/poporgs

Post an array of orgs to insert into database. Will be called by NWHCP-data_cleaning microservice.
