use std::collections::HashMap;
use std::fs;
use std::io::Result;

const START_OF_MESSAGE_MARKER_PART_1: usize = 4;
const START_OF_MESSAGE_MARKER_PART_2: usize = 14;

trait StrCustomMethods {
    fn are_chars_all_different(&self) -> bool;
}

struct Seen;

impl StrCustomMethods for &str {
    fn are_chars_all_different(&self) -> bool {
        let mut map: HashMap<char, Seen> = HashMap::new();
        for c in self.chars() {
            match map.get(&c) {
                Some(_) => return false,
                None => { map.insert(c.clone(), Seen); },
            }
        }
        true
    }
}


fn main() -> Result<()> {
    let contents = fs::read_to_string("input.txt")?;
    let line = contents.lines().next().expect("The file is empty!");

    let mut part_1 = 0;
    for i in 0..line.len()-START_OF_MESSAGE_MARKER_PART_1 {
        let slice = &line[i..i+START_OF_MESSAGE_MARKER_PART_1];
        if slice.are_chars_all_different() {
            part_1 = i + START_OF_MESSAGE_MARKER_PART_1;
            break;
        }
    }
    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 1542);

    let mut part_2 = 0;
    for i in 0..line.len()-START_OF_MESSAGE_MARKER_PART_2 {
        let slice = &line[i..i+START_OF_MESSAGE_MARKER_PART_2];
        if slice.are_chars_all_different() {
            part_2 = i + START_OF_MESSAGE_MARKER_PART_2;
            break;
        }
    }
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 3153);

    Ok(())
}
