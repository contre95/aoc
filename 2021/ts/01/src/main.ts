import { readFileSync } from "fs";

const part1 = () => {
  const input = readFileSync("src/input", "utf-8").toString().split("\n");
  let increase: number = 0;
  for (var i in input) {
    let n1 = Number(input[i]);
    let n2 = Number(input[Number(i) + 1]);
    if (n1 < n2) {
      increase++;
    }
  }
  console.log(increase);
};

const part2 = () => {
  const input = readFileSync("src/input", "utf-8").toString().split("\n").map(Number)
  let increase: number = 0;
  for (var index in input) {
    let i = Number(index)
    let n1 = Number(input[i] + input[i + 1] + input[i + 2]);
    let n2 = Number(input[i + 1] + input[i + 2] + input[i + 3]);
    if (n1 < n2) {
      increase++;
    }
  }
  console.log(increase);
};

part1();
console.log('----')
part2();
