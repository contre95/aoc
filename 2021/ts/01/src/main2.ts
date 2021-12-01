import { log } from "console";
import { readFileSync } from "fs";

const input = readFileSync("src/testput", "utf-8")
  .toString()
  .split("\n")
  .map(Number);
const part = (input: number[], pos: number, x: number): number => {
  if (pos + x >= input.length) {
    return 0;
  }

  let a: number = 0;
  for (let i = pos; i < x; i++) {
    a += input[Number(i)];
  }
  console.log(a);
  console.log("yes");
  if (a < part(input, pos + 1, x + 1)) {
    return 1 + part(input, pos + 1, x + 1);
  }
};

console.log(part(input, 0, 3));
