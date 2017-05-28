# Install instructions

## Install Golang
Fetch GoLang 1.8.3 from https://golang.org/dl/

## Install PostgreSQL

sudo add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main"
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get install postgresql-9.6

## (!) Important create a bin/install.sh script
For example bin/install.sh:

export ENV_FB_CLIENT_ID="123"
export ENV_FB_CLIENT_SECRET="123"

## To run the server run the run.sh script
./run.sh