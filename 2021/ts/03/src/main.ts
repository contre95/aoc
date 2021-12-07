import { readFileSync } from "fs";


type Counter = { [id: string]: number }

const part1 = (): number => {
  const input: string[] = readFileSync("src/input", "utf-8").toString().split("\n");

  let gamma: string = ""
  let epsilon: string = ""
  const cant_cols: number = input[0].length;
  for (let col = 0; col < cant_cols; col++) {
    if (hasMore(input, col) == "1") {
      gamma += "1"
      epsilon += "0"
    } else {
      gamma += "0"
      epsilon += "1"
    }
  }
  return parseInt(epsilon, 2) * parseInt(gamma, 2)
};

const hasMore = (input: string[], i: number, check: string = "0"): string => {
  let count = 0
  for (const line of input) {
    if (line[i] == check) {
      count++
    }
  }
  let result: string
  result = count >= input.length / 2 ? "1" : "0"
  if (check == "0") {
    result = count > input.length / 2 ? "1" : "0"
  }
  return result
}

const removeItemsWithChecks = (input: string[], col: number, check: string): string[] => {
  let temp = input
  temp = temp.filter((word) => { return word.length > 0 })
  console.log("Remove every item on input which col : ", col, " has the value ", check)
  for (const i in temp) {
    if (temp[i][col] == check) {
      delete temp[i]
    }
  }
  temp = temp.filter((word) => { return word.length >= 0 })
  return temp
}


const part2a = () => {
  let input: string[] = readFileSync("src/input", "utf-8").toString().split("\n");
  let column = 0
  do {
    input = removeItemsWithChecks(input, column, hasMore(input, column, "1"))
    column++
  } while (input.length > 1);
  console.log("temp: ", input[0])
  console.log("temp: ", parseInt(input[0], 2))
  return input[0]
};
const part2b = () => {
  let input: string[] = readFileSync("src/input", "utf-8").toString().split("\n");
  let column = 0
  do {
    input = removeItemsWithChecks(input, column, hasMore(input, column))
    column++
  } while (input.length > 1);
  console.log("temp: ", input.length)
  console.log("temp: ", input[0])
  console.log("temp: ", parseInt(input[0], 2))
  return input[0]
};

const part2 = () => {
  return parseInt(part2a(), 2) * parseInt(part2b(), 2)
}

console.log();
console.log("Part 1: ", part1());
console.log("----");
console.log("Part 2: ", part2());
//console.log("Part 2: ", part2());
