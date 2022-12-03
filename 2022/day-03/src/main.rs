use std::fs;
mod tests;

#[derive(Debug, PartialEq, Eq)]
struct Item(char);

impl Item {
    fn new(c: char) -> Self {
        assert!(c.is_ascii_alphabetic());
        Item(c)
    }

    fn priority(&self) -> u32 {
        if self.0.is_ascii_uppercase() {
            self.0 as u32 - 65 + 27
        } else {
            self.0 as u32 - 97 + 1
        }
    }
}

type Compartment = Vec<Item>;

#[derive(Debug)]
struct Rucksack {
    compartment_1: Compartment,
    compartment_2: Compartment,
}

impl Rucksack {
    fn new(items: &str) -> Self {
        let (compartment_1, compartment_2) = items.split_at(items.len() / 2);
        Rucksack {
            compartment_1: compartment_1.chars().map(Item::new).collect(),
            compartment_2: compartment_2.chars().map(Item::new).collect(),
        }
    }

    fn common_item(&self) -> Option<&Item> {
        self.compartment_1
            .iter()
            .find(|c| self.compartment_2.contains(c))
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let mut sum = 0;
    for line in contents.lines() {
        if line.is_empty() {
            break;
        }

        let rucksack = Rucksack::new(line);
        sum += rucksack.common_item().unwrap().priority();
    }

    println!("Part 1: {}", sum);
    assert_eq!(sum, 7824);
}
