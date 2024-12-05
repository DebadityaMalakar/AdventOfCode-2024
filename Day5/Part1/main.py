def main():
    # Read and split input
    with open("input.txt") as file:
        content = [line.strip() for line in file]

    # Find split index
    split_index = content.index("")

    rules = parse_rules(content[:split_index])
    updates = parse_updates(content[split_index + 1:])

    n_part1 = 0

    # Check if all rules are correct for each update
    for values in updates:
        relevant_rules = filter_relevant_rules(rules, values)
        if all_correct(values, relevant_rules):
            mid_index = len(values) // 2
            val = atoi(values[mid_index])
            n_part1 += val

    print(n_part1)


def parse_rules(lines):
    rules = []
    for line in lines:
        parts = line.split("|")
        rules.append((parts[0], parts[1]))
    return rules


def parse_updates(lines):
    updates = []
    for line in lines:
        updates.append(line.split(","))
    return updates


def filter_relevant_rules(rules, values):
    relevant_rules = []
    for rule in rules:
        if contains(values, rule[0]) and contains(values, rule[1]):
            relevant_rules.append(rule)
    return relevant_rules


def all_correct(values, rules):
    for rule in rules:
        if index_of(values, rule[0]) >= index_of(values, rule[1]):
            return False
    return True


def contains(lst, item):
    return item in lst


def index_of(lst, item):
    try:
        return lst.index(item)
    except ValueError:
        return -1


def atoi(s):
    try:
        return int(s)
    except ValueError:
        return 0


if __name__ == "__main__":
    main()
