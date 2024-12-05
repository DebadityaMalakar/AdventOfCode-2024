import re


# Checks if the update is valid according to the ordering rules
def is_valid_update(update, ordering_rules):
    for index, page in enumerate(update):
        previous_pages = update[:index]
        should_be_later_pages = ordering_rules.get(page, [])
        for p in previous_pages:
            if p in should_be_later_pages:
                return False
    return True


# Corrects an invalid update based on ordering rules
def correct_update(update, ordering_rules):
    corrected_update = update[:]

    while not is_valid_update(corrected_update, ordering_rules):
        for index, page in enumerate(corrected_update):
            previous_pages = corrected_update[:index]
            should_be_later_pages = ordering_rules.get(page, [])
            for p in previous_pages:
                if p in should_be_later_pages:
                    first_page_to_swap_index = corrected_update.index(p)
                    corrected_update.pop(first_page_to_swap_index)
                    corrected_update.insert(index, p)
                    break
    return corrected_update


# Main function to solve the problem
def solve(input_file):
    # Read the file
    try:
        with open(input_file, 'r') as file:
            content = file.read()
    except FileNotFoundError:
        print(f"Error: File '{input_file}' not found.")
        return ""

    # Extract ordering rules
    ordering_rule_regex = re.compile(r"(\d+)\|(\d+)")
    matches = ordering_rule_regex.findall(content)
    ordering_rules = {}
    for match in matches:
        if match[0] not in ordering_rules:
            ordering_rules[match[0]] = []
        ordering_rules[match[0]].append(match[1])

    # Extract pages to check
    page_check_regex = re.compile(r"\d+(?:,\d+)+")
    page_matches = page_check_regex.findall(content)
    pages_to_check = [match.split(",") for match in page_matches]

    # Process all invalid updates sequentially and calculate the sum
    total_sum = 0
    for update in pages_to_check:
        if not is_valid_update(update, ordering_rules):
            corrected_update = correct_update(update, ordering_rules)
            mid_index = len(corrected_update) // 2
            value = int(corrected_update[mid_index])
            total_sum += value

    return str(total_sum)


if __name__ == "__main__":
    # Call the solve function with the input file
    result = solve("input.txt")
    print(result)
