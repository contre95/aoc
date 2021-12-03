import { readFileSync } from "fs";

const part1 = () => {
  let directions: { [key: string]: number } = {};
  const input = readFileSync("input", "utf-8").toString().split("\n");
  for (let d of input) {
    let direction: string = d.split(" ")[0];
    let mod: number = Number(d.split(" ")[1]);
    if (direction == "up") {
      mod = -mod;
      direction = "down";
    }
    if (direction in directions) {
      directions[direction] += mod;
    } else {
      directions[direction] = mod;
    }
  }
  return directions["down"] * directions["forward"];
};

const part2 = () => {
  const input = readFileSync("input", "utf-8").toString().split("\n");
  let directions = {
    aim: 0,
    forward: 0,
    down: 0,
  };
  for (let d of input) {
    let direction: string = d.split(" ")[0];
    let mod: number = Number(d.split(" ")[1]);
    if (direction == "up") {
      mod = -mod;
      direction = "down";
    }
    console.log(direction, mod);
    if (direction == "down") {
      directions.aim += mod;
    }
    if (direction == "forward") {
      directions.down += mod * directions.aim;
      directions.forward += mod;
    }
  }
  return directions["down"] * directions["forward"];
};

console.log(part1())
console.log("----");
console.log(part2())
