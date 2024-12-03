import re

if __name__ == "__main__":
    with open("input.txt") as fp:
        sigma = 0
        do = True
        for mat in re.findall("mul\(\d+,\d+\)|don\'t\(\)|do\(\)", fp.read()):
            mat = str(mat)
            match mat:
                case "don't()": do = False
                case "do()": do = True
                case _ if do:
                    it = re.findall("\d+", mat)
                    sigma += int(it[0]) * int(it[1])
        print(sigma)
            