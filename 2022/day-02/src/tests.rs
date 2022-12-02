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
