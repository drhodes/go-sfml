// Move a rectangle by the given offset
// param Rect : Rectangle to move
// param OffsetX : Horizontal offset
// param OffsetY : Vertical offset
// void sfFloatRect_Offset(sfFloatRect* Rect, float OffsetX, float OffsetY);
func (self *FloatRect) Offset(offsetX float, offsetY float) {
    C.sfFloatRect_Offset(self.Cref, offsetX, offsetY) }
}

// void sfIntRect_Offset(sfIntRect* Rect, int OffsetX, int OffsetY);
func (self *IntRect) Offset(offsetX int, offsetY int) {
    C.sfIntRect_Offset(self.Cref, offsetX, offsetY) 
}

// Check if a point is inside a rectangle's area
// param Rect : Rectangle to test
// param X : X coordinate of the point to test
// param Y : Y coordinate of the point to test
// return sfTrue if the point is inside
// sfBool sfFloatRect_Contains(sfFloatRect* Rect, float X, float Y);
func (self *FloatRect) Contains(x float, y float) bool {
    return SfBool2GoBool( C.sfFloatRect_Contains(self.Cref, x, y) )
}

// sfBool sfIntRect_Contains(sfIntRect* Rect, int X, int Y);
func (self *IntRect) Contains(x int, y int) sfBool {
    return SfBool2GoBool( C.sfIntRect_Contains(self.Cref, x, y) )
}

// Check intersection between two rectangles
// param Rect1 : First rectangle to test
// param Rect2 : Second rectangle to test
// param OverlappingRect : Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfFloatRect_Intersects(sfFloatRect* Rect1, sfFloatRect* Rect2, sfFloatRect* OverlappingRect);
func (self *FloatRect) Intersects(rect2 sfFloatRect*, overlappingRect sfFloatRect*) sfBool {
	return SfBool2GoBool( C.sfFloatRect_Intersects(self.Cref, rect2, overlappingRect) )
}

// sfBool sfIntRect_Intersects(sfIntRect* Rect1, sfIntRect* Rect2, sfIntRect* OverlappingRect);
func (self *IntRect) Intersects(rect2 sfIntRect*, overlappingRect sfIntRect*) sfBool {
    return SfBool2GoBool( C.sfIntRect_Intersects(self.Cref, rect2, overlappingRect) )
}
