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
  namechanges-> {"Recipe":"Model3D", "recipe":"model3d", "FastAPI":"Fiber", "RECIPE":"MODEL3D, "https://christophergs.com/":"asdf"}
  add something about bootstrapping? 

git branch <branch>
git branch -a

   git add . &&
 git commit -a &&
git push -u origin FNamechanges


git add . &&
 git commit -a &&
git push -u origin <branch>

after namechanges to text and files..
add semicolons to console(err)??
config.js change apiBasePath
comment out catch in index.js to stare at jwt issue in browser

https://stackoverflow.com/questions/70585472/websocketclient-js16-websocket-connection-to-ws-localhost3000-ws-failed-r
commented out localStorageTokenInterceptor and related in client.js fix later..
don't really want jwt requirements for signup and login..?

https://github.com/request/request/issues/579
but with post


9/21/2022
reformatted apiClient post info (email, pass) to conform with axios params (not headers btw)

https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k#bodyparser
use BodyParser instead of FormValue to read in params from backend
commment out more dashboard headers stuff
changed route in App.js / => Home to /home => Home
useState default loading true to false for Home

https://www.geeksforgeeks.org/file-uploading-in-react-js/
refactor using Hook instead of class/component for model publishing page

todo:
decide whether/how to get FormInput going

9/27/2022
decide.. popup form, or not.. ok popups bad..use separate page?

9/28/2022
no render.. -rewrite dom todo
get rid of non rendering issue due to infiinite loop by ommiting refreshing conditional..

bad jwt
why does localStorage("token") appear to not persist when switching pages?
or its it a problem simply with jwtdecode..?
just change "token_string" to "token" since that is the new json.parse output format

todos:
indicator/text to file being opened/queued for publishing
css formatting for formInputs 

9/29/2022
changed background/forminput/button with tailwind css
replaced this. and this.state with useState sets
changed formData.append to json-like formatting

refactor to have setModel3DForm as the standard, and to have FormData persist and not overwrite other Form object data onChange

todo: fix apiClient header issues in client.js

todo fix fetchUser routing/param issue on apiClient.get

curl --request GET --url http://localhost:3000/users/1

curl --data "email=iqiao2011@gmail.com&password=fakepassword" http://localhost:8000/auth/login

somewhat strange that changing the "pass" field to "password" in the login struct is required to get the API endpoint to work
about just as strange as having to change "lastname" to "surname" in the postgres table

9/30/2022
--data "email=iqiao2011@gmail.com" 
curl --request GET --url http://localhost:8000/users?email=iqiao2011%40gmail.com -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY0ODE3NzI1LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.Pa0-MFTVAKSu_mvgxo1IcVjobSf-nAD-OkmEMwZSmqM"

https://www.filestack.com/fileschool/react/react-file-upload/
https://www.codegrepper.com/code-examples/javascript/curl+authorization+bearer+in+axios+get
https://masteringjs.io/tutorials/axios/axios-multi-form-data
https://www.w3schools.com/react/react_usestate.asp
https://reqbin.com/req/c-sma2qrvp/curl-post-form-example
https://stackoverflow.com/questions/64480353/how-to-use-curl-commands-with-reactjs
https://www.codegrepper.com/code-examples/javascript/curl+authorization+bearer+in+axios+get


10/3/2022
whats the diff that makes the 401?

11243 8fae53ee-78fc-44ee-9eca-3624993955ac  401  -  GET      /users
 Sec-Ch-Ua-Mobile=?0&User-Agent=Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.83 Safari/537.36&Connection=keep-alive&Accept=application/json, text/plain, */*&Sec-Fetch-Dest=empty&Authorization=Bearer {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MDc1OTE1LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.VbPTA7qoI_bFj0eS6zQiyz2uzACaIFlPFj2YmK8Gcz8"}&Sec-Ch-Ua-Platform="Linux"&Sec-Fetch-Site=same-site&Host=localhost:8000&Accept-Language=en-US,en;q=0.9&Sec-Fetch-Mode=cors&Referer=http://localhost:3000/&Accept-Encoding=gzip, deflate, br&Sec-Ch-Ua="Chromium";v="93", " Not;A Brand";v="99"&Dnt=1&Origin=http://localhost:3000
email=iqiao2011%40gmail.com
getting users

11243 90ae53ee-78fc-44ee-9eca-3624993955ac  200  -  GET      /users
 Host=localhost:8000&User-Agent=curl/7.68.0&Accept=*/*&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY0ODE3NzI1LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.Pa0-MFTVAKSu_mvgxo1IcVjobSf-nAD-OkmEMwZSmqM
email=iqiao2011%40gmail.com

https://docs.gofiber.io/api/middleware/logger
token formatting issue, use resp.data.token to store token data as string instead of obj in localstorage

https://k3d.ivank.net/?p=download

rewrote req to /model3d fixed url, body/header format fixed state update to handle event.target info before it get thrown away

todo get K3D working or write your own.. or look for another library

9017 e02c68d5-332e-4784-a4da-371cdc0acf49  400  -  POST     /model3ds
Host=localhost:8000&Origin=http://localhost:3000&Accept-Language=en-US,en;q=0.9&Connection=keep-alive&Sec-Ch-Ua-Mobile=?0&Sec-Ch-Ua-Platform="Linux"&Sec-Fetch-Mode=cors&Sec-Fetch-Site=same-site&Content-Type=application/json&User-Agent=Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.83 Safari/537.36&Sec-Ch-Ua="Chromium";v="93", " Not;A Brand";v="99"&Accept=application/json, text/plain, */*&Referer=http://localhost:3000/&Accept-Encoding=gzip, deflate, br&Content-Length=190&Dnt=1&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MDkxMjkwLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.gkfn56WmxtyTuYXKpD5h88YjWyMwvgx0XpXBkrIW4kU&Sec-Fetch-Dest=empty

{"title":"asdf1","author":1,"description":"asdf2","price":"asdf3","blob_data":{"groups":{},"c_verts":[],"c_uvt":[],"c_norms":[],"i_verts":[],"i_uvt":[],"i_norms":[]},"file_name":"Blob1.obj"}

curl --request POST \
  --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99,
    "BlobData":{},
    "FileName":"dezznutz.obj"
  }' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MDkzMjYwLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.pFXmMW7qPdXYph9SKGG1ysgnO59My7bQjhHsK1cxfY8"

