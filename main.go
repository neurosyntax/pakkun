/*
    main.go

    Given a programming language type and parameter types, Black Hayate searches a directory and retrieves matching functions.

    Black Hayate is the name of Riza's dog from Fullmetal Alchemist
    
    Author: Justin Chen
    2.12.2017

    Boston University 
    Computer Science

    Dependencies:        exuberant ctags, and mongodb driver for go (http://labix.org/mgo)
    Operating systems:   GNU Linux, OS X
    Supported languages: C, C++, C#, Erlang, Lisp, Lua, Java, Javascript, and Python
*/

package main

import (
    "os"
	"flag"
    "search"
    "utils"
    "hash/fnv"
    "runtime"
)

type Test struct {
    Id   uint32
    Name string
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // Parse args
    flag.String("dir", "", "Directory to search")
	flag.Parse()

    // Store args
    options := make(map[string]string)

    key := ""
    for i, arg := range os.Args {
        if i > 0 {
            if i % 2 != 0 {
                key = arg[1:]
            } else {
                options[key] = arg
            } 
        }
    }

    // Search directory for functions of desired types
    searchDir := options["dir"]
    extension := ".java"

    // User-defined map to detect valid types. Valid if the key exists. Desired if that key/val is true.
    funcTypes := map[string]bool{"int":true, "double":true, "float":true, "boolean":true, "long":true,
                                 "short":true, "byte":true, "public":false, "private":false, "protected":false,
                                 "static":false, "strictfp":false, "native":false, "String":false, "void":false}
                              
    session := utils.ConnectDB()   
    search.SearchAndSaveFunc(session, searchDir, extension, funcTypes)
}