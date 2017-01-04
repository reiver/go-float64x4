package float64x4;



// WWWW returns a new float64x4 vector using the WWWW lanes of the original float64x4 vector.
func (receiver T) WWWW() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[w],
        receiver[w],
    }
}


// WWWX returns a new float64x4 vector using the WWWX lanes of the original float64x4 vector.
func (receiver T) WWWX() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[w],
        receiver[x],
    }
}


// WWWY returns a new float64x4 vector using the WWWY lanes of the original float64x4 vector.
func (receiver T) WWWY() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[w],
        receiver[y],
    }
}


// WWWZ returns a new float64x4 vector using the WWWZ lanes of the original float64x4 vector.
func (receiver T) WWWZ() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[w],
        receiver[z],
    }
}


// WWXW returns a new float64x4 vector using the WWXW lanes of the original float64x4 vector.
func (receiver T) WWXW() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[x],
        receiver[w],
    }
}


// WWXX returns a new float64x4 vector using the WWXX lanes of the original float64x4 vector.
func (receiver T) WWXX() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[x],
        receiver[x],
    }
}


// WWXY returns a new float64x4 vector using the WWXY lanes of the original float64x4 vector.
func (receiver T) WWXY() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[x],
        receiver[y],
    }
}


// WWXZ returns a new float64x4 vector using the WWXZ lanes of the original float64x4 vector.
func (receiver T) WWXZ() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[x],
        receiver[z],
    }
}


// WWYW returns a new float64x4 vector using the WWYW lanes of the original float64x4 vector.
func (receiver T) WWYW() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[y],
        receiver[w],
    }
}


// WWYX returns a new float64x4 vector using the WWYX lanes of the original float64x4 vector.
func (receiver T) WWYX() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[y],
        receiver[x],
    }
}


// WWYY returns a new float64x4 vector using the WWYY lanes of the original float64x4 vector.
func (receiver T) WWYY() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[y],
        receiver[y],
    }
}


// WWYZ returns a new float64x4 vector using the WWYZ lanes of the original float64x4 vector.
func (receiver T) WWYZ() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[y],
        receiver[z],
    }
}


// WWZW returns a new float64x4 vector using the WWZW lanes of the original float64x4 vector.
func (receiver T) WWZW() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[z],
        receiver[w],
    }
}


// WWZX returns a new float64x4 vector using the WWZX lanes of the original float64x4 vector.
func (receiver T) WWZX() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[z],
        receiver[x],
    }
}


// WWZY returns a new float64x4 vector using the WWZY lanes of the original float64x4 vector.
func (receiver T) WWZY() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[z],
        receiver[y],
    }
}


// WWZZ returns a new float64x4 vector using the WWZZ lanes of the original float64x4 vector.
func (receiver T) WWZZ() T {
    return T{
        receiver[w],
        receiver[w],
        receiver[z],
        receiver[z],
    }
}


// WXWW returns a new float64x4 vector using the WXWW lanes of the original float64x4 vector.
func (receiver T) WXWW() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[w],
        receiver[w],
    }
}


// WXWX returns a new float64x4 vector using the WXWX lanes of the original float64x4 vector.
func (receiver T) WXWX() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[w],
        receiver[x],
    }
}


// WXWY returns a new float64x4 vector using the WXWY lanes of the original float64x4 vector.
func (receiver T) WXWY() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[w],
        receiver[y],
    }
}


// WXWZ returns a new float64x4 vector using the WXWZ lanes of the original float64x4 vector.
func (receiver T) WXWZ() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[w],
        receiver[z],
    }
}


// WXXW returns a new float64x4 vector using the WXXW lanes of the original float64x4 vector.
func (receiver T) WXXW() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[x],
        receiver[w],
    }
}


// WXXX returns a new float64x4 vector using the WXXX lanes of the original float64x4 vector.
func (receiver T) WXXX() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[x],
        receiver[x],
    }
}


// WXXY returns a new float64x4 vector using the WXXY lanes of the original float64x4 vector.
func (receiver T) WXXY() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[x],
        receiver[y],
    }
}


// WXXZ returns a new float64x4 vector using the WXXZ lanes of the original float64x4 vector.
func (receiver T) WXXZ() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[x],
        receiver[z],
    }
}


// WXYW returns a new float64x4 vector using the WXYW lanes of the original float64x4 vector.
func (receiver T) WXYW() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[y],
        receiver[w],
    }
}


// WXYX returns a new float64x4 vector using the WXYX lanes of the original float64x4 vector.
func (receiver T) WXYX() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[y],
        receiver[x],
    }
}


// WXYY returns a new float64x4 vector using the WXYY lanes of the original float64x4 vector.
func (receiver T) WXYY() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[y],
        receiver[y],
    }
}


// WXYZ returns a new float64x4 vector using the WXYZ lanes of the original float64x4 vector.
func (receiver T) WXYZ() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[y],
        receiver[z],
    }
}


// WXZW returns a new float64x4 vector using the WXZW lanes of the original float64x4 vector.
func (receiver T) WXZW() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[z],
        receiver[w],
    }
}


// WXZX returns a new float64x4 vector using the WXZX lanes of the original float64x4 vector.
func (receiver T) WXZX() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[z],
        receiver[x],
    }
}


// WXZY returns a new float64x4 vector using the WXZY lanes of the original float64x4 vector.
func (receiver T) WXZY() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[z],
        receiver[y],
    }
}


// WXZZ returns a new float64x4 vector using the WXZZ lanes of the original float64x4 vector.
func (receiver T) WXZZ() T {
    return T{
        receiver[w],
        receiver[x],
        receiver[z],
        receiver[z],
    }
}


