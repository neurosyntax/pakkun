# Pakkun![Go gopher](./images/gopherbelly50.jpg)
### Essentially a refactored github-scraper. Currently, does not have the GitHub API features, but will add these soon. That repo really needed refactoring...

Language: Go (Golang)

A tool for scraping repositories from GitHub and extracting source code function information using Golang and GitHub API v3. Assumes you're running Linux. If running on OSX, use the appropriate dependencies. This script relies on bash commands and so will not work on Windows.

#### Setup/ Dependencies

[Install MongoDB](https://golang.org/doc/install)

[Install MongoDB on Ubuntu](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

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

#### Test:
By default the script looks for functions containing numeric/boolean input parameters and outputs*.
You should only get back test3() and test7() since that's the only one with only numeric or boolean values.
```
go run main.go -dir ~/pakkun

output:
{
	"_id" : 911846790,
	"name" : "foobar.java",
	"path" : "/home/ubuntu/research/sandbox/go_workspace/pakkun/test/foobar.java",
	"funcs" : 
		[
			{
				"id" : NumberLong("2298940924"),
					"name" : "test7",
					"header" : "public static double test7(int i, int j, int k)",
					"intype" : [ "int", "int", "int" ],
					"outtype" : [ "double" ],
					"source" : "public static double test7(int i, int j, int k) {}"
			}
		]
}
{
	"_id" : 1170655353,
	"name" : "helloworld.java",
	"path" : "/home/ubuntu/research/sandbox/go_workspace/pakkun/test/helloworld.java",
	"funcs" : 
		[ 
			{
				"id" : NumberLong("3700423157"),
				"name" : "test3",
				"header" : "public static float test3(int i, double d, float f)",
				"intype" : [ "float", "int", "double" ],
				"outtype" : [ "float" ],
				"source" : "public static float test3(int i, double d, float f) {}"
			}
		]
}
```

*  *Java functions headers should only specify one return type. However, languages like Golang allow for specifying multiple return types. So allowing for multiple types will be useful for extending this to languages like Go, but it also requires assuming the programs are syntactically correct.


