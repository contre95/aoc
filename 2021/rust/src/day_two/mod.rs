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
    return input;
}

pub fn part1() -> u32 {  
    let input = lines_to_vector("src/day_two/testput");
    for vector in input {
        let direction = vector.split(' ');
        println!("{:?}",direction)
    }
    21
}

