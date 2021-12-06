"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var part1 = function () {
    var input = (0, fs_1.readFileSync)('./input');
    for (var i = 0; i < input.length; i++) {
        var element = input[i];
        console.log(element);
    }
};
var part2 = function () {
};
part1();
console.log('----');
part2();
