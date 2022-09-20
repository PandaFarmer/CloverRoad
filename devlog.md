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
    Surname varchar(255),
    FirstName varchar(255),
    City varchar(255)
);

8/18/2022
https://www.databasestar.com/sql-show-databases/
https://stackoverflow.com/questions/7695962/password-authentication-failed-for-user-postgres
https://nixcp.com/postgresql-restart-command/

https://gorm.io/docs/connecting_to_the_database.html
ALTER USER user_name WITH PASSWORD 'new_password';

 git add . &&
 git commit -a &&
 git push -u origin master

 curl --request POST \
  --url http://localhost:3000/users \
  --header 'Content-Type: application/json' \
  --data '{
    "email": "iqiao2011@gmail.com",
    "user_name": "PandaFarmer",
    "first_name": "Isaac",
    "surname": "Qiao",
    "password": "fakepassword",
    "date_joined": "2012-04-23T18:25:43.511Z",
    "is_super_user": true
  }'

   curl --request PUT \
  --url http://localhost:3000/users/1 \
  --header 'Content-Type: application/json' \
  --data '{
    "email": "iqiao2011@gmail.com",
    "user_name": "PandaFarmer",
    "first_name": "Isaac",
    "surname": "Qiao",
    "password": "fakepassword",
    "date_joined": "2012-04-23T18:25:43.511Z",
    "is_super_user": true
  }'

  curl --request GET --url http://localhost:3000/users/1


curl --request POST \
  --url http://localhost:3000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99
  }'

curl --request PUT \
  --url http://localhost:3000/model3ds/1 \
  --header 'Content-Type: application/json' \
  --data '{
 "Title":"StormTrooper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"Ligma",
    "Price":-6.99
}'

curl --request DELETE \
  --url http://localhost:3000/users/2


curl --request GET --url http://localhost:3000/model3ds/1

List of todos:
api versioning?
logging w/ flags? - https://docs.gofiber.io/api/middleware/logger
async
iam w/ jwts
add blob/[]byte to Model3D struct
react/threejs frontend
deploy on aws

rewrite in asp or actix haha, asp more likely

https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
https://github.com/gofiber/jwt

https://www.jajaldoang.com/post/how-to-update-golang/


curl --data "email=iqiao2011@gmail.com&pass=fakepassword" http://localhost:8000/auth/login

curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxMTIyNDI1LCJuYW1lIjoiSm9obiBEb2UifQ.E3z-ecQYJ0UZCyy0uT8rS6XH7SE6M_XgkLhiQ9UsK1g"

curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxNDQwOTUyLCJuYW1lIjoiSm9obiBEb2UifQ.sU0bcnmX7Kv-1TVrhX627TRDkIC5JCnH2HSBX0kwe2Q"

curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxNDUwNTczLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.3V2aYPIsgsXaZpKniLGyD41uwCU6liMcEEWxd0Z9rNM"

  curl --request GET --url http://localhost:8000/users/1 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxNjMwNDUzLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.qg2b_x2gZRQ8fPQeo-qmMp_gukfFtBs5CJdm9DvIGic"

8/22/2022
jwt issues..
no username in token?
ok issue with User schema..response needs to have a UserName and FirstName, not user_name and first_name..

psql -U postgres -h localhost -W
db pass: 3t4t5LQS

\l
\c <database-name>
\dt
\dS <table-name>

\x
\df+

DROP DATABASE <database-name>;
DROP TABLE <table-name>;

PUT does not behave properly
make sure curl json input is sanitized? LastName last_name json value/sql column mismatch.. issue with struct?
response values for add_user bad, issue..
just fix/exclude routing middleware jwtware addon for some endpoint-> done


mkdir repo &&
cd repo &&
git init &&
git remote add -f origin $URL &&
git config core.sparseCheckout true
echo '$DIRNAME' >> .git/info/sparse-checkout # Do this for every directory you want.
git fetch
git checkout $BRANCH # Typically master


https://github.com/ChristopherGS/ultimate-fastapi-tutorial
part-12-react-frontend/frontend/

https://linuxconfig.org/install-npm-on-linux
https://phoenixnap.com/kb/update-node-js-version


https://phoenixnap.com/kb/update-node-js-version

8/23/2022
write tests for backend
make frontend
add blob handling

npm i react
https://create-react-app.dev

better->

mkdir repo &&
cd repo &&
git init &&
git remote add -f origin https://github.com/ChristopherGS/ultimate-fastapi-tutorial &&
git config core.sparseCheckout true &&
echo 'part-12-react-frontend/frontend/' >> .git/info/sparse-checkout &&
git fetch &&
git checkout main

8/24/2022
the frontend:
App-Name and Logo ... ->  search bar ...JOIN / LOG IN -> cart?
Categories... A|B|C|D ... SORT BY.. or filter by range  LICENSE? Start/Sell

DropDowns for filter/sort.. DropDowns for asdf
filter for free/nsfw
filter by file type
filter low/high poly
drop downs for categories?
sort by price asc/desc (high/low) newest

changed dev.env to 8000

-notes:
include frontend
exclude *.go, *.md, *.lock, *.json
namechanges: Recipe, recipe, RECIPE
including folder and file names..
changed backend path /login to auth/login to match frontend
changed parameters 'username' and 'password' to 'email' and 'pass' in client.js login fn to match backend
later change some text in markup doctre

8/26/2022
notion- cloud storage https://gocloud.dev/howto/blob/

curl --data "email=iqiao2011@gmail.com&pass=fakepassword" http://localhost:8000/auth/login

curl -F 'file=@/home/iqiao/Documents/BlenderBlobs/Blob1.obj' \
  --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99
    "FileName":"Blob1.obj"
  }' \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxNzkxOTkzLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.dgWkUTshCnfjF6mqJN54wrwUcEs3tRdDol1EBMp_r_A"

/home/iqiao/Documents/BlenderBlobs/Blob1.obj

curl --request POST \
  --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99
  }'

curl --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99,
    "":
    "FileName":"Blob1.obj"
  }' \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjYxNzkxOTkzLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.dgWkUTshCnfjF6mqJN54wrwUcEs3tRdDol1EBMp_r_A"

  9/1/2022
  home/index.jsx
  handle commented
  fetchRecipes
  RecipeTable/index.jsx
  Recipe/index.jsx
  
  9/5/2022
  find where to disable jwt requirement for login page

  9/16/2022
  remove dashboardheader in jsx ref, jwt requirement stems from there
  
  9/20/2022
  just git clone and comment out jwt in dashboard header, jwt/moment- fix later
  namechanges-> {"Recipe":"Model3D", "recipe":"model3d"}

   git add . &&
 git commit -a &&
git push -u FNamechanges~


$ git remote add new-remote-repo https://bitbucket.com/user/repo.git
# Add remote repo to local repo config
$ git push <new-remote-repo> crazy-experiment~
# pushes the crazy-experiment branch to new-remote-repo