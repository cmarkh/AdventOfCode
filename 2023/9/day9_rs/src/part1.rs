fn parse_input(input: &str) {}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    fn get_input(file_name: &str) {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let _ = get_input(file_name);
    }
}
