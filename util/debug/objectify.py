from pyparsing import *
import sys

''' 
func Obj_Font_FontDestroy(arg1 SfFont) {
func Obj_Font_FontGetCharacterSize(arg1 SfFont) uint {
'''
for line in sys.stdin.readlines():
    line = line.strip()
    
    if line[:9] != "func Obj_":
        print line
        continue

    prototype = line[9:]


    FirstArg = Literal("arg1") + Word(alphanums + "_")
    className, sig = prototype.split("_") # assumes function name has no underscores
                                          # which it shouldn't and will error if it does.

    arg = list(FirstArg.scanString(sig))
    
    if len(arg) == 0:    
        print prototype
        print "Can't objectify a function with no arguments."    
        sys.exit()
    
    else:
        # [((['arg1', 'SfFont'], {}), 35, 46)]
        name, start, stop = arg[0]
        receiver = name[1]

    sig = sig[:start] + sig[stop:]
        
    print "func (arg1 %s) " % (receiver) + sig

