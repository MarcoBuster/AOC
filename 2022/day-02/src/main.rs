use std::fs;

mod tests;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone)]
enum Shape {
    Rock = 1,
    Paper = 2,
    Scissors = 3,
}

impl TryFrom<&str> for Shape {
    type Error = &'static str;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        match value {
            "A" | "X" => Ok(Self::Rock),
            "B" | "Y" => Ok(Self::Paper),
            "C" | "Z" => Ok(Self::Scissors),
            _ => Err("Can't derive a Shape from illegal string"),
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

impl TryFrom<&str> for Outcome {
    type Error = &'static str;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        match value {
            "X" => Ok(Self::Lose),
            "Y" => Ok(Self::Draw),
            "Z" => Ok(Self::Win),
            _ => Err("Can't derive a Shape from illegal string"),
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

        let opponent = Shape::try_from(opponent).unwrap();
        let my_shape = Shape::try_from(myself).unwrap();
        let prediction = Outcome::try_from(myself).unwrap();

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
