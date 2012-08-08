package sfml

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Context.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Types.h>
import "C"

import "errors"

type Context struct {
	Cref *C.sfContext
}

// Create a new context
// This function activates the new context.
// \return New sfContext object
// sfContext* sfContext_create(void);
func NewContext() (Context, error) {
	ctx := C.sfContext_create()
	if ctx == nil {
		return Context{nil}, errors.New("Couldn't make a new context")
	}
	return Context{ctx}, nil
}

//  Destroy a context
// \param context Context to destroy
// void sfContext_destroy(sfContext* context);
func (self Context) Destroy() {
	C.sfContext_destroy(self.Cref)
}

//  Activate or deactivate explicitely a context
//
// \param context Context object
// \param active  sfTrue to activate, sfFalse to deactivate
// void sfContext_setActive(sfContext* context, sfBool active);
func (self Context) SetActive(active bool) {
	if active {
		C.sfContext_setActive(self.Cref, C.sfBool(1))
	} else {
		C.sfContext_setActive(self.Cref, C.sfBool(0))
	}
}
