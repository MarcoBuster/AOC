use std::fs;

#[derive(PartialEq, Eq, PartialOrd, Ord)]
struct Elf {
    calories: u32
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let mut elves: Vec<Elf> = vec![];
    let mut count: u32 = 0;
    for line in contents.split('\n') {
        if line.is_empty() {
            elves.push(Elf { calories: count });
            count = 0;
            continue;
        }
        let number: u32 = line.parse().unwrap_or_else(|_| panic!("Line {line} is not a number."));
        count += number;
    }

    println!("Part 1: {}", elves.iter().max().unwrap().calories);

    elves.sort();
    elves.reverse();

    println!("Part 2: {}", elves.iter().take(3).map(|e| e.calories).sum::<u32>());
}
