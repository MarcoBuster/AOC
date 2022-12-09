use std::{collections::HashSet, fs};

#[cfg(test)]
mod tests;

#[derive(Debug)]
enum Direction {
    Left,
    Right,
    Up,
    Down,
}

impl Direction {
    fn from_string(value: &str) -> Self {
        match value {
            "L" => Self::Left,
            "R" => Self::Right,
            "U" => Self::Up,
            "D" => Self::Down,
            _ => panic!("Invalid direction"),
        }
    }
}

#[derive(Debug)]
struct Motion {
    direction: Direction,
    step: usize,
}

impl Motion {
    fn from_string(value: &str) -> Self {
        let (direction, step) = value.split_once(' ').expect("Invalid command");
        Self {
            direction: Direction::from_string(direction),
            step: step.parse().unwrap(),
        }
    }
}

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Position(isize, isize);

impl Position {
    fn start() -> Self {
        Self(0, 0)
    }
}

#[derive(Debug)]
struct State {
    head: Position,
    knots: Vec<Position>,
    visited: HashSet<Position>,
}

impl State {
    fn new(n_knots: usize) -> Self {
        assert!(n_knots > 1);
        Self {
            head: Position::start(),
            knots: vec![Position::start(); n_knots - 1],
            visited: HashSet::new(),
        }
    }

    fn align_knot(&mut self, knot_idx: usize) {
        assert!(knot_idx < self.knots.len());

        let knot = self.knots.get(knot_idx).unwrap();
        let prec = if knot_idx > 0 {
            self.knots.get(knot_idx - 1).unwrap()
        } else {
            &self.head
        };

        if knot.0.abs_diff(prec.0) <= 1 && knot.1.abs_diff(prec.1) <= 1 {
            return;
        }
        let delta_x = prec.0 - knot.0;
        let delta_y = prec.1 - knot.1;

        let knot = self.knots.get_mut(knot_idx).unwrap();

        knot.0 += match delta_x.cmp(&0) {
            std::cmp::Ordering::Less => -1,
            std::cmp::Ordering::Equal => 0,
            std::cmp::Ordering::Greater => 1,
        };
        knot.1 += match delta_y.cmp(&0) {
            std::cmp::Ordering::Less => -1,
            std::cmp::Ordering::Equal => 0,
            std::cmp::Ordering::Greater => 1,
        };
    }

    fn perform(&mut self, motion: &Motion) {
        for _ in 0..motion.step {
            match motion.direction {
                Direction::Left => self.head.0 -= 1,
                Direction::Right => self.head.0 += 1,
                Direction::Up => self.head.1 -= 1,
                Direction::Down => self.head.1 += 1,
            }

            for knot_idx in 0..self.knots.len() {
                self.align_knot(knot_idx);
            }

            self.visited.insert(self.knots.last().unwrap().clone());
        }
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Can't read input file");

    let mut state = State::new(2);
    contents
        .lines()
        .map(Motion::from_string)
        .for_each(|motion| state.perform(&motion));

    let part_1 = state.visited.len();
    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 6745);

    let mut state = State::new(10);
    contents
        .lines()
        .map(Motion::from_string)
        .for_each(|motion| state.perform(&motion));

    let part_2 = state.visited.len();
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 2793);
}
