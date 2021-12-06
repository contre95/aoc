import { readFileSync } from "fs";


type Counter = { [id: string]: number }

const part1 = () => {
  const input = readFileSync("src/input", "utf-8").toString().split("\n");
  let c: Counter = {}
  for (let line = 0; line < input.length - 1; line++) {
    for (let char = 0; char < input[0].length; char++) {
      if (c[char] == undefined) {
        c[char] = 0
      }
      if (input[line][char] == "1") {
      } else {
        c[char] += 1
      }
    }
  }
  let gamma: string = ""
  let epsilon: string = ""
  console.log(c)
  for (let key = 0; key < input[0].length; key++) {
    if (c[key.toString()] < input.length / 2) {
      gamma += "1"
      epsilon += "0"
    } else {
      gamma += "0"
      epsilon += "1"
    }
  }
  console.log(epsilon)
  console.log(gamma)
  console.log(parseInt(epsilon, 2) * parseInt(gamma, 2))
};

const part2 = () => {
  console.log("");
};

part1();
console.log("----");
part2();