whats the diff?..

StormTrouuuper Helmet High Poly
11645 c0f932cd-3c57-4d05-904e-da447099756d  201  -  POST     /model3ds
Content-Length=186&Content-Type=application/json&User-Agent=curl/7.68.0&Accept=*/*&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MDkzMjYwLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.pFXmMW7qPdXYph9SKGG1ysgnO59My7bQjhHsK1cxfY8&Host=localhost:8000

{
    "Title":"StormTrouuuper Helmet High Poly",
    "Author":"PandaFarmer",
    "Description":"I like hats haha",
    "Price":-6.99,
    "BlobData":{},
    "FileName":"dezznutz.obj"
  }

12637 c84389f1-55a5-4ea1-8ed3-271df0afab00  400  -  POST     /model3ds
User-Agent=Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.83 Safari/537.36&Sec-Ch-Ua="Chromium";v="93", " Not;A Brand";v="99"&Sec-Fetch-Mode=cors&Host=localhost:8000&Content-Length=188&Sec-Ch-Ua-Mobile=?0&Sec-Ch-Ua-Platform="Linux"&Accept-Language=en-US,en;q=0.9&Connection=keep-alive&Accept=application/json, text/plain, */*&Sec-Fetch-Site=same-site&Accept-Encoding=gzip, deflate, br&Content-Type=application/json&Dnt=1&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MDkzMjYwLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.pFXmMW7qPdXYph9SKGG1ysgnO59My7bQjhHsK1cxfY8&Origin=http://localhost:3000&Sec-Fetch-Dest=empty&Referer=http://localhost:3000/

{"Title":"asdf1","Author":1,"Description":"asdf2","Price":"28.99","BlobData":{"groups":{},"c_verts":[],"c_uvt":[],"c_norms":[],"i_verts":[],"i_uvt":[],"i_norms":[]},"FileName":"Blob1.obj"}

just use Math.round(number * 100) / 100
