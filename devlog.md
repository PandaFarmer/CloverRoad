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

  curl --request GET --url http://localhost:8000/users


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


curl --data "email=iqiao2011@gmail.com&password=fakepassword" http://localhost:8000/auth/login

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

\q

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
curl --request GET --url http://localhost:8000/users?email=iqiao2011%40gmail.com -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3NDk3OTg0LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.ovFy2c_-grXVP2sJ6mPJ0pKvsL0-Ag-ePS6gqbigYnk"

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
fixed formatting for model3ds post request, matching field names in model3dData with backend request struct, formatting/type correctness, localStorage changes

10/4/2022
curl --request GET --url http://localhost:8000/model3ds -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY1MTU3MzEyLCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.pVcFIHZ2V8OSKHalxsw9ePceqb9qetIPu6aqwTIryTs"

fixed issue with localstorage and username

10/5/2022
Login indicator in dashboard needs to work, state isLoggedIn must be consistent
cleanup any lazy state to true sets, disallow jwt requirements for login/signup/default homepage
https://reactjs.org/docs/conditional-rendering.html
https://reactjs.org/docs/rendering-elements.html
Fixed login indicator for dashboard.. ++failure to login text

10/7/2022
https://stackoverflow.com/questions/37134433/convert-input-file-to-byte-array
lmao just convert to byte array, imagine following request download struct field typing defined in golang backend
fixed/completed upload/publish..?

https://gorm.io/docs/query.html

my-model3ds.. redo model3dForm to accept and convert byte arrays?-else there will be a undefined error..

npm i Buffer

WHY DOES BODYPARSER NOT WORK AS INTENDED?

10/11/2022
read the fiber bodyparser source code,
https://github.com/gofiber/fiber/blob/master/ctx.go
see the strings.HasPrefix(ctype, MIMEApplicationForm)
has
data := make(map[string][]string)
which i suppose allows fiber to not drop data and make an actual copy of the byte array
https://stackoverflow.com/questions/43111772/golang-byte-array-communication-through-channel-loses-data
https://stackoverflow.com/questions/55550834/how-to-set-mime-type-for-post-multipart-form-data-in-axios
so, if you want to pass file data, use MIMEApplicationForm instead of json, in both front and backend
createModel3D(client.js) AddModel3dRequestBody(add_model3d.go)
also the issue seems to have nothing to do with the app.config ReadBufferSize param

to reset db, go into the psql and DELETE FROM <tablename>;
https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-delete/

https://stackoverflow.com/questions/50696857/append-array-of-values-to-the-current-array-in-a-state-in-react-js



10/13/2022
https://r105.threejsfundamentals.org/threejs/lessons/threejs-load-obj.html
https://github.com/willbamford/react-with-threejs-example/blob/master/src/VisWithHooks.jsx
https://mockstacks.com/Write-byte-array-to-file-in-Javascript

10/14/2022
https://hacks.mozilla.org/2012/02/saving-images-and-files-in-localstorage/

https://developer.mozilla.org/en-US/docs/Web/API/FileReader
https://developer.mozilla.org/en-US/docs/Web/API/Blob


10/16/2022
https://stackoverflow.com/questions/35038884/download-file-from-bytes-in-javascript

https://stackoverflow.com/questions/680929/how-to-extract-extension-from-filename-string-in-javascript

10/17/2022
https://gist.github.com/greggman/d34fbccd00b283bfb54aec1ebf3d973a

https://stackoverflow.com/questions/18283772/how-to-create-a-file-object-from-binary-data-in-javascript

don't know why it wont render
possible reasons:
-scuffed resource url
-tailwind conflicts with threejs rendering properly
-useEffect not executing? due to external or internal?
-some threejs parameter required for render is missing
-mesh material is undefined, issue with render?
-rootNode of DOMTree aint being referenced properly kappa
-

10/21/2022
todo:
some tailwind changes to center model
display info as well
fix lack of multiple model3d components being displayed-elements get replaced per render?
Add lighting or background or camera changes to improve presentation..
Camera rotation on mouse drag

file formats more than glb/gltf

10/22/2022
local webserver notion? client side ish?-but without fiber or drei :(
https://stackoverflow.com/questions/64684300/cant-load-an-obj-file-with-three-js-and-objloader2

fixed dynamic population for Model3DTable
https://thewebdev.info/2021/03/13/how-to-push-or-append-an-element-to-a-state-array-with-react-hooks/

changed to using gltf, now shows models in pixelated profile

to consider:
model3d packages->referencing multiple entries?
pointerDown vs. mouseDown -> rotate model or update orthographic camera?

10/23/2022
screw three js
just have users upload png/gif, +blob that is an array of files

https://stackoverflow.com/questions/35038884/download-obj-from-bytes-in-javascript
https://mockstacks.com/How-to-Convert-Blob-to-File-Using-Javascript


10/27/2022
?
clearing localstorage/ application/browser cache required for backend to return body/resBody?????????????????????????????????????????????????????????

for issue with Access-Control-Allow-Origin
https://docs.gofiber.io/api/constants
https://github.com/gofiber/fiber/blob/master/ctx.go

10/28/2022
some interesting examples?:
https://github.com/gofiber/recipes

Ok api works with cors and csrf?
cors header settings needed:AllowOrigins, AllowMethods, AllowHeaders,AllowCredentials
single defer after (only for registering endpoint requiring jwt auth)
client side header addition??-not needed..
3rd param setting needed for cors, not just inclusive of 2

https://gorm.io/docs/method_chaining.html
https://www.golinuxcloud.com/golang-defer-keyword/
https://docs.gofiber.io/v/2.x/api/middleware/cors
https://stackoverflow.com/questions/28252117/golang-nested-class-inside-function

10/30/2022
https://pkg.go.dev/time#Time.Format
https://www.w3schools.com/jsref/jsref_tostring_date.asp
https://javascript.info/date

curl --request GET \
  --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "keywords" : "asdf",
    "max_results" : 10,
    "lastUpdated" : ,
    "minPrice" : ,
    "maxPrice" : ,
  }'

