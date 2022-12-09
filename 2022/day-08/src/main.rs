use std::fs;
use std::io::Result;

#[derive(Debug)]
struct Tree {
    height: isize,
    visible: bool,
    scenic_score: usize,
}

impl Tree {
    fn new(height: isize) -> Self {
        Tree {
            height,
            visible: false,
            scenic_score: 0,
        }
    }
}

fn mark_visible_trees(sequence: &mut [&mut Tree]) {
    assert!(sequence.len() > 1);

    for i in 0..sequence.len() {
        if sequence[i].height
            > sequence[0..i]
                .iter()
                .map(|tree| tree.height)
                .max()
                .unwrap_or(-1)
            || sequence[i].height
                > sequence[i + 1..sequence.len()]
                    .iter()
                    .map(|tree| tree.height)
                    .max()
                    .unwrap_or(-1)
        {
            sequence[i].visible = true;
        }
    }
}

fn viewing_distance(sequence: &[&Tree]) -> usize {
    assert!(sequence.len() > 1);

    let ref_tree = sequence[0];
    let mut count = sequence[1..]
        .iter()
        .take_while(|tree| tree.height < ref_tree.height)
        .count();

    // If not on an edge, the blocking tree still counts
    if count < sequence[1..].len() {
        count += 1;
    }

    count
}

fn main() -> Result<()> {
    let contents = fs::read_to_string("input.txt")?;

    // Parse input
    let mut matrix: Vec<Vec<Tree>> = contents
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| Tree::new(c.to_string().parse::<isize>().unwrap()))
                .collect()
        })
        .collect();

    let n = matrix.len();

    // Part 1: mark visible trees
    for i in 0..n {
        // Select i-row
        mark_visible_trees(&mut matrix[i].iter_mut().collect::<Vec<_>>()[..]);

        // Select i-column
        mark_visible_trees(
            &mut matrix
                .iter_mut()
                .map(|row| row.get_mut(i).unwrap())
                .collect::<Vec<_>>()[..],
        );
    }

    // Count visible trees
    let part_1: usize = matrix
        .iter()
        .map(|row| {
            row.iter()
                .map(|tree| if tree.visible { 1 } else { 0 })
                .sum::<usize>()
        })
        .sum();
    println!("Part 1: {}", part_1);
    assert_eq!(part_1, 1827);

    // Part 2: calculate scenic score by viewing distance
    for i in 1..n - 1 {
        for j in 1..n - 1 {
            let right = viewing_distance(&matrix[i][j..n].iter().collect::<Vec<&Tree>>()[..]);
            let left = viewing_distance(&matrix[i][0..=j].iter().rev().collect::<Vec<&Tree>>()[..]);
            let down = viewing_distance(
                &matrix[i..n]
                    .iter()
                    .map(|row| &row[j])
                    .collect::<Vec<&Tree>>()[..],
            );
            let up = viewing_distance(
                &matrix[0..=i]
                    .iter()
                    .map(|row| &row[j])
                    .rev()
                    .collect::<Vec<&Tree>>()[..],
            );
            matrix[i][j].scenic_score = right * left * down * up;
        }
    }

    // Find maximum scenic score
    let part_2: usize = matrix
        .iter()
        .map(|row| row.iter().map(|tree| tree.scenic_score).max().unwrap())
        .max()
        .unwrap();
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 335_580);

    Ok(())
}
