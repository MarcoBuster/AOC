use std::{fs, ops::RangeInclusive};

#[cfg(test)]
mod tests;

#[derive(Debug)]
struct Section(RangeInclusive<u32>);

impl Section {
    fn new(value: &str) -> Self {
        let (l, r) = value.split_once("-").unwrap();
        let l: u32 = l.parse().unwrap();
        let r: u32 = r.parse().unwrap();
        Section(l..=r)
    }

    fn fully_contains(&self, other: &Self) -> bool {
        self.0.start() <= other.0.start() && self.0.end() >= other.0.end()
    }
}

#[derive(Debug)]
struct ElfPair(Section, Section);

impl ElfPair {
    fn new(value: &str) -> Self {
        let (l, r) = value.split_once(",").unwrap();
        ElfPair(Section::new(l), Section::new(r))
    }

    fn overlaps(&self) -> bool {
        self.0.0.clone().filter(|x| self.1.0.contains(&x)).peekable().peek().is_some()
    }

    fn one_fully_contains_other(&self) -> bool {
        self.0.fully_contains(&self.1) || self.1.fully_contains(&self.0)
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();
    
    let mut part_1 = 0;
    let mut part_2 = 0;
    for line in contents.lines() {
        if line.is_empty() {
            break;
        }

        let pair = ElfPair::new(line);
        if pair.one_fully_contains_other() {
            part_1 += 1;
        }
        if pair.overlaps() {
            part_2 += 1;
        }
    }

    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 528);
    
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 881);
}
