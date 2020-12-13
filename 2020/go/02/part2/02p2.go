package part2

import (
	"aoc02/part0"
)



func V1() (int,int){
    dat, err := part0.FileToMap("input.txt")
    part0.Check(err)
    valid_passwords := 0
    n := 0
    for _, d := range dat {
        n++
        if ( string(d.Password[d.Constraints[0]-1]) == d.Letter ) != (string(d.Password[d.Constraints[1]-1]) == d.Letter) {
            valid_passwords++
        }
    }
    return valid_passwords, n
}
