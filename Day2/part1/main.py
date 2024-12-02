def parse_input(input_content):
    reports = []
    for line in input_content.splitlines():
        levels = [int(x) for x in line.split()]
        reports.append(levels)
    return reports

def check_falling(l, r):
    return l > r and 1 <= (l - r) <= 3

def check_rising(l, r):
    return r > l and 1 <= (r - l) <= 3

def part1(data):
    safe_count = 0
    for report in data:
        check = check_rising
        if report[0] >= report[-1]:
            check = check_falling
        
        safe = True
        for i in range(len(report) - 1):
            if not check(report[i], report[i+1]):
                safe = False
                break
        
        if safe:
            safe_count += 1
    
    return str(safe_count)

def main():
    with open('input.txt', 'r') as file:
        raw_data = file.read()
    
    parsed_data = parse_input(raw_data)
    result_one = part1(parsed_data)
    print(f"Result 1: {result_one}")

if __name__ == "__main__":
    main()