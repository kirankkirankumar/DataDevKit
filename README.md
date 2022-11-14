# DataDevKit


# IMPORTING GQLGEN PACKAGE
DataDevKit > go mod init github.com/kirankkirankumar/gqlgen-ddk  
go: creating new go.mod: module github.com/kirankkirankumar/gqlgen-ddk

# CREATING "tools.go"




# IMPORTING postgres PACKAGE ->  go get gorm.io/driver/postgres
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


# IMPORTING GORM PACKAGE ->  go get gorm.io/gorm
kirankkirankumar@US-NQXF0JXKQ0 DataDevKit % go get gorm.io/gorm
go: downloading gorm.io/gorm v1.24.1
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.4
go: added gorm.io/gorm v1.24.1

# Testing gqlgen generate
 - go run github.com/99designs/gqlgen generate

# To start the gqlgen server ->  go run server.go
 - go run server.go