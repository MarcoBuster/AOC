mod part1 {
    use crate::Monkey;

    #[test]
    fn monkey_parse_1() {
        let input = r#"Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3"#;
        let monkey = Monkey::parse(input.into());

        assert_eq!(monkey.items, vec![79, 98]);
        assert_eq!(monkey.if_true_dest, 2);
        assert_eq!(monkey.if_false_dest, 3);
    }
}