// WYWW returns a new float64x4 vector using the WYWW lanes of the original float64x4 vector.
func (receiver T) WYWW() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[w],
        receiver[w],
    }
}


// WYWX returns a new float64x4 vector using the WYWX lanes of the original float64x4 vector.
func (receiver T) WYWX() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[w],
        receiver[x],
    }
}


// WYWY returns a new float64x4 vector using the WYWY lanes of the original float64x4 vector.
func (receiver T) WYWY() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[w],
        receiver[y],
    }
}


// WYWZ returns a new float64x4 vector using the WYWZ lanes of the original float64x4 vector.
func (receiver T) WYWZ() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[w],
        receiver[z],
    }
}


// WYXW returns a new float64x4 vector using the WYXW lanes of the original float64x4 vector.
func (receiver T) WYXW() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[x],
        receiver[w],
    }
}


// WYXX returns a new float64x4 vector using the WYXX lanes of the original float64x4 vector.
func (receiver T) WYXX() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[x],
        receiver[x],
    }
}


// WYXY returns a new float64x4 vector using the WYXY lanes of the original float64x4 vector.
func (receiver T) WYXY() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[x],
        receiver[y],
    }
}


// WYXZ returns a new float64x4 vector using the WYXZ lanes of the original float64x4 vector.
func (receiver T) WYXZ() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[x],
        receiver[z],
    }
}


// WYYW returns a new float64x4 vector using the WYYW lanes of the original float64x4 vector.
func (receiver T) WYYW() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[y],
        receiver[w],
    }
}


// WYYX returns a new float64x4 vector using the WYYX lanes of the original float64x4 vector.
func (receiver T) WYYX() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[y],
        receiver[x],
    }
}


// WYYY returns a new float64x4 vector using the WYYY lanes of the original float64x4 vector.
func (receiver T) WYYY() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[y],
        receiver[y],
    }
}


// WYYZ returns a new float64x4 vector using the WYYZ lanes of the original float64x4 vector.
func (receiver T) WYYZ() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[y],
        receiver[z],
    }
}


// WYZW returns a new float64x4 vector using the WYZW lanes of the original float64x4 vector.
func (receiver T) WYZW() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[z],
        receiver[w],
    }
}


// WYZX returns a new float64x4 vector using the WYZX lanes of the original float64x4 vector.
func (receiver T) WYZX() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[z],
        receiver[x],
    }
}


// WYZY returns a new float64x4 vector using the WYZY lanes of the original float64x4 vector.
func (receiver T) WYZY() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[z],
        receiver[y],
    }
}


// WYZZ returns a new float64x4 vector using the WYZZ lanes of the original float64x4 vector.
func (receiver T) WYZZ() T {
    return T{
        receiver[w],
        receiver[y],
        receiver[z],
        receiver[z],
    }
}


// WZWW returns a new float64x4 vector using the WZWW lanes of the original float64x4 vector.
func (receiver T) WZWW() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[w],
        receiver[w],
    }
}


// WZWX returns a new float64x4 vector using the WZWX lanes of the original float64x4 vector.
func (receiver T) WZWX() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[w],
        receiver[x],
    }
}


// WZWY returns a new float64x4 vector using the WZWY lanes of the original float64x4 vector.
func (receiver T) WZWY() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[w],
        receiver[y],
    }
}


// WZWZ returns a new float64x4 vector using the WZWZ lanes of the original float64x4 vector.
func (receiver T) WZWZ() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[w],
        receiver[z],
    }
}


// WZXW returns a new float64x4 vector using the WZXW lanes of the original float64x4 vector.
func (receiver T) WZXW() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[x],
        receiver[w],
    }
}


// WZXX returns a new float64x4 vector using the WZXX lanes of the original float64x4 vector.
func (receiver T) WZXX() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[x],
        receiver[x],
    }
}


// WZXY returns a new float64x4 vector using the WZXY lanes of the original float64x4 vector.
func (receiver T) WZXY() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[x],
        receiver[y],
    }
}


// WZXZ returns a new float64x4 vector using the WZXZ lanes of the original float64x4 vector.
func (receiver T) WZXZ() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[x],
        receiver[z],
    }
}


// WZYW returns a new float64x4 vector using the WZYW lanes of the original float64x4 vector.
func (receiver T) WZYW() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[y],
        receiver[w],
    }
}


// WZYX returns a new float64x4 vector using the WZYX lanes of the original float64x4 vector.
func (receiver T) WZYX() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[y],
        receiver[x],
    }
}


// WZYY returns a new float64x4 vector using the WZYY lanes of the original float64x4 vector.
func (receiver T) WZYY() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[y],
        receiver[y],
    }
}


// WZYZ returns a new float64x4 vector using the WZYZ lanes of the original float64x4 vector.
func (receiver T) WZYZ() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[y],
        receiver[z],
    }
}


// WZZW returns a new float64x4 vector using the WZZW lanes of the original float64x4 vector.
func (receiver T) WZZW() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[z],
        receiver[w],
    }
}


// WZZX returns a new float64x4 vector using the WZZX lanes of the original float64x4 vector.
func (receiver T) WZZX() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[z],
        receiver[x],
    }
}


// WZZY returns a new float64x4 vector using the WZZY lanes of the original float64x4 vector.
func (receiver T) WZZY() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[z],
        receiver[y],
    }
}


// WZZZ returns a new float64x4 vector using the WZZZ lanes of the original float64x4 vector.
func (receiver T) WZZZ() T {
    return T{
        receiver[w],
        receiver[z],
        receiver[z],
        receiver[z],
    }
}


