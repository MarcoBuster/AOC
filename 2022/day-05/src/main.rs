#[cfg(test)]
mod tests;

use std::fs;
use regex::Regex;

type Crate = char;

fn parse_crate(value: &str) -> Crate {
    value.chars().nth(1).unwrap()
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
struct Stack {
    crates: Vec<Crate>,
    number: usize,
}

impl Stack {
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
    source: usize,
    destination: usize,
    quantity: usize,
}

type StackList = Vec<Stack>;

trait StackMapMethods {
    fn peek_str(&self) -> String;
}

impl StackMapMethods for StackList {
    fn peek_str(&self) -> String {
        let mut result = "".to_string();
        for stack in self.iter() {
            if let Some(crate_) = stack.peek() {
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
            source: captures[2].parse::<usize>().unwrap(),
            destination: captures[3].parse::<usize>().unwrap(),
            quantity: captures[1].parse::<usize>().unwrap(),
        }
    }

    fn exec_part1(&self, stacks: &mut [Stack]) {
        for _i in 0..self.quantity {
            let source = stacks.get_mut(self.source - 1).unwrap();
            let element = source.pop();
            let dest = stacks.get_mut(self.destination - 1).unwrap();
            dest.push(element);
        }
    }

    fn exec_part2(&self, stacks: &mut [Stack]) {
        let source = stacks.get_mut(self.source - 1).unwrap();
        let mut elements: Vec<Crate> = Vec::new();
        for _i in 0..self.quantity {
            elements.push(source.pop());
        }
        elements.reverse();
        let dest = stacks.get_mut(self.destination - 1).unwrap();
        dest.extend(elements);
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let first_line = contents.lines().next().unwrap();
    let mut stacks: Vec<Stack> = Vec::new();
    for i in 0..(first_line.len() + 1) / 4 {
        stacks.push(Stack { crates: vec![], number: i + 1 })
    }

    let mut parse_stacks = true;
    let mut stacks_part2: Vec<Stack> = Vec::new();
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

                let key = num_chars / 4;
                let crate_name = parse_crate(group);
                stacks.get_mut(key - 1).unwrap().insert_last(crate_name);
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
