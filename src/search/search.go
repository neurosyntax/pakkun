/*
    search.go

    A module for searching directories
    
    Author: Justin Chen
    2.12.2017

    Boston University 
    Computer Science

    Dependencies:        exuberant ctags, and mongodb driver for go (http://labix.org/mgo)
    Operating systems:   GNU Linux, OS X
    Supported languages: C, C++, C#, Erlang, Lisp, Lua, Java, Javascript, and Python
*/

package search

import (
    "path/filepath"
	"strings"
	"os"
	"parse"
    "utils"
    "gopkg.in/mgo.v2"
)

func SearchAndSaveFunc(session *mgo.Session, searchDir string, extension string, funcTypes map[string]bool) {
    // var total int
    // Walk directory and parse java files as they're found
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
    	if strings.HasSuffix(path, extension) {
            file, ok := parse.ParseFile(path, funcTypes)

            if ok {
                utils.SaveMgoDoc("github_repos", "source", file)
            }
        }
        return nil
    })
}