// XWWW returns a new float64x4 vector using the XWWW lanes of the original float64x4 vector.
func (receiver T) XWWW() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[w],
        receiver[w],
    }
}


// XWWX returns a new float64x4 vector using the XWWX lanes of the original float64x4 vector.
func (receiver T) XWWX() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[w],
        receiver[x],
    }
}


// XWWY returns a new float64x4 vector using the XWWY lanes of the original float64x4 vector.
func (receiver T) XWWY() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[w],
        receiver[y],
    }
}


// XWWZ returns a new float64x4 vector using the XWWZ lanes of the original float64x4 vector.
func (receiver T) XWWZ() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[w],
        receiver[z],
    }
}


// XWXW returns a new float64x4 vector using the XWXW lanes of the original float64x4 vector.
func (receiver T) XWXW() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[x],
        receiver[w],
    }
}


// XWXX returns a new float64x4 vector using the XWXX lanes of the original float64x4 vector.
func (receiver T) XWXX() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[x],
        receiver[x],
    }
}


// XWXY returns a new float64x4 vector using the XWXY lanes of the original float64x4 vector.
func (receiver T) XWXY() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[x],
        receiver[y],
    }
}


// XWXZ returns a new float64x4 vector using the XWXZ lanes of the original float64x4 vector.
func (receiver T) XWXZ() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[x],
        receiver[z],
    }
}


// XWYW returns a new float64x4 vector using the XWYW lanes of the original float64x4 vector.
func (receiver T) XWYW() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[y],
        receiver[w],
    }
}


// XWYX returns a new float64x4 vector using the XWYX lanes of the original float64x4 vector.
func (receiver T) XWYX() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[y],
        receiver[x],
    }
}


// XWYY returns a new float64x4 vector using the XWYY lanes of the original float64x4 vector.
func (receiver T) XWYY() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[y],
        receiver[y],
    }
}


// XWYZ returns a new float64x4 vector using the XWYZ lanes of the original float64x4 vector.
func (receiver T) XWYZ() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[y],
        receiver[z],
    }
}


// XWZW returns a new float64x4 vector using the XWZW lanes of the original float64x4 vector.
func (receiver T) XWZW() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[z],
        receiver[w],
    }
}


// XWZX returns a new float64x4 vector using the XWZX lanes of the original float64x4 vector.
func (receiver T) XWZX() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[z],
        receiver[x],
    }
}


// XWZY returns a new float64x4 vector using the XWZY lanes of the original float64x4 vector.
func (receiver T) XWZY() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[z],
        receiver[y],
    }
}


// XWZZ returns a new float64x4 vector using the XWZZ lanes of the original float64x4 vector.
func (receiver T) XWZZ() T {
    return T{
        receiver[x],
        receiver[w],
        receiver[z],
        receiver[z],
    }
}


// XXWW returns a new float64x4 vector using the XXWW lanes of the original float64x4 vector.
func (receiver T) XXWW() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[w],
        receiver[w],
    }
}


// XXWX returns a new float64x4 vector using the XXWX lanes of the original float64x4 vector.
func (receiver T) XXWX() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[w],
        receiver[x],
    }
}


// XXWY returns a new float64x4 vector using the XXWY lanes of the original float64x4 vector.
func (receiver T) XXWY() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[w],
        receiver[y],
    }
}


// XXWZ returns a new float64x4 vector using the XXWZ lanes of the original float64x4 vector.
func (receiver T) XXWZ() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[w],
        receiver[z],
    }
}


// XXXW returns a new float64x4 vector using the XXXW lanes of the original float64x4 vector.
func (receiver T) XXXW() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[x],
        receiver[w],
    }
}


// XXXX returns a new float64x4 vector using the XXXX lanes of the original float64x4 vector.
func (receiver T) XXXX() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[x],
        receiver[x],
    }
}


// XXXY returns a new float64x4 vector using the XXXY lanes of the original float64x4 vector.
func (receiver T) XXXY() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[x],
        receiver[y],
    }
}


// XXXZ returns a new float64x4 vector using the XXXZ lanes of the original float64x4 vector.
func (receiver T) XXXZ() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[x],
        receiver[z],
    }
}


// XXYW returns a new float64x4 vector using the XXYW lanes of the original float64x4 vector.
func (receiver T) XXYW() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[y],
        receiver[w],
    }
}


// XXYX returns a new float64x4 vector using the XXYX lanes of the original float64x4 vector.
func (receiver T) XXYX() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[y],
        receiver[x],
    }
}


// XXYY returns a new float64x4 vector using the XXYY lanes of the original float64x4 vector.
func (receiver T) XXYY() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[y],
        receiver[y],
    }
}


// XXYZ returns a new float64x4 vector using the XXYZ lanes of the original float64x4 vector.
func (receiver T) XXYZ() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[y],
        receiver[z],
    }
}


// XXZW returns a new float64x4 vector using the XXZW lanes of the original float64x4 vector.
func (receiver T) XXZW() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[z],
        receiver[w],
    }
}


// XXZX returns a new float64x4 vector using the XXZX lanes of the original float64x4 vector.
func (receiver T) XXZX() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[z],
        receiver[x],
    }
}


// XXZY returns a new float64x4 vector using the XXZY lanes of the original float64x4 vector.
func (receiver T) XXZY() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[z],
        receiver[y],
    }
}


// XXZZ returns a new float64x4 vector using the XXZZ lanes of the original float64x4 vector.
func (receiver T) XXZZ() T {
    return T{
        receiver[x],
        receiver[x],
        receiver[z],
        receiver[z],
    }
}


// XYWW returns a new float64x4 vector using the XYWW lanes of the original float64x4 vector.
func (receiver T) XYWW() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[w],
        receiver[w],
    }
}


