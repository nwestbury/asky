# Install instructions

## Install Golang
Fetch GoLang 1.8.3 from https://golang.org/dl/

## Install Dependencies
go get github.com/gorilla/mux
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/facebook
go get github.com/asaskevich/govalidator
go get github.com/lib/pq

## Install PostgreSQL

sudo add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main"
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get install postgresql-9.6

## Import Schema
psql -d flash -f sql/database-flash.sql

## (!) Important create a bin/install.sh script
For example bin/install.sh:

export ENV_FB_CLIENT_ID="123"
export ENV_FB_CLIENT_SECRET="123"

## To run the server run the run.sh script
./run.sh
