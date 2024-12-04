def read_file():
    """Reads the input file and returns a 2D list of characters."""
    with open("input.txt", "r") as file:
        return [list(line.strip()) for line in file]

def check_bounds(rows, cols, x, y):
    """Checks if the coordinates (x, y) are within the grid boundaries."""
    return 0 <= x < rows and 0 <= y < cols

def check_xmas_pattern(top_left, top_right, bottom_left, bottom_right):
    """Checks if the given characters form a valid XMAS pattern."""
    return ((top_left == 'M' and top_right == 'M' and bottom_left == 'S' and bottom_right == 'S') or
            (top_left == 'S' and top_right == 'S' and bottom_left == 'M' and bottom_right == 'M') or
            (top_left == 'M' and top_right == 'S' and bottom_left == 'M' and bottom_right == 'S') or
            (top_left == 'S' and top_right == 'M' and bottom_left == 'S' and bottom_right == 'M'))

def search_xmas(grid):
    """Searches for custom XMAS patterns around the character 'A'."""
    rows = len(grid)
    cols = len(grid[0])
    matches = 0

    for x in range(rows):
        for y in range(cols):
            if grid[x][y] != 'A':
                continue

            top_left_x, top_left_y = x - 1, y - 1
            top_right_x, top_right_y = x - 1, y + 1
            bottom_left_x, bottom_left_y = x + 1, y - 1
            bottom_right_x, bottom_right_y = x + 1, y + 1

            if not (check_bounds(rows, cols, top_left_x, top_left_y) and
                    check_bounds(rows, cols, top_right_x, top_right_y) and
                    check_bounds(rows, cols, bottom_left_x, bottom_left_y) and
                    check_bounds(rows, cols, bottom_right_x, bottom_right_y)):
                continue

            top_left = grid[top_left_x][top_left_y]
            top_right = grid[top_right_x][top_right_y]
            bottom_left = grid[bottom_left_x][bottom_left_y]
            bottom_right = grid[bottom_right_x][bottom_right_y]

            if check_xmas_pattern(top_left, top_right, bottom_left, bottom_right):
                matches += 1

    return matches

def part2():
    grid = read_file()
    matches = search_xmas(grid)
    print("Part 2 answer:", matches)

if __name__ == "__main__":
    part2()