// XYWX returns a new float64x4 vector using the XYWX lanes of the original float64x4 vector.
func (receiver T) XYWX() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[w],
        receiver[x],
    }
}


// XYWY returns a new float64x4 vector using the XYWY lanes of the original float64x4 vector.
func (receiver T) XYWY() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[w],
        receiver[y],
    }
}


// XYWZ returns a new float64x4 vector using the XYWZ lanes of the original float64x4 vector.
func (receiver T) XYWZ() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[w],
        receiver[z],
    }
}


// XYXW returns a new float64x4 vector using the XYXW lanes of the original float64x4 vector.
func (receiver T) XYXW() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[x],
        receiver[w],
    }
}


// XYXX returns a new float64x4 vector using the XYXX lanes of the original float64x4 vector.
func (receiver T) XYXX() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[x],
        receiver[x],
    }
}


// XYXY returns a new float64x4 vector using the XYXY lanes of the original float64x4 vector.
func (receiver T) XYXY() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[x],
        receiver[y],
    }
}


// XYXZ returns a new float64x4 vector using the XYXZ lanes of the original float64x4 vector.
func (receiver T) XYXZ() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[x],
        receiver[z],
    }
}


// XYYW returns a new float64x4 vector using the XYYW lanes of the original float64x4 vector.
func (receiver T) XYYW() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[y],
        receiver[w],
    }
}


// XYYX returns a new float64x4 vector using the XYYX lanes of the original float64x4 vector.
func (receiver T) XYYX() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[y],
        receiver[x],
    }
}


// XYYY returns a new float64x4 vector using the XYYY lanes of the original float64x4 vector.
func (receiver T) XYYY() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[y],
        receiver[y],
    }
}


// XYYZ returns a new float64x4 vector using the XYYZ lanes of the original float64x4 vector.
func (receiver T) XYYZ() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[y],
        receiver[z],
    }
}


// XYZW returns a new float64x4 vector using the XYZW lanes of the original float64x4 vector.
func (receiver T) XYZW() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[z],
        receiver[w],
    }
}


// XYZX returns a new float64x4 vector using the XYZX lanes of the original float64x4 vector.
func (receiver T) XYZX() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[z],
        receiver[x],
    }
}


// XYZY returns a new float64x4 vector using the XYZY lanes of the original float64x4 vector.
func (receiver T) XYZY() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[z],
        receiver[y],
    }
}


// XYZZ returns a new float64x4 vector using the XYZZ lanes of the original float64x4 vector.
func (receiver T) XYZZ() T {
    return T{
        receiver[x],
        receiver[y],
        receiver[z],
        receiver[z],
    }
}


// XZWW returns a new float64x4 vector using the XZWW lanes of the original float64x4 vector.
func (receiver T) XZWW() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[w],
        receiver[w],
    }
}


// XZWX returns a new float64x4 vector using the XZWX lanes of the original float64x4 vector.
func (receiver T) XZWX() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[w],
        receiver[x],
    }
}


// XZWY returns a new float64x4 vector using the XZWY lanes of the original float64x4 vector.
func (receiver T) XZWY() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[w],
        receiver[y],
    }
}


// XZWZ returns a new float64x4 vector using the XZWZ lanes of the original float64x4 vector.
func (receiver T) XZWZ() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[w],
        receiver[z],
    }
}


// XZXW returns a new float64x4 vector using the XZXW lanes of the original float64x4 vector.
func (receiver T) XZXW() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[x],
        receiver[w],
    }
}


// XZXX returns a new float64x4 vector using the XZXX lanes of the original float64x4 vector.
func (receiver T) XZXX() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[x],
        receiver[x],
    }
}


// XZXY returns a new float64x4 vector using the XZXY lanes of the original float64x4 vector.
func (receiver T) XZXY() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[x],
        receiver[y],
    }
}


// XZXZ returns a new float64x4 vector using the XZXZ lanes of the original float64x4 vector.
func (receiver T) XZXZ() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[x],
        receiver[z],
    }
}


// XZYW returns a new float64x4 vector using the XZYW lanes of the original float64x4 vector.
func (receiver T) XZYW() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[y],
        receiver[w],
    }
}


// XZYX returns a new float64x4 vector using the XZYX lanes of the original float64x4 vector.
func (receiver T) XZYX() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[y],
        receiver[x],
    }
}


// XZYY returns a new float64x4 vector using the XZYY lanes of the original float64x4 vector.
func (receiver T) XZYY() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[y],
        receiver[y],
    }
}


// XZYZ returns a new float64x4 vector using the XZYZ lanes of the original float64x4 vector.
func (receiver T) XZYZ() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[y],
        receiver[z],
    }
}


// XZZW returns a new float64x4 vector using the XZZW lanes of the original float64x4 vector.
func (receiver T) XZZW() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[z],
        receiver[w],
    }
}


// XZZX returns a new float64x4 vector using the XZZX lanes of the original float64x4 vector.
func (receiver T) XZZX() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[z],
        receiver[x],
    }
}


// XZZY returns a new float64x4 vector using the XZZY lanes of the original float64x4 vector.
func (receiver T) XZZY() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[z],
        receiver[y],
    }
}


// XZZZ returns a new float64x4 vector using the XZZZ lanes of the original float64x4 vector.
func (receiver T) XZZZ() T {
    return T{
        receiver[x],
        receiver[z],
        receiver[z],
        receiver[z],
    }
}


// YWWW returns a new float64x4 vector using the YWWW lanes of the original float64x4 vector.
func (receiver T) YWWW() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[w],
        receiver[w],
    }
}


