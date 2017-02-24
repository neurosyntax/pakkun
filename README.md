# Pakkun
### Essentially a refactored github-scraper. Currently, does not have the GitHub API features, but will add these soon. That repo really need refactoring...

# Pakkun![Go gopher](./images/gopherbelly50.jpg)

Language: Go (Golang)

A tool for scraping repositories from GitHub and extracting source code function information using Golang and GitHub API v3. Assumes you're running Linux. If running on OSX, use the appropriate dependencies. This script relies on bash commands and so will not work on Windows.

#### Setup/ Dependencies

[Install MongoDB](https://golang.org/doc/install)

[Install MongoDB](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

Install Exuberant Ctags:
```sh
Ubuntu:
sudo apt install exuberant-ctags
```
Install MongoDB driver for Go:
```sh
go get gopkg.in/mgo.v2
```
Refer to [mgo](https://github.com/go-mgo/mgo) for further and more up-to-date instructions.

GOPATH setup
```sh
Linux:
sudo nano ~/.bashrc
export GOPATH=$HOME/<path to this repo>
source ~/.bashrc

Unix:
sudo nano ~/.bash_profile
export GOPATH=$HOME/<path to this repo>
touch ~/.bash_profile
```

If you get a message about GOPATH not able to find e.g. `parse/parseFuncHeader.go`, do the following:
```sh
cd pakkun/src/parse
cd pakkun/src/search
cd pakkun/src/utils
go build
go install
```

#### Basic usage:
```sh
go run main.go -dir <absolute path>
```


