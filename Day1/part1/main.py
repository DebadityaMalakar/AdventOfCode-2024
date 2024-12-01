def abs(x):
    return -x if x < 0 else x

def main():
    input_file = "input.txt"
    
    try:
        with open(input_file, 'r') as file:
            l1 = []
            l2 = []
            
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
                l2.append(right_num)

    except Exception as e:
        print(f"Error opening file: {e}")
        return

    l1.sort()
    l2.sort()

    total = 0
    for i in range(len(l1)):
        total += abs(l1[i] - l2[i])

    print(total)

if __name__ == "__main__":
    main()
