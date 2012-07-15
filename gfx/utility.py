def odd(n): return n % 2 != 0

def pairify(lst):
    if odd(len(lst)):
        raise ValueError("Odd number of elements passed to pairify")

    if len(lst) == 0:
        return []

    return [(lst[0], lst[1])] + pairify(lst[2:])


def lookupErr(line):
    line3 = int(line.split("tmpfile-pass-3.ally:")[-1]) - 1
    val3 = open("./temp/tmpfile-pass-3.ally").readlines()[line3]

    line2 = int(val3.split("tmpfile-pass-2.ally:")[-1]) - 1        
    val2 = open("./temp/tmpfile-pass-2.ally").readlines()[line2]

    line1 = int(val2.split("tmpfile-pass-1.ally:")[-1]) - 1        
    val1 = open("./temp/tmpfile-pass-1.ally").readlines()[line1]
                
    userfilename = val1.split(":")[-2]
    userlinenum = int(val1.split(":")[-1])         
    userfile = open(userfilename.strip(), 'r').readlines()
    userline = userfile[userlinenum]

    return "%s:%s" % (userfilename, userlinenum)
    return (userfilename, userline, str(userlinenum))


def error(msg, line):
    fn, ln, num = lookupErr( line)        
    tmp = "\n\n%s\n\nfile:    %s\nline:    %s\n\nline number:    %s"
    raise SyntaxError(tmp % ( msg.strip(),
                              fn.strip(),
                              ln.strip(),
                              num.strip()))

def has_function(mod, s):
    return s in mod.__dict__
    


if __name__ == "__main__":
    print pairify([1, 2,3,4]) == [(1,2),(3,4)] 
    
