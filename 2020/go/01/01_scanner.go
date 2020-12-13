package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        i, err := strconv.Atoi(scanner.Text())
        check(err)
        lines = append(lines, i)
    }
    defer f.Close()
    return lines, scanner.Err()

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

