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


func fileToMap( fp string ) (map[int]int, error) {
    lines := make(map[int]int)
    f, err := os.Open(fp)
    check(err)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        i, err := strconv.Atoi(scanner.Text())
        check(err)
        lines[i] = 0
    }
    defer f.Close()
    return lines, scanner.Err()

}
func main() {
    dat, err := fileToMap("./intput.txt")
    check(err)
    var reports []int
    for i,_:= range dat {
        if _,found := dat[2020-i]; found {
            reports = append(reports,i*(2020-i))
        }
    }
    fmt.Println(reports)
}

