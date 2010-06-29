package win

// #include <SFML/Window/Context.h>
import "C"

import(
	"runtime"
	//"unsafe"
)

// utility functions must be from in this package to make happy cgo.
func SfBool2GoBool(cVal C.sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) C.sfBool {
	if goVal {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}

/*
func NewColor(val C.sfColor) Color {
	return Color{ val }
}

type Color struct { 
	Cref C.sfColor 
}
*/

//-------------------------------------
func NewContext(val *C.sfContext) Context {
    tmp := Context{ val }
    runtime.SetFinalizer(&tmp, (*Context).Destroy)
    return tmp
}

type Context struct {
	Cref *C.sfContext
}




// _Context_
// -------------------------------------------------------------------------------
/// Construct a new context
/// return New context
// sfContext* sfContext_Create();
func CreateContext() Context {
    return NewContext( C.sfContext_Create() )
}
/// Destroy an existing context
/// param Context : Context to destroy
// void sfContext_Destroy(sfContext* Context);
func (self *Context) Destroy() {
    C.sfContext_Destroy(self.Cref)
	self.Cref = nil
}

/// Activate or deactivate a context
/// param Context : Context to activate or deactivate
/// param Active : sfTrue to activate, sfFalse to deactivate
// void sfContext_SetActive(sfContext* Context, sfBool Active);
func (self *Context) SetActive(active bool) {
    C.sfContext_SetActive(self.Cref, GoBool2SfBool(active))
}

