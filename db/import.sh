#! /bin/bash

echo "importing"

mongoimport --db mongodb --collection organization --file ./json/orgs.json --jsonArray