// YWWX returns a new float64x4 vector using the YWWX lanes of the original float64x4 vector.
func (receiver T) YWWX() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[w],
        receiver[x],
    }
}


// YWWY returns a new float64x4 vector using the YWWY lanes of the original float64x4 vector.
func (receiver T) YWWY() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[w],
        receiver[y],
    }
}


// YWWZ returns a new float64x4 vector using the YWWZ lanes of the original float64x4 vector.
func (receiver T) YWWZ() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[w],
        receiver[z],
    }
}


// YWXW returns a new float64x4 vector using the YWXW lanes of the original float64x4 vector.
func (receiver T) YWXW() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[x],
        receiver[w],
    }
}


// YWXX returns a new float64x4 vector using the YWXX lanes of the original float64x4 vector.
func (receiver T) YWXX() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[x],
        receiver[x],
    }
}


// YWXY returns a new float64x4 vector using the YWXY lanes of the original float64x4 vector.
func (receiver T) YWXY() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[x],
        receiver[y],
    }
}


// YWXZ returns a new float64x4 vector using the YWXZ lanes of the original float64x4 vector.
func (receiver T) YWXZ() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[x],
        receiver[z],
    }
}


// YWYW returns a new float64x4 vector using the YWYW lanes of the original float64x4 vector.
func (receiver T) YWYW() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[y],
        receiver[w],
    }
}


// YWYX returns a new float64x4 vector using the YWYX lanes of the original float64x4 vector.
func (receiver T) YWYX() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[y],
        receiver[x],
    }
}


// YWYY returns a new float64x4 vector using the YWYY lanes of the original float64x4 vector.
func (receiver T) YWYY() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[y],
        receiver[y],
    }
}


// YWYZ returns a new float64x4 vector using the YWYZ lanes of the original float64x4 vector.
func (receiver T) YWYZ() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[y],
        receiver[z],
    }
}


// YWZW returns a new float64x4 vector using the YWZW lanes of the original float64x4 vector.
func (receiver T) YWZW() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[z],
        receiver[w],
    }
}


// YWZX returns a new float64x4 vector using the YWZX lanes of the original float64x4 vector.
func (receiver T) YWZX() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[z],
        receiver[x],
    }
}


// YWZY returns a new float64x4 vector using the YWZY lanes of the original float64x4 vector.
func (receiver T) YWZY() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[z],
        receiver[y],
    }
}


// YWZZ returns a new float64x4 vector using the YWZZ lanes of the original float64x4 vector.
func (receiver T) YWZZ() T {
    return T{
        receiver[y],
        receiver[w],
        receiver[z],
        receiver[z],
    }
}


// YXWW returns a new float64x4 vector using the YXWW lanes of the original float64x4 vector.
func (receiver T) YXWW() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[w],
        receiver[w],
    }
}


// YXWX returns a new float64x4 vector using the YXWX lanes of the original float64x4 vector.
func (receiver T) YXWX() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[w],
        receiver[x],
    }
}


// YXWY returns a new float64x4 vector using the YXWY lanes of the original float64x4 vector.
func (receiver T) YXWY() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[w],
        receiver[y],
    }
}


// YXWZ returns a new float64x4 vector using the YXWZ lanes of the original float64x4 vector.
func (receiver T) YXWZ() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[w],
        receiver[z],
    }
}


// YXXW returns a new float64x4 vector using the YXXW lanes of the original float64x4 vector.
func (receiver T) YXXW() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[x],
        receiver[w],
    }
}


// YXXX returns a new float64x4 vector using the YXXX lanes of the original float64x4 vector.
func (receiver T) YXXX() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[x],
        receiver[x],
    }
}


// YXXY returns a new float64x4 vector using the YXXY lanes of the original float64x4 vector.
func (receiver T) YXXY() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[x],
        receiver[y],
    }
}


// YXXZ returns a new float64x4 vector using the YXXZ lanes of the original float64x4 vector.
func (receiver T) YXXZ() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[x],
        receiver[z],
    }
}


// YXYW returns a new float64x4 vector using the YXYW lanes of the original float64x4 vector.
func (receiver T) YXYW() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[y],
        receiver[w],
    }
}


// YXYX returns a new float64x4 vector using the YXYX lanes of the original float64x4 vector.
func (receiver T) YXYX() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[y],
        receiver[x],
    }
}


// YXYY returns a new float64x4 vector using the YXYY lanes of the original float64x4 vector.
func (receiver T) YXYY() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[y],
        receiver[y],
    }
}


// YXYZ returns a new float64x4 vector using the YXYZ lanes of the original float64x4 vector.
func (receiver T) YXYZ() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[y],
        receiver[z],
    }
}


// YXZW returns a new float64x4 vector using the YXZW lanes of the original float64x4 vector.
func (receiver T) YXZW() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[z],
        receiver[w],
    }
}


// YXZX returns a new float64x4 vector using the YXZX lanes of the original float64x4 vector.
func (receiver T) YXZX() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[z],
        receiver[x],
    }
}


// YXZY returns a new float64x4 vector using the YXZY lanes of the original float64x4 vector.
func (receiver T) YXZY() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[z],
        receiver[y],
    }
}


// YXZZ returns a new float64x4 vector using the YXZZ lanes of the original float64x4 vector.
func (receiver T) YXZZ() T {
    return T{
        receiver[y],
        receiver[x],
        receiver[z],
        receiver[z],
    }
}


// YYWW returns a new float64x4 vector using the YYWW lanes of the original float64x4 vector.
func (receiver T) YYWW() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[w],
        receiver[w],
    }
}


// YYWX returns a new float64x4 vector using the YYWX lanes of the original float64x4 vector.
func (receiver T) YYWX() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[w],
        receiver[x],
    }
}


