package win

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Context.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Types.h>
import "C"


type Context struct {
	Cref *C.sfContext
}

/// Create a new context
/// This function activates the new context.
/// \return New sfContext object
// sfContext* sfContext_create(void);
func NewContext() Context { 
    return Context{C.sfContext_create()}
}
            
///  Destroy a context
/// \param context Context to destroy
// void sfContext_destroy(sfContext* context);
func (self Context) Destroy() { 
    C.sfContext_destroy(self.Cref)
}
            
///  Activate or deactivate explicitely a context
///
/// \param context Context object
/// \param active  sfTrue to activate, sfFalse to deactivate
// void sfContext_setActive(sfContext* context, sfBool active);
func (self Context) SetActive(active bool) { 
	if active {
		C.sfContext_setActive(self.Cref, C.sfBool(1))
	} else {
		C.sfContext_setActive(self.Cref, C.sfBool(0))
	}
}

            
