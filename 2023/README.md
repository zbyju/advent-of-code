# Year 2023 - Rust

## Usage

The project has a very strict structure. There is a folder for each `day` inside `days` that contains `solution.rs` where the solution is implemented. There are also `input.txt` (includes the input for that day) and `testX.txt` (includes the test inputs for that day) files inside `inputs` folder.

### Running the project

To run the project for day X:

```sh
cargo run dayX

# Example:
cargo run day09
```

To run the project for day X with some input:

```sh
cargo run dayX input.txt

# Example
cargo run day09 test1.txt
```
