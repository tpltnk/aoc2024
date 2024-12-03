if __name__ == "__main__":
    with open("input.txt") as fp:
        safe = 0
        for record in fp.readlines():
            levels = list(map(lambda x: int(x), record.split(" ")))
            ok = []
            for j in range(len(levels)):
                levelsafe = True
                rem = levels[:j] + levels[j+1:]
                diffs = [rem[i+1]-rem[i] for i in range(len(rem)-1)]
                if any(abs(d) < 1 or abs(d) > 3 for d in diffs):
                    levelsafe = False
                inc = all(d < 0 for d in diffs)
                dec = all(d > 0 for d in diffs)
                if not inc and not dec:
                    levelsafe = False
                if inc and dec:
                    levelsafe = False
                ok.append(levelsafe)
            if any(ok):
                safe += 1
        print(safe)
            
