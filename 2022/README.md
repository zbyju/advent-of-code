# 2022 - Go

These are my Go solutions, I'm **attempting** to solve **all days** and **all parts** under **<1s** when all the times are **combined** together.

## Times

All the solutions were ran on my PC: AMD Ryzen 7 5800X; 16GB RAM

| **Day / time[s]** | **Part 1 [ms]** | **Part 2 [ms]** |
| :---------------: | --------------: | --------------: |
|      Day 01       |           0.112 |           0.085 |
|      Day 02       |           0.161 |           0.087 |
|      Day 03       |           0.232 |           0.224 |
|      Day 04       |           0.207 |           0.162 |
|      Day 05       |           0.204 |           0.083 |
|      Day 06       |           0.101 |           0.054 |
|      Day 07       |           0.242 |           0.163 |
|      Day 08       |           3.112 |           0.444 |
|      Day 09       |           1.017 |           1.171 |
|      Day 10       |           0.094 |           1.706 |
|      Day 11       |           0.429 |          14.578 |
|      Day 12       |           3.003 |           2.340 |
|      Day 13       |           1.149 |           2.097 |
|      Day 14       |           0.919 |           8.157 |
|      Day 15       |           0.174 |           0.251 |
|      Day 16       |          11.688 |            2954 |
|      Day 17       |           4.281 |           7.749 |
|      Day 18       |           2.908 |          18.039 |
|      Day 19       |         469.723 |           11000 |
|      Day 20       |           3.760 |          43.070 |
|      Day 21       |           0.793 |           7.213 |
|      Day 22       |           0.872 |           0.591 |
|      Day 23       |          34.469 |            2869 |
|      Day 24       |          60.673 |         850.692 |
|      Day 25       |           0.115 |           0.013 |

## How to run

You can either run _all the days and parts_, _both parts for a specific date_ or a _specific part on a specific date_.

**Run all solutions:**

```bash
go run .
```

**Run both parts on day `x`**

```bash
go run . x # Where x is the day [0-25]
```

**Run specific part `y` on a specific day `x`**

```bash
go run . x y # Where x is the day [0-25]; y is the part [1-2]
```
