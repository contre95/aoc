use std::fs;
use std::io::{BufRead, BufReader};

fn lines_to_vector(path: &str) -> Vec<u32> {
    let mut input: Vec<u32> = vec![];
    let file = fs::File::open(path).expect("Error while reading the file");
    let reader = BufReader::new(file);
    for line in reader.lines() {
        let line = line.expect("Unable to read line");
        input.push(line.parse().expect("Could not parse string to u32"))
    }
    return input;
}

pub fn part1() -> u32 {
    let input = lines_to_vector("./inputs/input_1");
    let mut count = 0;
    let mut increase: u32 = 0;
    loop {
        let n1: u32 = input[count];
        let n2: u32 = input[count + 1];
        if n1 < n2 {
            increase += 1;
        }
        count += 1;
        if count + 1 >= input.len() {
            break increase;
        }
    }
}

pub fn part2() -> u32 {
    let input = lines_to_vector("./inputs/input_1");
    let mut count = 0;
    let mut increase: u32 = 0;
    while count + 4 <= input.len() {
        let n1: u32 = input[count] + input[count + 1] + input[count + 2];
        let n2: u32 = input[count + 1] + input[count + 2] + input[count + 3];
        if n1 < n2 {
            increase += 1;
        }
        count += 1;
    }
    increase
}
