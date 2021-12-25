use std::collections::HashMap;
use std::io::BufRead;
use std::{fs, io};

fn lines_to_vector(path: &str) -> Vec<String> {
    let mut input: Vec<String> = vec![];
    let file = fs::File::open(path).expect("Error while reading the file");
    let reader = io::BufReader::new(file);
    for line in reader.lines() {
        let line = line.expect("Unable to read line");
        input.push(line)
    }
    input
}

pub fn part1() -> i32 {
    let input = lines_to_vector("./inputs/input_2");
    let mut directions = HashMap::new();

    for vector in input {
        let vector: Vec<&str> = vector.split(' ').collect();
        let mut direc: &str = vector[0];
        let mut module: i32 = vector[1].parse().expect("Could not parse the module");
        if direc == "up" {
            module = -module;
            direc = "down";
        }
        // What's the cost of this to_owned(), do i have choice?
        let count = directions.entry(direc.to_owned()).or_insert(0);
        *count += module
    }
    directions["forward"] * directions["down"]
}

pub fn part2() -> i32 {
    let input = lines_to_vector("./inputs/input_2");
    let mut directions = HashMap::new();

    for vector in input {
        let vector: Vec<&str> = vector.split(' ').collect();
        let mut direc: &str = vector[0];
        let mut module: i32 = vector[1].parse().expect("Could not parse the module");
        if direc == "up" {
            module = -module;
            direc = "down";
        }
        // What's the cost of this to_owned(), do i have choice?
        if direc == "down" {
            let count = directions.entry("aim").or_insert(0);
            *count += module;
        }
        //let count = directions.entry(direc.to_owned()).or_insert(0);
        //*count+=module;
        if direc == "forward" {
            if directions.contains_key("aim") {
                *directions.entry("down").or_insert(0) = module * *directions.get("aim").unwrap();
            }
            *directions.entry("forward").or_insert(0) += module
        }
    }
    directions["forward"] * directions["down"]
}

#[allow(dead_code)]
fn main() {
    println!("Day2");
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}
