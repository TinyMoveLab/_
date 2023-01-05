package main

import (
    // "fmt"
    "os"
	"time"
	"io"
)

func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}

	

func main() {
    arg := os.Args[1]
	
	dateTimeNow := string(time.Now().AddDate(543, 0, 0).Format("06.01.02.15.04.05"))

	// if err := os.Mkdir("_"+arg, os.ModePerm); err != nil {
    //     fmt.Println(err)
    // }

	os.Mkdir("_"+arg, os.ModePerm);
	copyFileContents("./"+arg,"./_"+arg+"/"+dateTimeNow+"_"+arg)
	
	
}