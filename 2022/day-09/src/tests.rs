mod part1 {
    use std::{collections::HashSet, vec};

    use crate::{Position, State};

    fn state_from_pos(head: [isize; 2], tail: [isize; 2]) -> State {
        State {
            head: Position(head[0], head[1]),
            knots: vec![Position(tail[0], tail[1])],
            visited: HashSet::new(),
        }
    }

    #[test]
    /// .....    .....    .....
    /// .TH.. -> .T.H. -> ..TH.
    /// .....    .....    .....
    fn align_tail_1() {
        let mut state = state_from_pos([3, 1], [1, 1]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 2);
        assert_eq!(state.knots[0].1, 1);
    }

    #[test]
    fn align_tail_2() {
        let mut state = state_from_pos([1, 3], [1, 1]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 1);
        assert_eq!(state.knots[0].1, 2);
    }

    #[test]
    /// ...    ...    ...
    /// .T.    .T.    ...
    /// .H. -> ... -> .T.
    /// ...    .H.    .H.
    /// ...    ...    ...
    fn align_tail_3() {
        let mut state = state_from_pos([1, 1], [3, 1]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 2);
        assert_eq!(state.knots[0].1, 1);
    }

    #[test]
    fn align_tail_4() {
        let mut state = state_from_pos([1, 1], [1, 3]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 1);
        assert_eq!(state.knots[0].1, 2);
    }

    #[test]
    /// .....    .....    .....
    /// .....    ..H..    ..H..
    /// ..H.. -> ..... -> ..T..
    /// .T...    .T...    .....
    /// .....    .....    .....
    fn align_tail_5() {
        let mut state = state_from_pos([2, 1], [1, 3]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 2);
        assert_eq!(state.knots[0].1, 2);
    }

    #[test]
    /// .....    .....    .....
    /// .....    .....    .....
    /// ..H.. -> ...H. -> ..TH.
    /// .T...    .T...    .....
    /// .....    .....    .....
    fn align_tail_6() {
        let mut state = state_from_pos([3, 2], [1, 3]);
        state.align_knot(0);
        assert_eq!(state.knots[0].0, 2);
        assert_eq!(state.knots[0].1, 2);
    }
}
