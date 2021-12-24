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

pub fn part1() -> u32 {
    let input = lines_to_vector("./inputs/testput_2");
    
    let directions : HashMap<&str,i32> = HashMap::new();
    for vector in input {
        let vector: Vec<&str> = vector.split(' ').collect();
        let mut direc: &str = vector[0];
        let mut module: i32 = vector[1].parse().expect("Could not parse the module");
        if direc == "up" {
            module = -module;
            direc = "down";
        }
        //let direc_wrong: Vec<String> = vector.split(' ').map(|x| x.to_string()).collect(); // Just if I wanted them to be a string.
        println!("{} - {}", direc, module);
    }
    21
}

#[allow(dead_code)]
fn main() {
    println!("Day2");
    println!("Part1: {}", part1());
    //println!("Part2: {}", part2());
}