// YYWY returns a new float64x4 vector using the YYWY lanes of the original float64x4 vector.
func (receiver T) YYWY() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[w],
        receiver[y],
    }
}


// YYWZ returns a new float64x4 vector using the YYWZ lanes of the original float64x4 vector.
func (receiver T) YYWZ() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[w],
        receiver[z],
    }
}


// YYXW returns a new float64x4 vector using the YYXW lanes of the original float64x4 vector.
func (receiver T) YYXW() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[x],
        receiver[w],
    }
}


// YYXX returns a new float64x4 vector using the YYXX lanes of the original float64x4 vector.
func (receiver T) YYXX() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[x],
        receiver[x],
    }
}


// YYXY returns a new float64x4 vector using the YYXY lanes of the original float64x4 vector.
func (receiver T) YYXY() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[x],
        receiver[y],
    }
}


// YYXZ returns a new float64x4 vector using the YYXZ lanes of the original float64x4 vector.
func (receiver T) YYXZ() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[x],
        receiver[z],
    }
}


// YYYW returns a new float64x4 vector using the YYYW lanes of the original float64x4 vector.
func (receiver T) YYYW() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[y],
        receiver[w],
    }
}


// YYYX returns a new float64x4 vector using the YYYX lanes of the original float64x4 vector.
func (receiver T) YYYX() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[y],
        receiver[x],
    }
}


// YYYY returns a new float64x4 vector using the YYYY lanes of the original float64x4 vector.
func (receiver T) YYYY() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[y],
        receiver[y],
    }
}


// YYYZ returns a new float64x4 vector using the YYYZ lanes of the original float64x4 vector.
func (receiver T) YYYZ() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[y],
        receiver[z],
    }
}


// YYZW returns a new float64x4 vector using the YYZW lanes of the original float64x4 vector.
func (receiver T) YYZW() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[z],
        receiver[w],
    }
}


// YYZX returns a new float64x4 vector using the YYZX lanes of the original float64x4 vector.
func (receiver T) YYZX() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[z],
        receiver[x],
    }
}


// YYZY returns a new float64x4 vector using the YYZY lanes of the original float64x4 vector.
func (receiver T) YYZY() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[z],
        receiver[y],
    }
}


// YYZZ returns a new float64x4 vector using the YYZZ lanes of the original float64x4 vector.
func (receiver T) YYZZ() T {
    return T{
        receiver[y],
        receiver[y],
        receiver[z],
        receiver[z],
    }
}


// YZWW returns a new float64x4 vector using the YZWW lanes of the original float64x4 vector.
func (receiver T) YZWW() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[w],
        receiver[w],
    }
}


// YZWX returns a new float64x4 vector using the YZWX lanes of the original float64x4 vector.
func (receiver T) YZWX() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[w],
        receiver[x],
    }
}


// YZWY returns a new float64x4 vector using the YZWY lanes of the original float64x4 vector.
func (receiver T) YZWY() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[w],
        receiver[y],
    }
}


// YZWZ returns a new float64x4 vector using the YZWZ lanes of the original float64x4 vector.
func (receiver T) YZWZ() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[w],
        receiver[z],
    }
}


// YZXW returns a new float64x4 vector using the YZXW lanes of the original float64x4 vector.
func (receiver T) YZXW() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[x],
        receiver[w],
    }
}


// YZXX returns a new float64x4 vector using the YZXX lanes of the original float64x4 vector.
func (receiver T) YZXX() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[x],
        receiver[x],
    }
}


// YZXY returns a new float64x4 vector using the YZXY lanes of the original float64x4 vector.
func (receiver T) YZXY() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[x],
        receiver[y],
    }
}


// YZXZ returns a new float64x4 vector using the YZXZ lanes of the original float64x4 vector.
func (receiver T) YZXZ() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[x],
        receiver[z],
    }
}


// YZYW returns a new float64x4 vector using the YZYW lanes of the original float64x4 vector.
func (receiver T) YZYW() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[y],
        receiver[w],
    }
}


// YZYX returns a new float64x4 vector using the YZYX lanes of the original float64x4 vector.
func (receiver T) YZYX() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[y],
        receiver[x],
    }
}


// YZYY returns a new float64x4 vector using the YZYY lanes of the original float64x4 vector.
func (receiver T) YZYY() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[y],
        receiver[y],
    }
}


// YZYZ returns a new float64x4 vector using the YZYZ lanes of the original float64x4 vector.
func (receiver T) YZYZ() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[y],
        receiver[z],
    }
}


// YZZW returns a new float64x4 vector using the YZZW lanes of the original float64x4 vector.
func (receiver T) YZZW() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[z],
        receiver[w],
    }
}


// YZZX returns a new float64x4 vector using the YZZX lanes of the original float64x4 vector.
func (receiver T) YZZX() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[z],
        receiver[x],
    }
}


// YZZY returns a new float64x4 vector using the YZZY lanes of the original float64x4 vector.
func (receiver T) YZZY() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[z],
        receiver[y],
    }
}


// YZZZ returns a new float64x4 vector using the YZZZ lanes of the original float64x4 vector.
func (receiver T) YZZZ() T {
    return T{
        receiver[y],
        receiver[z],
        receiver[z],
        receiver[z],
    }
}


// ZWWW returns a new float64x4 vector using the ZWWW lanes of the original float64x4 vector.
func (receiver T) ZWWW() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[w],
        receiver[w],
    }
}


// ZWWX returns a new float64x4 vector using the ZWWX lanes of the original float64x4 vector.
func (receiver T) ZWWX() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[w],
        receiver[x],
    }
}


