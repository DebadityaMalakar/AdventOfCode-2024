import re

with open("input.txt", "r") as file:
  input_data = file.read()

input_data = re.sub(r"don't\(\).*?(?:do\(\)|$)", "", input_data, flags=re.DOTALL)
matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", input_data)
print(sum([int(match[0])*int(match[1]) for match in matches]))