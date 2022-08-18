8/15/2022
attempt at creating a 3Dmodeling marketplace webapp
stack of golang-fiber, sqlite->redis, react+three js?
things to keep in mind:
Logging, db/models/schema??- how does orm handle this?, database migration handler?, virtual envs/manager/containers..
iam, cors, api versioning/prod structure, async, dependency injection system, 

/*
https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.36.0

GOPROXY=https://proxy.golang.org GO111MODULE=on go get example.com/my/module@v1.0.0

GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/gofiber/fiber/v2@v2.36.0

https://betterprogramming.pub/simple-api-in-go-with-fiber-1449198e0237

go mod init github.com/YOUR_USERNAME/go-fiber-api



go get github.com/spf13/viper &&
go get github.com/gofiber/fiber/v2 &&
go get gorm.io/gorm &&
go get gorm.io/driver/postgres

touch Makefile cmd/main.go pkg/books/add_book.go pkg/books/controller.go pkg/books/delete_book.go pkg/books/get_book.go pkg/books/get_books.go pkg/books/update_book.go pkg/common/db/db.go pkg/common/config/envs/.env pkg/common/models/book.go pkg/common/config/config.go
*/

mkdir -p cmd pkg/model3ds pkg/common/db pkg/common/config/envs pkg/common/models

touch Makefile cmd/main.go pkg/model3ds/add_model3d.go pkg/model3ds/controller.go pkg/model3ds/delete_model3d.go pkg/model3ds/get_model3d.go pkg/model3ds/get_model3ds.go pkg/model3ds/update_model3d.go pkg/common/db/db.go pkg/common/config/envs/.env pkg/common/models/model3d.go pkg/common/config/config.go

8/17/2022
git init && git add . && git commit -m 'Initial Commit' && git push -u origin master 
git push -u origin master

finished cleaning up namechanges and iqiao->PandaFarmer
https://sqlserverguides.com/postgresql-installation-on-linux/

->not needed due to automigrate?

CREATE TABLE Model3Ds (
    Id int,
    Title varchar(255),
    Author varchar(255),
    Description varchar(255),
    Price float8
);

CREATE TABLE Users (
    PersonID int,
    LastName varchar(255),
    FirstName varchar(255),
    City varchar(255)
);

8/18/2022
https://gorm.io/docs/connecting_to_the_database.html
ALTER USER user_name WITH PASSWORD 'new_password';

