use eval::Expr;
use std::fs;

#[cfg(test)]
mod tests;

type Item = u64;

#[derive(Debug, Clone)]
struct Monkey {
    items: Vec<Item>,
    operation: Expr,
    divisible_by: u64,
    if_true_dest: u64,
    if_false_dest: u64,
    inspections: u64,
}

impl Monkey {
    fn parse(value: String) -> Self {
        let mut lines = value.lines();
        lines.next();
        Monkey {
            items: lines
                .next()
                .expect("Monkey definition group has too few lines")
                .split_once(": ")
                .expect("'Starting items' not properly delimited")
                .1
                .split(", ")
                .map(|item| item.parse().expect("This starting item is not a number"))
                .collect(),
            operation: Expr::new(
                lines
                    .next()
                    .expect("Monkey definition group has too few lines")
                    .split_once("= ")
                    .expect("'Operation' not properly delimited")
                    .1,
            )
            .compile()
            .expect("Invalid operation expression"),
            divisible_by: lines
                .next()
                .expect("Monkey definition group has too few lines")
                .split_once("divisible by ")
                .expect("Expected a 'divisible by' condition")
                .1
                .parse()
                .expect("The operator is not a number"),
            if_true_dest: lines
                .next()
                .expect("Monkey definition group has too few lines")
                .split_once("monkey ")
                .expect("'If true' not properly delimited")
                .1
                .parse()
                .expect(
                    "The selected monkey for the 'if true' clause is not identified with a number",
                ),
            if_false_dest: lines
                .next()
                .expect("Monkey definition group has too few lines")
                .split_once("monkey ")
                .expect("'If false' not properly delimited")
                .1
                .parse()
                .expect(
                    "The selected monkey for the 'if false' clause is not identified with a number",
                ),
            inspections: 0,
        }
    }
}

// TODO: Find a better name (lol)
#[derive(Debug)]
struct MonkeyManager {
    monkeys: Vec<Monkey>,
    common_multiplier: u64,
}

impl MonkeyManager {
    fn new() -> Self {
        MonkeyManager {
            monkeys: vec![],
            common_multiplier: 1,
        }
    }

    fn insert(&mut self, monkey: Monkey) {
        self.common_multiplier *= monkey.divisible_by;
        self.monkeys.push(monkey);
    }

    fn turn(&mut self, index: usize, divide_by_three: bool) {
        let monkey = self.monkeys[index].clone();
        let item_count = monkey.items.len();

        for item in monkey.items {
            let mut worry_level = monkey
                .operation
                .clone()
                .value("old", item)
                .exec()
                .unwrap()
                .as_u64()
                .unwrap();

            if divide_by_three {
                worry_level /= 3;
            } else {
                worry_level %= self.common_multiplier;
            }

            let other = self
                .monkeys
                .get_mut(if worry_level % monkey.divisible_by == 0 {
                    monkey.if_true_dest
                } else {
                    monkey.if_false_dest
                } as usize)
                .unwrap();
            other.items.push(worry_level);
        }

        self.monkeys[index] = Monkey {
            items: vec![],
            inspections: monkey.inspections + (item_count as u64),
            ..monkey
        };
    }

    fn round(&mut self, divide_by_three: bool) {
        for i in 0..self.monkeys.len() {
            self.turn(i, divide_by_three);
        }
    }

    fn monkey_business(&self) -> u64 {
        let mut monkey_inspections = self
            .monkeys
            .iter()
            .map(|monkey| monkey.inspections)
            .collect::<Vec<_>>();
        monkey_inspections.sort_unstable();
        monkey_inspections.iter().rev().take(2).product()
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();
    let lines: Vec<&str> = contents.lines().collect();

    let mut manager = MonkeyManager::new();
    for chunk in lines.chunks(7) {
        let monkey = Monkey::parse(chunk.join("\n"));
        manager.insert(monkey);
    }
    for _ in 0..20 {
        manager.round(true);
    }

    let part_1: u64 = manager.monkey_business();
    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 102_399);

    let mut manager = MonkeyManager::new();
    for chunk in lines.chunks(7) {
        let monkey = Monkey::parse(chunk.join("\n"));
        manager.insert(monkey);
    }
    for _ in 0..10_000 {
        manager.round(false);
    }

    let part_2: u64 = manager.monkey_business();
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 23_641_658_401);
}
