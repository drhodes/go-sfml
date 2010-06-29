from pyparsing import Word, Suppress, Literal, \
    commaSeparatedList, Optional, alphanums, ZeroOrMore, Group, OneOrMore
import sys
'''
func (self *IntRect) Contains(arg2 int, arg3 int) bool {
    return SfIntRect_Contains(self.Cref, arg2, arg3)
}
'''

#ex = "CSFML_API unsigned int sfShape_GetNbPoints(sfShape* Shape);"
ex = "CSFML_API sfFont* sfFont_CreateFromMemory(const char* Data, size_t SizeInBytes, unsigned int CharSize, const sfUint32* Charset);"
ex = sys.stdin.read()
#ex= "CSFML_API void sfFont_Destroy(sfFont* Font);"

ex = ex.strip()
if not ex.startswith("CSFML"):
    if ex.startswith("////////////////"):
        sys.exit()
    if ex == "///":
        sys.exit()
    print ex
    sys.exit()

ex = ex.replace("CSFML_API ", "")
print "//", ex

#func Class_Name(Arglist) ReturnValue
LPAREN, RPAREN = Suppress("("), Suppress(")")
COMMA = Suppress(",")

Identifier = Word(alphanums + "*_")
IdentNo_ =  Word(alphanums + "*")
Params = Group( OneOrMore( Identifier ) ) # last one is the ParamName
FuncClass = Suppress(Literal("sf")) + IdentNo_
FuncName = Identifier
Args = Group( LPAREN + ZeroOrMore( Params + COMMA ) + Optional( Params ) + RPAREN )
FuncGroup = FuncClass + Suppress("_") + FuncName
ReturnType = Group(Optional(Identifier + Literal(" ")) + Identifier)
Prototype =  ReturnType + FuncGroup + Args + Suppress(Literal(";"))

result = list( Prototype.parseString(ex) )
#print result

return_type = result[0]
if len(return_type) == 2:
    if return_type[0] == "unsigned":
        return_type = "u" + return_type[1]
    if return_type[0] == "const":
        return_type = return_type[1:]

#print type(return_type)
#print return_type
return_type = str(" ".join(return_type.asList()))


class_name = result[1]
params = result[3]
func_name =  result[2]

#print return_type
#print func_name


'''
if params[0][1] == class_name:
    params = params[1:]
'''

params = [p.asList() for p in params]
newParams = []
for p in params:
    temp = p
    if len(p) == 3:
        if p[0] == "const":
            temp = p[1:]
        if p[0] == "unsigned":
            temp = ['u'+p[1], p[2]]
    temp.reverse()
    name = temp[0][0].lower() + temp[0][1:]
    temp = [name] + temp[1:]
    newParams.append(temp)

params = newParams            
        
#print params


class Wrapper(object):
    def __init__(self, class_name, func_name, params, return_type):
        self.class_name = class_name
        self.func_name = func_name
        self.params = params # list
        self.return_type = return_type

        self.crefify() # tack on a .Cref to params that need it.
        self.emit_sig()
        self.emit_body()
        
    def emit_sig(self):
        params = ', '.join([' '.join(x) for x in self.params[1:]])
        params = params.replace("Sf", "")
        params = params.replace(".Cref", "")
        
        sig = "func (self *%s) %s(%s) %s {"
        if self.return_type.startswith("Sf"):
            rtype = self.return_type[2:]
        else:
            rtype = self.return_type

        print sig % ( self.class_name, 
                      self.func_name, 
                      params,
                      rtype )

    def emit_body(self):
        if self.return_type == "":
            self.emit_body_no_return()
        else:
            self.emit_body_with_return()

    def crefify(self):
        for p in self.params:
            if p[1].startswith("sf"):
                p[1] = p[1][2:]
                p[0] = p[0] + ".Cref"

    def emit_body_with_return(self):
        params = ', '.join([x[0] for x in self.params[1:]])

        if self.return_type.startswith("Sf"):
            cCall = "C.sf%s_%s(self.Cref, %s)" % (self.class_name, self.func_name, params)
            output = "    return %s{ %s }\n}" % (self.return_type[2:], cCall)
            print output.replace(", )", ")")
        else:
            output = "    return C.sf%s_%s(self.Cref, %s)\n}\n\n" % (self.class_name, 
                                                               self.func_name, 
                                                               params)
            print output.replace(", )", ")")

    def emit_body_no_return(self):
        params = ', '.join([x[0] for x in self.params[1:]])
        print "    C.sf%s_%s(self.Cref, %s)\n}\n\n" % (self.class_name, self.func_name, params)

Wrapper(class_name, func_name, params, return_type)



