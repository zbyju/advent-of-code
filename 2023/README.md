# Year 2023 - Rust

## Challenge - Beat AoC in 1s

Current running of all solutions (part 1 + part 2 of all days) is: **`1.60159ms`**

|        | Part 1 | Part 2 | Combined |
|--------|-------:|-------:|---------:|
| Day 01 |199.87µs|65.04µs |264.91µs  |
| Day 02 |86.41µs |79.83µs |166.24µs  |
| Day 03 |203.28µs|192.19µs|395.47µs  |
| Day 04 |167.86µs|181.29µs|349.15µs  |
| Day 05 |29.52µs |40.35µs |69.87µs   |
| Day 06 |2.04µs  |1.39µs  |3.43µs    |
| Day 07 |189.99µs|162.53µs|352.52µs  |
| Day 08 |        |        |          |
| Day 09 |        |        |          |
| Day 10 |        |        |          |
| Day 11 |        |        |          |
| Day 12 |        |        |          |
| Day 13 |        |        |          |
| Day 14 |        |        |          |
| Day 15 |        |        |          |
| Day 16 |        |        |          |
| Day 17 |        |        |          |
| Day 18 |        |        |          |
| Day 19 |        |        |          |
| Day 20 |        |        |          |
| Day 21 |        |        |          |
| Day 22 |        |        |          |
| Day 23 |        |        |          |
| Day 24 |        |        |          |
| Day 25 |        |        |          |

*These times were measured on my personal computer: AMD Ryzen 7 5800X; 16GB RAM*

## Usage

The project has a very strict structure. There is a folder for each `day` inside `days` that contains `solution.rs` where the solution is implemented. There are also `input.txt` (includes the input for that day) and `testX.txt` (includes the test inputs for that day) files inside `inputs` folder.

### Running the project

To run all the solutions at once:

```sh
cargo run

# Optimized:
cargo run --release
```

To run the project for day X:

```sh
cargo run dayX

# Example:
cargo run day09

# Optimized:
cargo run --release day09
```

To run the project for day X with some input:

```sh
cargo run dayX input.txt

# Example
cargo run day09 test1.txt
```
