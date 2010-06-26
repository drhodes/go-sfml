package gfx

import "fmt"

func Debug(x interface{}){	fmt.Printf("%#v\n", x); }
func DebugT(x interface{}){ fmt.Printf("%#T\n", x); }
func DebugN(x interface{}){ fmt.Printf("%#v", x); }