// ZWWY returns a new float64x4 vector using the ZWWY lanes of the original float64x4 vector.
func (receiver T) ZWWY() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[w],
        receiver[y],
    }
}


// ZWWZ returns a new float64x4 vector using the ZWWZ lanes of the original float64x4 vector.
func (receiver T) ZWWZ() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[w],
        receiver[z],
    }
}


// ZWXW returns a new float64x4 vector using the ZWXW lanes of the original float64x4 vector.
func (receiver T) ZWXW() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[x],
        receiver[w],
    }
}


// ZWXX returns a new float64x4 vector using the ZWXX lanes of the original float64x4 vector.
func (receiver T) ZWXX() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[x],
        receiver[x],
    }
}


// ZWXY returns a new float64x4 vector using the ZWXY lanes of the original float64x4 vector.
func (receiver T) ZWXY() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[x],
        receiver[y],
    }
}


// ZWXZ returns a new float64x4 vector using the ZWXZ lanes of the original float64x4 vector.
func (receiver T) ZWXZ() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[x],
        receiver[z],
    }
}


// ZWYW returns a new float64x4 vector using the ZWYW lanes of the original float64x4 vector.
func (receiver T) ZWYW() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[y],
        receiver[w],
    }
}


// ZWYX returns a new float64x4 vector using the ZWYX lanes of the original float64x4 vector.
func (receiver T) ZWYX() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[y],
        receiver[x],
    }
}


// ZWYY returns a new float64x4 vector using the ZWYY lanes of the original float64x4 vector.
func (receiver T) ZWYY() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[y],
        receiver[y],
    }
}


// ZWYZ returns a new float64x4 vector using the ZWYZ lanes of the original float64x4 vector.
func (receiver T) ZWYZ() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[y],
        receiver[z],
    }
}


// ZWZW returns a new float64x4 vector using the ZWZW lanes of the original float64x4 vector.
func (receiver T) ZWZW() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[z],
        receiver[w],
    }
}


// ZWZX returns a new float64x4 vector using the ZWZX lanes of the original float64x4 vector.
func (receiver T) ZWZX() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[z],
        receiver[x],
    }
}


// ZWZY returns a new float64x4 vector using the ZWZY lanes of the original float64x4 vector.
func (receiver T) ZWZY() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[z],
        receiver[y],
    }
}


// ZWZZ returns a new float64x4 vector using the ZWZZ lanes of the original float64x4 vector.
func (receiver T) ZWZZ() T {
    return T{
        receiver[z],
        receiver[w],
        receiver[z],
        receiver[z],
    }
}


// ZXWW returns a new float64x4 vector using the ZXWW lanes of the original float64x4 vector.
func (receiver T) ZXWW() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[w],
        receiver[w],
    }
}


// ZXWX returns a new float64x4 vector using the ZXWX lanes of the original float64x4 vector.
func (receiver T) ZXWX() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[w],
        receiver[x],
    }
}


// ZXWY returns a new float64x4 vector using the ZXWY lanes of the original float64x4 vector.
func (receiver T) ZXWY() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[w],
        receiver[y],
    }
}


// ZXWZ returns a new float64x4 vector using the ZXWZ lanes of the original float64x4 vector.
func (receiver T) ZXWZ() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[w],
        receiver[z],
    }
}


// ZXXW returns a new float64x4 vector using the ZXXW lanes of the original float64x4 vector.
func (receiver T) ZXXW() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[x],
        receiver[w],
    }
}


// ZXXX returns a new float64x4 vector using the ZXXX lanes of the original float64x4 vector.
func (receiver T) ZXXX() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[x],
        receiver[x],
    }
}


// ZXXY returns a new float64x4 vector using the ZXXY lanes of the original float64x4 vector.
func (receiver T) ZXXY() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[x],
        receiver[y],
    }
}


// ZXXZ returns a new float64x4 vector using the ZXXZ lanes of the original float64x4 vector.
func (receiver T) ZXXZ() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[x],
        receiver[z],
    }
}


// ZXYW returns a new float64x4 vector using the ZXYW lanes of the original float64x4 vector.
func (receiver T) ZXYW() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[y],
        receiver[w],
    }
}


// ZXYX returns a new float64x4 vector using the ZXYX lanes of the original float64x4 vector.
func (receiver T) ZXYX() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[y],
        receiver[x],
    }
}


// ZXYY returns a new float64x4 vector using the ZXYY lanes of the original float64x4 vector.
func (receiver T) ZXYY() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[y],
        receiver[y],
    }
}


// ZXYZ returns a new float64x4 vector using the ZXYZ lanes of the original float64x4 vector.
func (receiver T) ZXYZ() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[y],
        receiver[z],
    }
}


// ZXZW returns a new float64x4 vector using the ZXZW lanes of the original float64x4 vector.
func (receiver T) ZXZW() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[z],
        receiver[w],
    }
}


// ZXZX returns a new float64x4 vector using the ZXZX lanes of the original float64x4 vector.
func (receiver T) ZXZX() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[z],
        receiver[x],
    }
}


// ZXZY returns a new float64x4 vector using the ZXZY lanes of the original float64x4 vector.
func (receiver T) ZXZY() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[z],
        receiver[y],
    }
}


// ZXZZ returns a new float64x4 vector using the ZXZZ lanes of the original float64x4 vector.
func (receiver T) ZXZZ() T {
    return T{
        receiver[z],
        receiver[x],
        receiver[z],
        receiver[z],
    }
}


// ZYWW returns a new float64x4 vector using the ZYWW lanes of the original float64x4 vector.
func (receiver T) ZYWW() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[w],
        receiver[w],
    }
}


