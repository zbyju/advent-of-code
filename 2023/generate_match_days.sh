#!/bin/bash

# Loop from 0 to 25
for i in {0..25}
do
    # Format day number with leading zero if needed
    day=$(printf "day%02d" $i)
    # Generate and print the formatted string
    echo "\"$day\" => days::$day::solution::${day^}.run(test_case),"
done


