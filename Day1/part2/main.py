def main():
    input_file = "input.txt"
    
    try:
        with open(input_file, 'r') as file:
            l1 = []
            l2 = {}

            for line in file:
                line = line.strip()
                if ' ' not in line and '\t' not in line:
                    print(f"Invalid input format in line: {line}")
                    continue
                
                space_idx = line.find(' ') if ' ' in line else line.find('\t')
                left_num, right_num = line[:space_idx].strip(), line[space_idx+1:].strip()
                
                try:
                    left_num = int(left_num)
                    right_num = int(right_num)
                except ValueError:
                    print(f"Error parsing numbers in line: {line}")
                    continue
                
                l1.append(left_num)
                if right_num in l2:
                    l2[right_num] += 1
                else:
                    l2[right_num] = 1

    except Exception as e:
        print(f"Error opening file: {e}")
        return

    total = 0
    for n1 in l1:
        total += n1 * l2.get(n1, 0)

    print(total)

if __name__ == "__main__":
    main()
