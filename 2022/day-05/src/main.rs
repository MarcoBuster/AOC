#[cfg(test)]
mod tests;

use itertools::Itertools;
use std::{collections::HashMap, fs};

use regex::Regex;

type Crate = char;

fn parse_crate(value: &str) -> Crate {
    value.chars().nth(1).unwrap()
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
struct Stack {
    crates: Vec<Crate>,
    number: u8,
}

impl Stack {
    fn new(number: u8) -> Self {
        Stack {
            crates: vec![],
            number,
        }
    }

    fn push(&mut self, crate_: Crate) {
        self.crates.push(crate_);
    }

    fn insert_last(&mut self, crate_: Crate) {
        self.crates.insert(0, crate_);
    }

    fn extend(&mut self, crates: Vec<Crate>) {
        self.crates.extend(crates);
    }

    fn pop(&mut self) -> Crate {
        assert!(!self.crates.is_empty());
        self.crates.pop().unwrap()
    }

    fn peek(&self) -> Option<&Crate> {
        self.crates.last()
    }
}

#[derive(Debug)]
struct Command {
    source: u8,
    destination: u8,
    quantity: u8,
}

type StackMap = HashMap<u8, Stack>;

trait StackMapMethods {
    fn peek_str(&self) -> String;
}

impl StackMapMethods for StackMap {
    fn peek_str(&self) -> String {
        let mut result = "".to_string();
        for stack in self.iter().sorted() {
            if let Some(crate_) = stack.1.peek() {
                result = format!("{}{}", result, crate_);
            }
        }
        result
    }
}

impl Command {
    fn parse_string(value: &str) -> Self {
        let re = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
        let captures = re.captures_iter(value).next().unwrap();
        Command {
            source: captures[2].parse::<u8>().unwrap(),
            destination: captures[3].parse::<u8>().unwrap(),
            quantity: captures[1].parse::<u8>().unwrap(),
        }
    }

    fn exec_part1(&self, stacks: &mut HashMap<u8, Stack>) {
        for _i in 0..self.quantity {
            let source = stacks.get_mut(&self.source).unwrap();
            let element = source.pop();
            let dest = stacks.get_mut(&self.destination).unwrap();
            dest.push(element);
        }
    }

    fn exec_part2(&self, stacks: &mut HashMap<u8, Stack>) {
        let source = stacks.get_mut(&self.source).unwrap();
        let mut elements: Vec<Crate> = Vec::new();
        for _i in 0..self.quantity {
            elements.push(source.pop());
        }
        elements.reverse();
        let dest = stacks.get_mut(&self.destination).unwrap();
        dest.extend(elements);
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let mut parse_stacks = true;
    let mut stacks: HashMap<u8, Stack> = HashMap::new();
    let mut stacks_part2: HashMap<u8, Stack> = HashMap::new();
    for line in contents.lines() {
        if line.is_empty() {
            if !parse_stacks {
                break;
            }

            parse_stacks = false;
            stacks_part2 = stacks.clone();
            continue;
        }

        if parse_stacks {
            if line.contains('1') {
                continue;
            }

            let mut num_chars = 0;
            for group in line.split_inclusive(' ') {
                num_chars += group.len();
                if group == " " {
                    continue;
                }

                if num_chars % 4 == 3 {
                    num_chars += 1;
                }

                let key = (num_chars / 4) as u8;
                let crate_name = parse_crate(group);
                match stacks.get_mut(&key) {
                    Some(stack) => stack.insert_last(crate_name),
                    None => {
                        let mut stack = Stack::new(key);
                        stack.push(crate_name);
                        stacks.insert(key, stack);
                    }
                };
            }
        } else {
            let command = Command::parse_string(line);
            command.exec_part1(&mut stacks);
            command.exec_part2(&mut stacks_part2);
        }
    }

    let part_1 = stacks.peek_str();
    println!("Part 1: {}", part_1);
    assert_eq!(part_1, "QGTHFZBHV");

    let part_2 = stacks_part2.peek_str();
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, "MGDMPSZTM");
}
