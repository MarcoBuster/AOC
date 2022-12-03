#[cfg(test)]
mod part1 {
    use crate::Rucksack;

    fn get_priority(input: &str) -> Option<u32> {
        let rucksack = Rucksack::new(input);
        match rucksack.common_item() {
            Some(item) => Some(item.priority()),
            None => None,
        }
    }

    #[test]
    fn input_1() {
        assert_eq!(get_priority("vJrwpWtwJgWrhcsFMMfFFhFp"), Some(16));
    }

    #[test]
    fn input_2() {
        assert_eq!(get_priority("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), Some(38));
    }

    #[test]
    fn input_3() {
        assert_eq!(get_priority("PmmdzqPrVvPwwTWBwg"), Some(42));
    }

    #[test]
    fn input_4() {
        assert_eq!(get_priority("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), Some(22));
    }

    #[test]
    fn input_5() {
        assert_eq!(get_priority("ttgJtRGJQctTZtZT"), Some(20));
    }

    #[test]
    fn input_6() {
        assert_eq!(get_priority("CrZsJsPPZsGzwwsLwLmpwMDw"), Some(19));
    }
}
