package part1

import (
	"aoc02/part0"
	"strings"
)

func V1() (int, int){
    dat, err := part0.FileToMap("./input.txt")
    part0.Check(err)
    valid_passwords := 0
    n := 0
    for _, d := range dat {
        occurrences := 0
        n++
       if strings.Contains(d.Password, d.Letter) {
           for _,l := range d.Password {
               if string(l) == d.Letter {
                   n++
                   occurrences++
               }
           }
       }
       if occurrences >= d.Constraints[0] && occurrences <= d.Constraints[1] {
           valid_passwords++
       }
   }
    return valid_passwords, n
}

func V2() (int,int){
    dat, err := part0.FileToMap("./input.txt")
    part0.Check(err)
    valid_passwords := 0
    n := 0
    for _, d := range dat {
       occurrences := strings.Count(d.Password, d.Letter) // part0.Check how much does this take.
       n++
       if occurrences >= d.Constraints[0] && occurrences <= d.Constraints[1] {
           valid_passwords++
       }
   }
    return valid_passwords, n
}

