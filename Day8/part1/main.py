import itertools

def in_bounds(pos, w, h):
  i = pos[0]
  j = pos[1]
  return i in range(h) and j in range(w)


with open("./input.txt", "r") as file:
  input_data = [list(line.strip()) for line in file.readlines()]


width = len(input_data[0])
height = len(input_data)

antennas = {}
for i in range(height):
  for j in range(width):
    freq, pos = input_data[i][j], [i, j]
    if freq != ".":
      if freq not in antennas.keys():
        antennas[freq] = [pos]
      else:
        antennas[freq].append(pos)

unique_antinodes = []
for frequency in antennas.keys():
  positions = antennas[frequency]
  pairs = list(itertools.combinations(positions, 2))  # get unique pairs of positions

  for pair in pairs:
    a = pair[0]
    b = pair[1]
    slope = [b[0] - a[0], b[1] - a[1]]  # slope from a to b

    # one antinode is a minus the slope, other is b plus the slope
    antinodes = [[a[0] - slope[0], a[1] - slope[1]], [b[0] + slope[0], b[1] + slope[1]]]
    for antinode in antinodes:
      if antinode not in unique_antinodes and in_bounds(antinode, width, height):
        unique_antinodes.append(antinode)

print(len(unique_antinodes))