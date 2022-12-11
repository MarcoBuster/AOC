use std::fs;

#[derive(Debug, Clone)]
enum Instruction {
    Nop,
    Add(isize),
    Bubble,
}

fn parse_instruction(value: &str) -> Instruction {
    if value == "noop" {
        Instruction::Nop
    } else if value.starts_with("addx") {
        let (_, operand) = value.split_once(' ').unwrap();
        Instruction::Add(operand.parse().unwrap())
    } else {
        panic!("Invalid instruction");
    }
}

struct VM {
    register: isize,
}

impl VM {
    fn new() -> Self {
        VM { register: 1 }
    }

    fn single_cycle(&mut self, instruction: &Instruction) {
        if let Instruction::Add(operand) = instruction {
            self.register += operand;
        }
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();
    let instructions: Vec<Instruction> = contents.lines().map(parse_instruction).collect();

    let mut spaced_instructions = vec![];
    instructions
        .iter()
        .for_each(|instruction| match instruction {
            Instruction::Nop => spaced_instructions.push(Instruction::Bubble),
            Instruction::Add(operand) => {
                spaced_instructions.push(Instruction::Nop);
                spaced_instructions.push(Instruction::Add(*operand));
            }
            Instruction::Bubble => unreachable!(),
        });
    std::mem::drop(instructions);

    const INTERESTING_CYCLES: [isize; 6] = [20, 60, 100, 140, 180, 220];
    
    let mut vm = VM::new();
    let mut part_1: isize = 0;
    for (i, instruction) in spaced_instructions.iter().enumerate() {
        let i = i as isize;
        if INTERESTING_CYCLES.contains(&(i + 1)) {
            part_1 += (i + 1) * vm.register;
        }
        vm.single_cycle(instruction);
    }
    println!("Part 1: {}", part_1);

    const WIDTH: isize = 40;
    const SPRITE_SIZE: isize = 3;

    println!("Part 2:");
    let mut vm = VM::new();
    for (i, instruction) in spaced_instructions.iter().enumerate() {
        let i = i as isize;
        let position = (i % WIDTH) as isize + 1;

        if position >= vm.register && position < vm.register + SPRITE_SIZE {
            print!("#");
        } else {
            print!(".");
        }

        if position == WIDTH {
            println!();
        }
        vm.single_cycle(instruction);
    }
}
