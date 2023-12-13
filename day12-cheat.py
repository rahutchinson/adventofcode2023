from functools import cache

@cache # too slow without the cache - stpres function values
def numlegal(s,c):

    s = s.lstrip('.') # ignore leading dots

    # ['', ()] is legal
    if s == '':
        return int(c == ()) 

    # [s, ()] is legal so long as s has no '#' (we can convert '?' to '.')
    if c == ():
        return int(s.find('#') == -1) 

    # s starts with '#' so remove the first spring
    if s[0] == '#':
        if len(s) < c[0] or '.' in s[:c[0]]:
            return 0 # impossible - not enough space for the spring
        elif len(s) == c[0]:
            return int(len(c) == 1) #single spring, right size
        elif s[c[0]] == '#':
            return 0 # springs must be separated by '.' (or '?') 
        else:
            return numlegal(s[c[0]+1:],c[1:]) # one less spring

    # numlegal springs if we convert the first '?' to '#' + '.'
    return numlegal('#'+s[1:],c) + numlegal(s[1:],c)


springs = [c.strip().split() for c in open("./inputs/day12/input.txt").readlines()]
ss = [[c[0],tuple(int(d) for d in c[1].split(','))] for c in springs]
print("Part 1 total:", sum(numlegal(s,c) for [s,c] in ss))


ss2 = [[(s[0]+'?')*4 + s[0],s[1]*5] for s in ss]
print("Part 2 total:", sum(numlegal(s,c) for [s,c] in ss2))

# thanks ai_prof
# found here https://www.reddit.com/r/adventofcode/comments/18ge41g/comment/kd0oj1t/?utm_source=share&utm_medium=web2x&context=3
