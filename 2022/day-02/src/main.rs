use std::fs;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone)]
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

impl Outcome {
    fn from_str(string: &str) -> Option<Self> {
        match string {
            "X" => Some(Self::Lose),
            "Y" => Some(Self::Draw),
            "Z" => Some(Self::Win),
            _ => None,
        }
    }
}

#[derive(Debug, PartialEq, Eq)]
struct GameResult {
    outcome: Outcome,
    my_points: u8,
}

impl Game {
    fn from_prediction(opponent: Shape, outcome: Outcome) -> Self {
        match outcome {
            Outcome::Lose => Game {
                myself: match &opponent {
                    Shape::Rock => Shape::Scissors,
                    Shape::Paper => Shape::Rock,
                    Shape::Scissors => Shape::Paper,
                },
                opponent,
            },
            Outcome::Draw => Game {
                myself: opponent.clone(),
                opponent,
            },
            Outcome::Win => Game {
                myself: match &opponent {
                    Shape::Rock => Shape::Paper,
                    Shape::Paper => Shape::Scissors,
                    Shape::Scissors => Shape::Rock,
                },
                opponent,
            },
        }
    }

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

    let mut part_1: u32 = 0;
    let mut part_2: u32 = 0;
    for line in contents.split('\n') {
        if line.is_empty() {
            break;
        }

        let line = line.replace(' ', "");
        let (opponent, myself) = line.split_at(1);

        let opponent = Shape::from_str(opponent).unwrap();
        let my_shape = Shape::from_str(myself).unwrap();
        let prediction = Outcome::from_str(myself).unwrap();

        let game = Game {
            opponent: opponent.clone(),
            myself: my_shape,
        };
        let result = game.play();
        part_1 += result.my_points as u32;

        let predicted_game = Game::from_prediction(opponent, prediction);
        let result = predicted_game.play();
        part_2 += result.my_points as u32;
    }

    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 12458);

    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 12683);
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

    #[test]
    fn test_part2_ay() {
        let game = Game::from_prediction(Shape::Rock, Outcome::Draw);
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Draw,
                my_points: 4
            }
        );
    }

    #[test]
    fn test_part2_bx() {
        let game = Game::from_prediction(Shape::Paper, Outcome::Lose);
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Lose,
                my_points: 1
            }
        );
    }

    #[test]
    fn test_part2_cz() {
        let game = Game::from_prediction(Shape::Scissors, Outcome::Win);
        assert_eq!(
            game.play(),
            GameResult {
                outcome: Outcome::Win,
                my_points: 7
            }
        );
    }
}
