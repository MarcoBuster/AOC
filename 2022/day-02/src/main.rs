use std::fs;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
enum Shape {
    Rock = 1,
    Paper = 2,
    Scissors = 3,
}

impl Shape {
    fn from_str(string: &str) -> Option<Self> {
        match string {
            "A" | "X" => Some(Self::Rock),
            "B" | "Y" => Some(Self::Paper),
            "C" | "Z" => Some(Self::Scissors),
            _ => None,
        }
    }
}

#[derive(Debug)]
struct Game {
    opponent: Shape,
    myself: Shape,
}

#[derive(Debug, PartialEq, Eq)]
enum Outcome {
    Lose,
    Draw,
    Win,
}

#[derive(Debug, PartialEq, Eq)]
struct GameResult {
    outcome: Outcome,
    my_points: u8,
}

impl Game {
    fn play(self) -> GameResult {
        match (&self.opponent, &self.myself) {
            (Shape::Rock, Shape::Scissors)
            | (Shape::Scissors, Shape::Paper)
            | (Shape::Paper, Shape::Rock) => GameResult {
                outcome: Outcome::Lose,
                my_points: self.myself as u8,
            },

            (Shape::Rock, Shape::Rock)
            | (Shape::Paper, Shape::Paper)
            | (Shape::Scissors, Shape::Scissors) => GameResult {
                outcome: Outcome::Draw,
                my_points: self.myself as u8 + 3,
            },

            (Shape::Scissors, Shape::Rock)
            | (Shape::Paper, Shape::Scissors)
            | (Shape::Rock, Shape::Paper) => GameResult {
                outcome: Outcome::Win,
                my_points: self.myself as u8 + 6,
            },
        }
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let mut total_points: u32 = 0;
    for line in contents.split('\n') {
        if line.is_empty() {
            break;
        }

        let line = line.replace(' ', "");
        let (opponent, myself) = line.split_at(1);

        let opponent = Shape::from_str(opponent).unwrap();
        let myself = Shape::from_str(myself).unwrap();

        let game = Game { opponent, myself };
        let winner = game.play();
        total_points += winner.my_points as u32;
    }

    assert_eq!(total_points, 12458);
    println!("Part 1: {}", total_points);
}

#[cfg(test)]
mod tests {
    use crate::{Game, GameResult, Outcome, Shape};

    #[test]
    fn test_part1_ay() {
        let game = Game {
            opponent: Shape::Rock,
            myself: Shape::Paper,
        };
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Win,
                my_points: 8
            }
        );
    }

    #[test]
    fn test_part1_bx() {
        let game = Game {
            opponent: Shape::Paper,
            myself: Shape::Rock,
        };
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Lose,
                my_points: 1
            }
        );
    }

    #[test]
    fn test_part1_cz() {
        let game = Game {
            opponent: Shape::Scissors,
            myself: Shape::Scissors,
        };
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Draw,
                my_points: 6
            }
        );
    }
}
