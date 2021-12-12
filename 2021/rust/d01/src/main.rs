use std::fs;
use std::io::{BufRead, BufReader};

fn lines_to_vector(path: &str) -> Vec<String> {
    println!("{}", path);
    let mut input: Vec<String> = vec![];
    let file = fs::File::open(path).expect("Error while reading the file");
    let reader = BufReader::new(file);
    for line in reader.lines() {
        let line = line.expect("Unable to read line");
        input.push(line)
    }
    return input;
}

fn part1()-> u32 {
    let input = lines_to_vector("input");
    let mut count = 0;
    let mut increase : u32 = 0;
    println!("Len {}", input.len());
    loop {
        let n1: u32 = input[count].parse().expect("Could not parse");
        let n2: u32 = input[count + 1].parse().expect("Could not parse");
        if n1 < n2 {
            increase += 1;
        }
        count += 1;
        if count + 1 >= input.len() {
            break increase
        }
    }
}


fn main(){ 
    println!("Part1: {}", part1());
}
