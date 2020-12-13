fs = require("fs");

let main = () => {
  let a = [];
  var dat = fs.readFileSync("intput.txt", "utf8").split("\n").map(x=>Number(x));
  let set = new Set(dat);
  for (i = 0; i <= dat.length; i++) {
    if (set.has(2020 - dat[i])) {
      a.push((2020 - dat[i]) * dat[i])
    }
  }
  return a;
};

console.log(main())
