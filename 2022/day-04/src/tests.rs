const LINES: [&'static str; 6] = [
    "2-4,6-8",
    "2-3,4-5",
    "5-7,7-9",
    "2-8,3-7",
    "6-6,4-6",
    "2-6,4-8"
];

mod part1 {
    use crate::{ElfPair, tests::LINES};

    #[test]
    fn input_1() {
        assert_eq!(ElfPair::new(LINES[0]).one_fully_contains_other(), false);
    }

    #[test]
    fn input_2() {
        assert_eq!(ElfPair::new(LINES[1]).one_fully_contains_other(), false);
    }

    #[test]
    fn input_3() {
        assert_eq!(ElfPair::new(LINES[2]).one_fully_contains_other(), false);
    }

    #[test]
    fn input_4() {
        assert_eq!(ElfPair::new(LINES[3]).one_fully_contains_other(), true);
    }

    #[test]
    fn input_5() {
        assert_eq!(ElfPair::new(LINES[4]).one_fully_contains_other(), true);
    }

    #[test]
    fn input_6() {
        assert_eq!(ElfPair::new(LINES[5]).one_fully_contains_other(), false);
    }
}

mod part2 {
    use crate::{ElfPair, tests::LINES};

    #[test]
    fn input_1() {
        assert_eq!(ElfPair::new(LINES[0]).overlaps(), false);
    }

    #[test]
    fn input_2() {
        assert_eq!(ElfPair::new(LINES[1]).overlaps(), false);
    }

    #[test]
    fn input_3() {
        assert_eq!(ElfPair::new(LINES[2]).overlaps(), true);
    }

    #[test]
    fn input_4() {
        assert_eq!(ElfPair::new(LINES[3]).overlaps(), true);
    }

    #[test]
    fn input_5() {
        assert_eq!(ElfPair::new(LINES[4]).overlaps(), true);
    }

    #[test]
    fn input_6() {
        assert_eq!(ElfPair::new(LINES[5]).overlaps(), true);
    }
}