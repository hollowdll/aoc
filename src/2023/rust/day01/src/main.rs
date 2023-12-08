use std::fs::File;
use std::io::{
    self,
    Lines,
    BufReader,
    BufRead,
};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut sum = 0;

    if let Ok(lines) = read_lines("./input/input.txt") {
        for line in lines {
            if let Ok(line) = line {
                let mut first_digit = 0;
                let mut last_digit = 0;

                for c in line.chars() {
                    if c.is_digit(10) {
                        first_digit = c.to_digit(10).unwrap() as i32;
                        break;
                    }
                }

                for c in line.chars().rev() {
                    if c.is_digit(10) {
                        last_digit = c.to_digit(10).unwrap() as i32;
                        break;
                    }
                }

                sum += format!("{}{}", first_digit, last_digit).parse::<i32>().unwrap();
            }
        }
    }

    println!("Part 1 answer: {}", sum);
}

fn part2() {
    let letter_digits = [
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    ];
    let mut sum = 0;

    if let Ok(lines) = read_lines("./input/input.txt") {
        for line in lines {
            if let Ok(line) = line {
                let mut digits = Vec::new();

                for (i, c) in line.chars().enumerate() {
                    if c.is_digit(10) {
                        digits.push(c.to_digit(10).unwrap() as usize);
                    }
                    for (j, val) in letter_digits.iter().enumerate() {
                        if line[i..].starts_with(val) {
                            digits.push(j + 1);
                        }
                    }
                }

                sum += format!("{}{}", digits[0], digits[digits.len() - 1]).parse::<i32>().unwrap();
            }
        }
    }

    println!("Part 2 answer: {}", sum);
}

// https://doc.rust-lang.org/rust-by-example/std_misc/file/read_lines.html#a-more-efficient-approach
fn read_lines<P>(filename: P) -> io::Result<Lines<BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(BufReader::new(file).lines())
}
