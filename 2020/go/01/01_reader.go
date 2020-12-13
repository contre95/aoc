package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func fileToSlice( fp string ) ([]int, error) {
    var lines []int
    f, err := os.Open(fp)
    check(err)
    r := bufio.NewReader(f)
    for {
        l, err := r.ReadString('\n')
        if err != nil && err == io.EOF {
            break
        }
        check(err)
        i, err := strconv.Atoi(strings.TrimSuffix(l, "\n"))
        check(err)
        lines = append(lines, i)
    }
    defer f.Close()
    return lines, nil
}
func main() {
    dat, err := fileToSlice("./intput.txt")
    check(err)
    var reports []int

    for i,_ := range dat {
        for j,_ := range dat {
            if j<=i {
                continue
            }
            r1 := dat[i]
            r2 := dat[j]
            if i!=j && r1+r2 == 2020 {
                reports = append(reports, r1*r2)
            }
        }
    }
    fmt.Println(reports)
}

