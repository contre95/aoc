use std::io;
use std::error::Error;
use std::collections::HashSet;
fn main() -> Result<(), Box<dyn Error>> {
    let mut numbers_seen = HashSet::new();
    loop {
        let mut line = String::new();
        io::stdin().read_line(&mut line)?;
        if line == "" {
            break;
        }
        let number = line.trim().parse::<i32>()?;
        let needed_number = 2020 - number;
        if numbers_seen.contains(&needed_number) {
            println!("{}", number * needed_number);
            return Ok(())
        }
        numbers_seen.insert(number);
    }
    Ok(())
}
