def read_file():
    """Reads the input file and returns a 2D list of characters."""
    with open("input.txt", "r") as file:
        return [list(line.strip()) for line in file]

DIRECTIONS = [
    (-1, 0),  # up
    (1, 0),   # down
    (0, -1),  # left
    (0, 1),   # right
    (-1, -1), # up-left
    (-1, 1),  # up-right
    (1, -1),  # down-left
    (1, 1),   # down-right
]

def check_bounds(rows, cols, x, y):
    """Checks if the coordinates (x, y) are within the grid boundaries."""
    return 0 <= x < rows and 0 <= y < cols

def search_word(grid, word):
    """Searches for occurrences of the word in the grid in all 8 directions."""
    rows = len(grid)
    cols = len(grid[0])
    matches = 0

    for x in range(rows):
        for y in range(cols):
            for dx, dy in DIRECTIONS:
                match = True
                for i in range(len(word)):
                    new_x = x + i * dx
                    new_y = y + i * dy
                    if not check_bounds(rows, cols, new_x, new_y) or grid[new_x][new_y] != word[i]:
                        match = False
                        break
                if match:
                    matches += 1
    return matches

def part1():
    grid = read_file()
    matches = search_word(grid, "XMAS")
    print("Part 1 answer:", matches)

if __name__ == "__main__":
    part1()
