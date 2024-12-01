if __name__ == "__main__":
    with open("input.txt", "rt") as fp:
        left, right = zip(*[(int(line.split()[0]), int(line.split()[1])) for line in fp.readlines()])
        left, right = sorted(left), sorted(right)
        print(sum([abs(r - l) for l, r in zip(left, right)]))
        print(sum(l * sum(1 for r in right if l == r) for l in left))