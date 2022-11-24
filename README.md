# Data Development Kit

================================

Go Libaries used in this project 

================================

- "github.com/99designs/gqlgen/" - Used for generating models and resolver from schema
- "gorm.io/gorm" - Used for connecting to DB Setup
- "github.com/go-chi/chi" - Used for http routing and adding middleware


# IMPORTING GQLGEN PACKAGE

DataDevKit > go mod init github.com/kirankkirankumar/gqlgen-ddk  
==================
- Expected Output
==================
go: creating new go.mod: module github.com/kirankkirankumar/gqlgen-ddk

# CREATING "tools.go" file in parallel to graph folder
======================================================

import (
	_ "github.com/99designs/gqlgen"
)

# IMPORTING postgres PACKAGE ->  go get gorm.io/driver/postgres
==================
- Expected Output
==================

go: added github.com/jackc/chunkreader/v2 v2.0.1

go: added github.com/jackc/pgconn v1.13.0

go: added github.com/jackc/pgio v1.0.0

go: added github.com/jackc/pgpassfile v1.0.0

go: added github.com/jackc/pgproto3/v2 v2.3.1

go: added github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b

go: added github.com/jackc/pgtype v1.12.0

go: added github.com/jackc/pgx/v4 v4.17.2

go: upgraded golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 => v0.0.0-20220722155217-630584e8d5aa

go: added gorm.io/driver/postgres v1.4.5

==================================

# Setting environmental Variables 

==================================
- create .env file and specify the Required DB details 
DataDevKit % source .env

# IMPORTING GORM PACKAGE ->  go get gorm.io/gorm
 DataDevKit % go get gorm.io/gorm
==================
- Expected Output
==================

go: downloading gorm.io/gorm v1.24.1

go: added github.com/jinzhu/inflection v1.0.0

go: added github.com/jinzhu/now v1.1.4

go: added gorm.io/gorm v1.24.1


# Testing gqlgen generate

=============================

 - go run github.com/99designs/gqlgen generate


# To start the gqlgen server ->  go run server.go

=================================================

 - go run server.go


 # Clone repository using "git clone github.com/kirankkirankumar/gqlgen-ddk" command
===============
Clone the Repo
===============

- Update the environment variables in .env file
- Run command "source .env" to load the variables
- Run command "go run server.go" to the run server




RULES :

- If we add any add/modify any mapping in graphql file we need to give definition in Service.go

- Run command "source .env" to load the variables

- Run command "go run server.go" to the run server
