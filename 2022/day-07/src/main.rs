// Inspired by https://github.com/nappa85/aoc_2022/blob/283907a4d4bf4a55f0a304a7dc4027390edbf34d/day7/src/main.rs

use std::fs;

enum Line<'a> {
    Cd(&'a str),
    CdDotDot,
    Ls,
    Dir(&'a str),
    File(u64, &'a str),
}

struct Folder {
    parent: Option<usize>,
    files: u64,
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();
    let mut folders: Vec<Folder> = vec![];
    let mut current = None;

    contents
        .lines()
        .map(|line| {
            let mut words = line.split(' ');
            match words.next().unwrap() {
                "$" => match words.next().unwrap() {
                    "cd" => match words.next().unwrap() {
                        ".." => Line::CdDotDot,
                        s => Line::Cd(s),
                    },
                    "ls" => Line::Ls,
                    _ => panic!("Command not found"),
                },
                "dir" => Line::Dir(words.next().unwrap()),
                n => Line::File(n.parse().unwrap(), words.next().unwrap()),
            }
        })
        .for_each(|l| match l {
            Line::Cd(_name) => {
                folders.push(Folder {
                    parent: current,
                    files: 0,
                });
                current = Some(folders.len() - 1);
            }
            Line::CdDotDot => {
                current = folders[current.unwrap()].parent;
            }
            Line::Ls => {}
            Line::Dir(_name) => {}
            Line::File(size, _name) => {
                let mut parent = current;
                while let Some(id) = parent {
                    folders[id].files += size;
                    parent = folders[id].parent;
                }
            }
        });

    let part_1 = folders
        .iter()
        .skip(1) // skip root
        .filter_map(|f| (f.files < 100_000).then_some(f.files))
        .sum::<u64>();
    println!("Part 1: {part_1}");
    assert_eq!(part_1, 1490523);

    let free_space = 70_000_000 - folders[0].files;
    let space_to_free = 30_000_000 - free_space;
    let mut part_2 = folders
        .iter()
        .filter_map(|f| (f.files >= space_to_free).then_some(f.files))
        .collect::<Vec<u64>>();
    part_2.sort_unstable();
    let part_2 = part_2[0];
    println!("Part 2: {}", part_2);
    assert_eq!(part_2, 12390492);
}