10/31/2022
jwt randomly decides to stop working on endpoints "/Publish" (which was working b4) and all others in deferred registration
403 on a previously working endpoint..

11/1/2022
How to solve to 403:
fiddle with config params for jwt initialization..
more logging?
why does GET users work but POST model3d or GET model3ds does not?

curl --request POST \
  --url http://localhost:8000/model3ds \
  --header 'Content-Type: application/json' \
  --data '{
    "title":"sugma",
    "author":"sugma",
    "description":"sugma",
    "price":"42",
    "serialized_preview_file":{},
    "preview_file_name_and_extension":"sugma.png",
    "serialized_file_3d":{},
    "file_name_and_extension":"sugma.obj",
  }' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3NDk3OTg0LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.ovFy2c_-grXVP2sJ6mPJ0pKvsL0-Ag-ePS6gqbigYnk"

WHY:
pid: 11896bd7002ed-2a79-4bac-a92f-4eb6464a14cc  200  -  GET      /users
reqHeaders:Accept=application/json, text/plain, */*&Sec-Ch-Ua-Platform="Linux"&Connection=keep-alive&Sec-Fetch-Mode=cors&Referer=http://localhost:3000/&Host=localhost:8000&User-Agent=Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.83 Safari/537.36&Sec-Ch-Ua="Chromium";v="93", " Not;A Brand";v="99"&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3NDk3OTg0LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.ovFy2c_-grXVP2sJ6mPJ0pKvsL0-Ag-ePS6gqbigYnk&Sec-Ch-Ua-Mobile=?0&Sec-Fetch-Dest=empty&Accept-Encoding=gzip, deflate, br&Accept-Language=en-US,en;q=0.9&Dnt=1&Origin=http://localhost:3000&Sec-Fetch-Site=same-site
queryParams:email=iqiao2011%40gmail.com
body:
resBody:[{"id":1,"email":"iqiao2011@gmail.com","user_name":"PandaFarmer","full_name":"","password":"$2a$04$CFrXGrx1/PS8MdaHIUPhY.njABElWe67.1kmctImErJu917SFMbPi","date_joined":"2022-08-22T09:49:41.885828-07:00","is_super_user":true}]
resHeaders:

??????????????//
pid: 1521186bf3c93-12a4-4abe-b851-3d0e0f293164  403  -  POST     /model3ds
reqHeaders:Content-Length=208340&Sec-Ch-Ua="Chromium";v="93", " Not;A Brand";v="99"&Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3NDk3OTg0LCJuYW1lIjoiUGFuZGFGYXJtZXIifQ.ovFy2c_-grXVP2sJ6mPJ0pKvsL0-Ag-ePS6gqbigYnk&Sec-Fetch-Mode=cors&Accept-Encoding=gzip, deflate, br&Connection=keep-alive&Dnt=1&Sec-Ch-Ua-Mobile=?0&Accept=application/json, text/plain, */*&Accept-Language=en-US,en;q=0.9&Host=localhost:8000&Content-Type=multipart/form-data; boundary=----WebKitFormBoundary9Odw9dbqWeT1cirq&Sec-Ch-Ua-Platform="Linux"&Origin=http://localhost:3000&Sec-Fetch-Dest=empty&Referer=http://localhost:3000/&User-Agent=Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.83 Safari/537.36&Sec-Fetch-Site=same-site
queryParams:


Todos:

for reasonable MVP:
search functionality in my-model3ds and/or homepage, same background
tags/categories
onHover display info?
onClick smaller link to buy/inspect model
margins/padding per tile,
publish/save -> show preview image
multiple files/filetypes per model3d entry
implement search and download
switch user?
Redirect login page to publish/my-model3ds/home if already logged in
fix my-model-3ds page not rendering if user has no publishings
Signin Failure Behavior, captcha?
NotLoggedIn Redirects
Login Session timed expiration
More Error Handling?
Form Input validation/checking for publish page must be finalized/catch edge cases?
sales prices/percent offs+duration
changes to dashboard-colours, size?
add paypal?-some transaction recording...
aws/other cloud hosting?
caching?
testing kekw-but more logging, at least..?
extension support for up to half of these:
obj, fbx, max, c3d, ma/mb, blend, unitypackage, upk/uasset, dae, 3ds, skp, Lxo, lwo/lws, stl

11/9/2022
https://github.com/hellokvn/go-gin-api-medium
https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0

circular dependencies- db.automigrate uses models type structs, but current impl of the user model utilizes db in a method
solution-> use a mediator, just like the kevj guy does, with h handler (which holds a DB field for reference) as mediator, however there is an issue with passing in c *gin.Ctx if not passing method as a callback parameter.. that i haven't figured out..
also will likely need to have some type conversion byte[] wtv before passing in params..? or at least milk a user object for its password field

also will have to replace some user method references (called from login/signup fn endpoint defs) with auth fns (defined as handler methods?)

also is it anti-idiomatic or just bad to call a endpoint fn without using http or gin to interface..?

also need to check for model/struct assignment on crud endpoints
todo: use docker, comprehensive testing (backend), 

11/10/2022
implemented, fixing compilation/syntax issues.. ++namechanges
write tests for db, users/model3ds endpoints, auth endpoints, 

https://stackoverflow.com/questions/36357791/gin-router-path-segment-conflicts-with-existing-wildcard

