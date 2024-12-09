class Day9_1:
    def __init__(self):
        self.label = "day9:1"
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

    def move_blocks(self, file_id, start, n):
        i = 0
        remaining = n

        while i < start:
            if i in self.free_blocks:
                free_space = self.free_blocks[i]
                size = min(free_space, remaining)
                self.file_blocks[i] = (file_id, size)
                del self.free_blocks[i]

                if free_space > size:
                    self.free_blocks[i + size] = free_space - size

                if size == remaining:
                    del self.file_blocks[start]
                    break

                remaining -= size
                self.file_blocks[start] = (file_id, remaining)
                i += size
            else:
                i += 1

    def move_files(self):
        for i in range(self.max_block - 1, -1, -1):
            if i in self.file_blocks:
                file_id, size = self.file_blocks[i]
                self.move_blocks(file_id, i, size)

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
    day1 = Day9_1()
    result1 = day1.pipeline()
    print("Checksum (Part 1):", result1)
