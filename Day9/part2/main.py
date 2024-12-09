class Day9_2:
    def __init__(self):
        self.label = "day9:2"
        self.path = "test"
        self.files = {}
        self.file_blocks = {}
        self.free_blocks = {}
        self.max_block = 0

    def parse_blocks(self, data):
        is_file = True

        for char in data:
            n = int(char)

            if is_file:
                self.files[len(self.files)] = (self.max_block, n)
                self.file_blocks[self.max_block] = (len(self.file_blocks), n)
            elif n != 0:
                self.free_blocks[self.max_block] = n

            self.max_block += n
            is_file = not is_file

    def move_file(self, file_id, start, size):
        for i in range(start):
            if i in self.free_blocks and self.free_blocks[i] >= size:
                self.free_blocks[start] = size
                del self.free_blocks[i]

                if self.free_blocks[i] > size:
                    self.free_blocks[i + size] = self.free_blocks[i] - size

                self.files[file_id] = (i, size)
                del self.file_blocks[start]
                self.file_blocks[i] = (file_id, size)
                break

    def move_files(self):
        sorted_files = sorted(self.files.items(), key=lambda x: x[1][0], reverse=True)

        for file_id, (start, size) in sorted_files:
            self.move_file(file_id, start, size)

    def checksum(self):
        result = 0
        i = 0

        while i < self.max_block:
            if i in self.file_blocks:
                file_id, size = self.file_blocks[i]
                for _ in range(size):
                    result += i * file_id
                    i += 1
            else:
                i += 1

        return result

    def pipeline(self):
        try:
            with open("input.txt", "r") as file:
                data = file.read().strip()
            self.parse_blocks(data)
            self.move_files()
            return self.checksum()
        except Exception as e:
            print("Error:", e)
            return 0


if __name__ == "__main__":
    day2 = Day9_2()
    result2 = day2.pipeline()
    print("Checksum (Part 2):", result2)
