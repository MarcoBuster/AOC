const LINES: [&'static str; 6] = [
    "vJrwpWtwJgWrhcsFMMfFFhFp",
    "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
    "PmmdzqPrVvPwwTWBwg",
    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
    "ttgJtRGJQctTZtZT",
    "CrZsJsPPZsGzwwsLwLmpwMDw",
];

#[cfg(test)]
mod part1 {
    use crate::{tests::LINES, Rucksack};

    fn get_priority(input: &str) -> Option<u32> {
        let rucksack = Rucksack::new(input);
        match rucksack.common_item() {
            Some(item) => Some(item.priority()),
            None => None,
        }
    }

    #[test]
    fn input_1() {
        assert_eq!(get_priority(LINES[0]), Some(16));
    }

    #[test]
    fn input_2() {
        assert_eq!(get_priority(LINES[1]), Some(38));
    }

    #[test]
    fn input_3() {
        assert_eq!(get_priority(LINES[2]), Some(42));
    }

    #[test]
    fn input_4() {
        assert_eq!(get_priority(LINES[3]), Some(22));
    }

    #[test]
    fn input_5() {
        assert_eq!(get_priority(LINES[4]), Some(20));
    }

    #[test]
    fn input_6() {
        assert_eq!(get_priority(LINES[5]), Some(19));
    }
}

#[cfg(test)]
mod part2 {
    use crate::{tests::*, Group};

    fn get_priority(line_1: &str, line_2: &str, line_3: &str) -> Option<u32> {
        let group = Group::new(line_1, line_2, line_3);
        match group.common_item() {
            Some(item) => Some(item.priority()),
            None => None,
        }
    }

    #[test]
    fn input_1() {
        assert_eq!(get_priority(LINES[0], LINES[1], LINES[2]), Some(18));
    }

    #[test]
    fn input_2() {
        assert_eq!(get_priority(LINES[3], LINES[4], LINES[5]), Some(52));
    }
}
