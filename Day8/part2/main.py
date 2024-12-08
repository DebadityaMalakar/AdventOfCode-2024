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
  pairs = list(itertools.combinations(positions, 2))

  for pair in pairs:
    a = pair[0]
    b = pair[1]
    slope = [b[0] - a[0], b[1] - a[1]]

    antinodes = [a, b]

    for i in itertools.count(start=1):
      antinode = [a[0] - i * slope[0], a[1] - i * slope[1]]
      if not in_bounds(antinode, width, height):
        break
      antinodes.append(antinode)

    for i in itertools.count(start=1):
      antinode = [b[0] + i * slope[0], b[1] + i * slope[1]]
      if not in_bounds(antinode, width, height):
        break
      antinodes.append(antinode)

    for antinode in antinodes:
      if antinode not in unique_antinodes:
        unique_antinodes.append(antinode)

print(len(unique_antinodes))