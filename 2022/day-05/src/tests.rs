use std::collections::HashMap;

use crate::Stack;

const LINES: [&'static str; 4] = [
    "move 1 from 2 to 1",
    "move 3 from 1 to 3",
    "move 2 from 2 to 1",
    "move 1 from 1 to 2",
];

fn get_crates() -> Vec<Stack> {
    let mut crates: Vec<Stack> = Vec::new();
    crates.push(
        Stack {
            number: 1,
            crates: vec!['Z', 'N'],
        },
    );
    crates.push(
        Stack {
            number: 2,
            crates: vec!['M', 'C', 'D'],
        },
    );
    crates.push(
        Stack {
            number: 3,
            crates: vec!['P'],
        },
    );
    crates
}

mod part1 {
    use crate::{Command, StackMapMethods};

    use super::{get_crates, LINES};

    #[test]
    fn test_command_1() {
        let mut crates = get_crates();
        Command::parse_string(LINES[0]).exec_part1(&mut crates);
        assert_eq!(crates.peek_str(), "DCP");
    }

    #[test]
    fn test_command_2() {
        let mut crates = get_crates();
        for i in 0..=1 {
            Command::parse_string(LINES[i]).exec_part1(&mut crates);
        }
        assert_eq!(crates.peek_str(), "CZ");
    }

    #[test]
    fn test_command_3() {
        let mut crates = get_crates();
        for i in 0..=2 {
            Command::parse_string(LINES[i]).exec_part1(&mut crates);
        }
        assert_eq!(crates.peek_str(), "MZ");
    }

    #[test]
    fn test_command_4() {
        let mut crates = get_crates();
        for i in 0..=3 {
            Command::parse_string(LINES[i]).exec_part1(&mut crates);
        }
        assert_eq!(crates.peek_str(), "CMZ");
    }
}

mod part2 {
    use crate::{Command, StackMapMethods};

    use super::{get_crates, LINES};

    #[test]
    fn test_command_1() {
        let mut crates = get_crates();
        Command::parse_string(LINES[0]).exec_part2(&mut crates);
        assert_eq!(crates.peek_str(), "DCP");
    }

    #[test]
    fn test_command_2() {
        let mut crates = get_crates();
        for i in 0..=1 {
            Command::parse_string(LINES[i]).exec_part2(&mut crates);
        }
        assert_eq!(crates.peek_str(), "CD");
    }

    #[test]
    fn test_command_3() {
        let mut crates = get_crates();
        for i in 0..=2 {
            Command::parse_string(LINES[i]).exec_part2(&mut crates);
        }
        assert_eq!(crates.peek_str(), "CD");
    }

    #[test]
    fn test_command_4() {
        let mut crates = get_crates();
        for i in 0..=3 {
            Command::parse_string(LINES[i]).exec_part2(&mut crates);
        }
        assert_eq!(crates.peek_str(), "MCD");
    }
}