// ZYWX returns a new float64x4 vector using the ZYWX lanes of the original float64x4 vector.
func (receiver T) ZYWX() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[w],
        receiver[x],
    }
}


// ZYWY returns a new float64x4 vector using the ZYWY lanes of the original float64x4 vector.
func (receiver T) ZYWY() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[w],
        receiver[y],
    }
}


// ZYWZ returns a new float64x4 vector using the ZYWZ lanes of the original float64x4 vector.
func (receiver T) ZYWZ() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[w],
        receiver[z],
    }
}


// ZYXW returns a new float64x4 vector using the ZYXW lanes of the original float64x4 vector.
func (receiver T) ZYXW() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[x],
        receiver[w],
    }
}


// ZYXX returns a new float64x4 vector using the ZYXX lanes of the original float64x4 vector.
func (receiver T) ZYXX() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[x],
        receiver[x],
    }
}


// ZYXY returns a new float64x4 vector using the ZYXY lanes of the original float64x4 vector.
func (receiver T) ZYXY() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[x],
        receiver[y],
    }
}


// ZYXZ returns a new float64x4 vector using the ZYXZ lanes of the original float64x4 vector.
func (receiver T) ZYXZ() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[x],
        receiver[z],
    }
}


// ZYYW returns a new float64x4 vector using the ZYYW lanes of the original float64x4 vector.
func (receiver T) ZYYW() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[y],
        receiver[w],
    }
}


// ZYYX returns a new float64x4 vector using the ZYYX lanes of the original float64x4 vector.
func (receiver T) ZYYX() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[y],
        receiver[x],
    }
}


// ZYYY returns a new float64x4 vector using the ZYYY lanes of the original float64x4 vector.
func (receiver T) ZYYY() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[y],
        receiver[y],
    }
}


// ZYYZ returns a new float64x4 vector using the ZYYZ lanes of the original float64x4 vector.
func (receiver T) ZYYZ() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[y],
        receiver[z],
    }
}


// ZYZW returns a new float64x4 vector using the ZYZW lanes of the original float64x4 vector.
func (receiver T) ZYZW() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[z],
        receiver[w],
    }
}


// ZYZX returns a new float64x4 vector using the ZYZX lanes of the original float64x4 vector.
func (receiver T) ZYZX() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[z],
        receiver[x],
    }
}


// ZYZY returns a new float64x4 vector using the ZYZY lanes of the original float64x4 vector.
func (receiver T) ZYZY() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[z],
        receiver[y],
    }
}


// ZYZZ returns a new float64x4 vector using the ZYZZ lanes of the original float64x4 vector.
func (receiver T) ZYZZ() T {
    return T{
        receiver[z],
        receiver[y],
        receiver[z],
        receiver[z],
    }
}


// ZZWW returns a new float64x4 vector using the ZZWW lanes of the original float64x4 vector.
func (receiver T) ZZWW() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[w],
        receiver[w],
    }
}


// ZZWX returns a new float64x4 vector using the ZZWX lanes of the original float64x4 vector.
func (receiver T) ZZWX() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[w],
        receiver[x],
    }
}


// ZZWY returns a new float64x4 vector using the ZZWY lanes of the original float64x4 vector.
func (receiver T) ZZWY() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[w],
        receiver[y],
    }
}


// ZZWZ returns a new float64x4 vector using the ZZWZ lanes of the original float64x4 vector.
func (receiver T) ZZWZ() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[w],
        receiver[z],
    }
}


// ZZXW returns a new float64x4 vector using the ZZXW lanes of the original float64x4 vector.
func (receiver T) ZZXW() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[x],
        receiver[w],
    }
}


// ZZXX returns a new float64x4 vector using the ZZXX lanes of the original float64x4 vector.
func (receiver T) ZZXX() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[x],
        receiver[x],
    }
}


// ZZXY returns a new float64x4 vector using the ZZXY lanes of the original float64x4 vector.
func (receiver T) ZZXY() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[x],
        receiver[y],
    }
}


// ZZXZ returns a new float64x4 vector using the ZZXZ lanes of the original float64x4 vector.
func (receiver T) ZZXZ() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[x],
        receiver[z],
    }
}


// ZZYW returns a new float64x4 vector using the ZZYW lanes of the original float64x4 vector.
func (receiver T) ZZYW() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[y],
        receiver[w],
    }
}


// ZZYX returns a new float64x4 vector using the ZZYX lanes of the original float64x4 vector.
func (receiver T) ZZYX() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[y],
        receiver[x],
    }
}


// ZZYY returns a new float64x4 vector using the ZZYY lanes of the original float64x4 vector.
func (receiver T) ZZYY() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[y],
        receiver[y],
    }
}


// ZZYZ returns a new float64x4 vector using the ZZYZ lanes of the original float64x4 vector.
func (receiver T) ZZYZ() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[y],
        receiver[z],
    }
}


// ZZZW returns a new float64x4 vector using the ZZZW lanes of the original float64x4 vector.
func (receiver T) ZZZW() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[z],
        receiver[w],
    }
}


// ZZZX returns a new float64x4 vector using the ZZZX lanes of the original float64x4 vector.
func (receiver T) ZZZX() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[z],
        receiver[x],
    }
}


// ZZZY returns a new float64x4 vector using the ZZZY lanes of the original float64x4 vector.
func (receiver T) ZZZY() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[z],
        receiver[y],
    }
}


// ZZZZ returns a new float64x4 vector using the ZZZZ lanes of the original float64x4 vector.
func (receiver T) ZZZZ() T {
    return T{
        receiver[z],
        receiver[z],
        receiver[z],
        receiver[z],
    }
}


