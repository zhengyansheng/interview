# reflect.DeepEqual

## 函数签名
```go
func DeepEqual(x, y interface{}) bool
```

## 函数代码
```go
func DeepEqual(x, y interface{}) bool {
	// 1. 如果x, y都是空接口，则返回true
	if x == nil || y == nil {
		return x == y
	}
	/* 
	    2. 通过反射获取变量的类型
	    如果x, y的类型不相等，则返回false
	 */
	v1 := ValueOf(x)
	v2 := ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return deepValueEqual(v1, v2, make(map[visit]bool), 0)
}
```

```go
type visit struct {
    a1  unsafe.Pointer
    a2  unsafe.Pointer
    typ Type
}

func deepValueEqual(v1, v2 Value, visited map[visit]bool, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}

	// if depth > 10 { panic("deepValueEqual") }	// for debugging

	// We want to avoid putting more in the visited map than we need to.
	// For any possible reference cycle that might be encountered,
	// hard(v1, v2) needs to return true for at least one of the types in the cycle,
	// and it's safe and valid to get Value's internal pointer.
	hard := func(v1, v2 Value) bool {
		switch v1.Kind() {
		case Ptr:
			if v1.typ.ptrdata == 0 {
				// go:notinheap pointers can't be cyclic.
				// At least, all of our current uses of go:notinheap have
				// that property. The runtime ones aren't cyclic (and we don't use
				// DeepEqual on them anyway), and the cgo-generated ones are
				// all empty structs.
				return false
			}
			fallthrough
		case Map, Slice, Interface:
			// Nil pointers cannot be cyclic. Avoid putting them in the visited map.
			return !v1.IsNil() && !v2.IsNil()
		}
		return false
	}

	if hard(v1, v2) {
		// For a Ptr or Map value, we need to check flagIndir,
		// which we do by calling the pointer method.
		// For Slice or Interface, flagIndir is always set,
		// and using v.ptr suffices.
		ptrval := func(v Value) unsafe.Pointer {
			switch v.Kind() {
			case Ptr, Map:
				return v.pointer()
			default:
				return v.ptr
			}
		}
		addr1 := ptrval(v1)
		addr2 := ptrval(v2)
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}

		// Short circuit if references are already seen.
		typ := v1.Type()
		v := visit{addr1, addr2, typ}
		if visited[v] {
			return true
		}

		// Remember for later.
		visited[v] = true
	}

	/*
	    获取 v1 的类型
	 */
	switch v1.Kind() {
	case Array:
		for i := 0; i < v1.Len(); i++ {
			/*
			    v1.Index(i) 返回v索引的元素
			 */
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited, depth+1) {
				return false
			}
		}
		return true
	case Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
            /*
               v1.Index(i) 返回v索引的元素
            */
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited, depth+1) {
				return false
			}
		}
		return true
	case Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited, depth+1)
	case Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited, depth+1)
	case Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i), visited, depth+1) {
				return false
			}
		}
		return true
	case Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(val1, val2, visited, depth+1) {
				return false
			}
		}
		return true
	case Func:
		if v1.IsNil() && v2.IsNil() {
			return true
		}
		// Can't do better than this:
		return false
	default:
		// Normal equality suffices
		return valueInterface(v1, false) == valueInterface(v2, false)
	}
}
```