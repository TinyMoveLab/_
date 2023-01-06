package main

import (
    "fmt"
    "os"
	"time"
	"io"
    "bufio"
    // "log"
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

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

// func main() {
//     lines, err := readLines("foo.in.txt")
//     if err != nil {
//         log.Fatalf("readLines: %s", err)
//     }
//     for i, line := range lines {
//         fmt.Println(i, line)
//     }

//     if err := writeLines(lines, "foo.out.txt"); err != nil {
//         log.Fatalf("writeLines: %s", err)
//     }
// }
	

func main() {
    argLen := len(os.Args)
    // fmt.Println(argLen)
    if argLen == 1 {
        lines, _ := readLines("._")
        for i, line := range lines {
                    fmt.Println(i, line)
                    dateTimeNow := string(time.Now().AddDate(543, 0, 0).Format("06.01.02.15.04.05"))
                    os.Mkdir("_"+line, os.ModePerm);
                    copyFileContents("./"+line,"./_"+line+"/"+dateTimeNow+"_"+line)    
                    
        }
    }else{
        arg := os.Args[1]
        dateTimeNow := string(time.Now().AddDate(543, 0, 0).Format("06.01.02.15.04.05"))
        if arg == "init" {
            writeLines([]string{},"._")
        }else{
            os.Mkdir("_"+arg, os.ModePerm);
            copyFileContents("./"+arg,"./_"+arg+"/"+dateTimeNow+"_"+arg)       
        }
    }
    
	


    
	
}