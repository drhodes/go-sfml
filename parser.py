#! /usr/bin/env python
import sys
import re, fileinput
import pyPEG
from pyPEG import parseLine
from pyPEG import keyword, _and, _not, ignore

# pyPEG:
#                    0: following element is optional
#                   -1: following element can be omitted or repeated endless
#                   -2: following element is required and can be repeated endless


def comment():          return [re.compile(r"//.*"), re.compile("/\*.*?\*/", re.S)]
def const():            return (re.compile(r"const"))
def unsigned():         return (re.compile(r"unsigned"))
def void():             return (re.compile(r"void"))
def symbol():           return re.compile(r"[A-z]+[A-z|0-9|_|\*]*")
def idgroup():          return 0, [unsigned, const], 0, symbol, 0, symbol
def parameterlist():    return "(", 0, idgroup, -1, (",", idgroup),  ")"
def proto():            return 0, [unsigned, const], symbol, symbol, parameterlist, ";"
def some():             return -1, proto


pyPEG.print_trace = False

def go(t, rule):    
    return parseLine( t,
                      rule,
                      resultSoFar = [], 
                      skipWS = True,
                      skipComments = None,
                      packrat = False)

#ID = "sfTexture*"
#print go(ID, symbol)
#group = "const void* data"
#print go(group, idgroup)
#plist = "(const void* data, size_t sizeInBytes, const sfIntRect* area)"
#print go(plist, parameterlist)
#txt = "sfTexture* sfTexture_createFromMemory(const void* data, size_t sizeInBytes, const sfIntRect* area);"
#print go(txt, proto)
#txt2 = "sfGlyph sfFont_getGlyph(sfFont* font, sfUint32 codePoint, unsigned int characterSize, sfBool bold);"
#print go(txt2, proto)
#t3 = "const sfFont* sfFont_getDefaultFont(void);"
#go(t3, proto)
t4 = "sfColor sfColor_fromRGB(sfUint8 red, sfUint8 green, sfUint8 blue);"
print go(t4, proto)

# def walk(node):
#     if node.__class__.__name__ == "Symbol":
#         print node.__class__.__name__

#     for n in node:
#         if type(n) == type([]):
#             for j in n:
#                 walk(j)
#         else:
#             print n









