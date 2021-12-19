use std::fs;
use std::io::{BufRead, BufReader};

fn lines_to_vector(path: &str) -> Vec<String> {
    let mut input: Vec<String> = vec![];
    let file = fs::File::open(path).expect("Error while reading the file");
    let reader = BufReader::new(file);
    for line in reader.lines() {
        let line = line.expect("Unable to read line");
        input.push(line)
    }
    input
}

pub fn part1() -> u32 {
    let input = lines_to_vector("./inputs/testput_2");
    for vector in input {
        println!("{}", vector)
    }
    21
}

fn main() {
    println!("Day2");
    println!("Part1: {}", part1());
    //println!("Part2: {}", part2());
}
