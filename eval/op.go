// Copyright 2015 The Neugram Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eval

import (
	"fmt"
	"math/big"
	"reflect"

	"neugram.io/ng/token"
)

// TODO redo
func valEq(x, y interface{}) bool {
	if x == y {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	switch x := x.(type) {
	case *big.Int:
		switch y := y.(type) {
		case *big.Int:
			return x.Cmp(y) == 0
		}
	case *big.Float:
		switch y := y.(type) {
		case *big.Float:
			return x.Cmp(y) == 0
		}
		/*case *StructVal:
		switch y := y.(type) {
		case *StructVal:
			if len(x.Fields) != len(y.Fields) { // TODO compare tipe.Type
				return false
			}
			for i := range x.Fields {
				if !valEq(x.Fields[i], y.Fields[i]) {
					return false
				}
			}
			return true
		}
		*/
	}
	return false
}

func binOp(op token.Token, x, y interface{}) (interface{}, error) {
	switch op {
	case token.Add:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x + y, nil
			}
		case int8:
			switch y := y.(type) {
			case int8:
				return x + y, nil
			}
		case int16:
			switch y := y.(type) {
			case int16:
				return x + y, nil
			}
		case int32:
			switch y := y.(type) {
			case int32:
				return x + y, nil
			}
		case int64:
			switch y := y.(type) {
			case int64:
				return x + y, nil
			}
		case uint:
			switch y := y.(type) {
			case uint:
				return x + y, nil
			}
		case uint8:
			switch y := y.(type) {
			case uint8:
				return x + y, nil
			}
		case uint16:
			switch y := y.(type) {
			case uint16:
				return x + y, nil
			}
		case uint32:
			switch y := y.(type) {
			case uint32:
				return x + y, nil
			}
		case uint64:
			switch y := y.(type) {
			case uint64:
				return x + y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x + y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x + y, nil
			case UntypedFloat:
				yf, _ := y.Float64()
				return x + yf, nil
			}
		case complex64:
			switch y := y.(type) {
			case complex64:
				return x + y, nil
			}
		case complex128:
			switch y := y.(type) {
			case complex128:
				return x + y, nil
			case UntypedInt:
				return x + complex(float64(y.Int64()), 0), nil
			}
		case UntypedInt:
			switch y := y.(type) {
			case UntypedInt:
				z := big.NewInt(0)
				return UntypedInt{z.Add(x.Int, y.Int)}, nil
			case UntypedFloat:
				z := big.NewFloat(float64(x.Int.Int64()))
				return UntypedFloat{z.Add(z, y.Float)}, nil
			case UntypedComplex:
				re := big.NewFloat(float64(x.Int.Int64()))
				im := big.NewFloat(0)
				return UntypedComplex{re.Add(re, y.Real), im.Add(im, y.Imag)}, nil
			case float64:
				return float64(x.Int64()) + y, nil
			case complex64:
				return complex(float32(x.Int64()), 0) + y, nil
			case complex128:
				return complex(float64(x.Int64()), 0) + y, nil
			}
		case UntypedFloat:
			z := big.NewFloat(0)
			switch y := y.(type) {
			case UntypedInt:
				z.SetInt(y.Int)
				return UntypedFloat{z.Add(z, x.Float)}, nil
			case UntypedFloat:
				return UntypedFloat{z.Add(x.Float, y.Float)}, nil
			case UntypedComplex:
				re := big.NewFloat(0)
				im := big.NewFloat(0)
				return UntypedComplex{re.Add(x.Float, y.Real), im.Add(im, y.Imag)}, nil
			case float32:
				v, _ := x.Float32()
				return v + y, nil
			case float64:
				v, _ := x.Float64()
				return v + y, nil
			case complex64:
				v, _ := x.Float32()
				return complex(v, 0) + y, nil
			case complex128:
				v, _ := x.Float64()
				return complex(v, 0) + y, nil
			}
		case UntypedComplex:
			re := big.NewFloat(0)
			im := big.NewFloat(0)
			switch y := y.(type) {
			case UntypedInt:
				re.SetInt(y.Int)
				return UntypedComplex{Real: re.Add(re, x.Real), Imag: im}, nil
			case UntypedFloat:
				re.Set(y.Float)
				return UntypedComplex{Real: re.Add(re, x.Real), Imag: im}, nil
			case UntypedComplex:
				return UntypedComplex{
					Real: re.Add(x.Real, y.Real),
					Imag: im.Add(x.Imag, y.Imag),
				}, nil
			case complex64:
				re, _ := x.Real.Float32()
				im, _ := x.Imag.Float32()
				return complex(re, im) + y, nil
			case complex128:
				re, _ := x.Real.Float64()
				im, _ := x.Imag.Float64()
				return complex(re, im) + y, nil
			}
		case UntypedString:
			switch y := y.(type) {
			case UntypedString:
				return UntypedString{x.String + y.String}, nil
			case string:
				return x.String + y, nil
			}
		case string:
			switch y := y.(type) {
			case UntypedString:
				return x + y.String, nil
			case string:
				return x + y, nil
			}
		default:
			xv := reflect.ValueOf(x)
			yv := reflect.ValueOf(y)
			res := xv.MethodByName("Add").Call([]reflect.Value{yv})
			return res[0].Interface(), nil
		}
	case token.Sub:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x - y, nil
			}
		case int8:
			switch y := y.(type) {
			case int8:
				return x - y, nil
			}
		case int16:
			switch y := y.(type) {
			case int16:
				return x - y, nil
			}
		case int32:
			switch y := y.(type) {
			case int32:
				return x - y, nil
			}
		case int64:
			switch y := y.(type) {
			case int64:
				return x - y, nil
			}
		case uint:
			switch y := y.(type) {
			case uint:
				return x - y, nil
			}
		case uint8:
			switch y := y.(type) {
			case uint8:
				return x - y, nil
			}
		case uint16:
			switch y := y.(type) {
			case uint16:
				return x - y, nil
			}
		case uint32:
			switch y := y.(type) {
			case uint32:
				return x - y, nil
			}
		case uint64:
			switch y := y.(type) {
			case uint64:
				return x - y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x - y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x - y, nil
			}
		case complex64:
			switch y := y.(type) {
			case complex64:
				return x - y, nil
			}
		case complex128:
			switch y := y.(type) {
			case complex128:
				return x - y, nil
			}
		case UntypedInt:
			switch y := y.(type) {
			case UntypedFloat:
				z := big.NewFloat(0)
				xf := big.NewFloat(float64(x.Int.Int64()))
				return UntypedFloat{z.Sub(xf, y.Float)}, nil
			case UntypedInt:
				z := big.NewInt(0)
				return UntypedInt{z.Sub(x.Int, y.Int)}, nil
			case UntypedComplex:
				re := big.NewFloat(0)
				xf := big.NewFloat(float64(x.Int.Int64()))
				im := big.NewFloat(0)
				return UntypedComplex{re.Sub(xf, y.Real), im.Sub(im, y.Imag)}, nil
			}
		case UntypedFloat:
			z := big.NewFloat(0)
			switch y := y.(type) {
			case UntypedInt:
				yf := big.NewFloat(0)
				yf.SetInt(y.Int)
				return UntypedFloat{z.Sub(x.Float, yf)}, nil
			case UntypedFloat:
				return UntypedFloat{z.Sub(x.Float, y.Float)}, nil
			case UntypedComplex:
				return UntypedComplex{z.Sub(x.Float, y.Real), big.NewFloat(0)}, nil
			}
		case UntypedComplex:
			re := big.NewFloat(0)
			im := big.NewFloat(0)
			switch y := y.(type) {
			case UntypedInt:
				yre := big.NewFloat(0)
				yre.SetInt(y.Int)
				return UntypedComplex{re.Sub(x.Real, yre), im}, nil
			case UntypedFloat:
				yre := big.NewFloat(0)
				yre.Set(y.Float)
				return UntypedComplex{re.Sub(x.Real, yre), im}, nil
			case UntypedComplex:
				return UntypedComplex{
					re.Sub(x.Real, y.Real),
					im.Sub(x.Imag, y.Imag),
				}, nil
			}
		}
	case token.Mul:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x * y, nil
			}
		case int8:
			switch y := y.(type) {
			case int8:
				return x * y, nil
			}
		case int16:
			switch y := y.(type) {
			case int16:
				return x * y, nil
			}
		case int32:
			switch y := y.(type) {
			case int32:
				return x * y, nil
			}
		case int64:
			switch y := y.(type) {
			case int64:
				return x * y, nil
			}
		case uint:
			switch y := y.(type) {
			case uint:
				return x * y, nil
			}
		case uint8:
			switch y := y.(type) {
			case uint8:
				return x * y, nil
			}
		case uint16:
			switch y := y.(type) {
			case uint16:
				return x * y, nil
			}
		case uint32:
			switch y := y.(type) {
			case uint32:
				return x * y, nil
			}
		case uint64:
			switch y := y.(type) {
			case uint64:
				return x * y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x * y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x * y, nil
			}
		case complex64:
			switch y := y.(type) {
			case complex64:
				return x * y, nil
			}
		case complex128:
			switch y := y.(type) {
			case complex128:
				return x * y, nil
			}
		case *big.Int:
			switch y := y.(type) {
			case *big.Int:
				z := big.NewInt(0)
				return z.Mul(x, y), nil
			}
		case *big.Float:
			switch y := y.(type) {
			case *big.Float:
				z := big.NewFloat(0)
				return z.Mul(x, y), nil
			}
		case UntypedInt:
			switch y := y.(type) {
			case UntypedFloat:
				z := big.NewFloat(0)
				xf := big.NewFloat(float64(x.Int.Int64()))
				return UntypedFloat{z.Mul(xf, y.Float)}, nil
			case UntypedInt:
				z := big.NewInt(0)
				return UntypedInt{z.Mul(x.Int, y.Int)}, nil
			case UntypedComplex:
				re := big.NewFloat(0)
				xf := big.NewFloat(float64(x.Int.Int64()))
				im := big.NewFloat(0)
				return UntypedComplex{re.Mul(xf, y.Real), im.Mul(im, y.Imag)}, nil
			case complex64:
				return complex(float32(x.Int64()), 0) * y, nil
			case complex128:
				return complex(float64(x.Int64()), 0) * y, nil
			}
		case UntypedComplex:
			re := big.NewFloat(0)
			im := big.NewFloat(0)
			switch y := y.(type) {
			case UntypedInt:
				yre := big.NewFloat(0)
				yre.SetInt(y.Int)
				return UntypedComplex{re.Mul(x.Real, yre), im}, nil
			case UntypedFloat:
				yre := big.NewFloat(0)
				yre.Set(y.Float)
				return UntypedComplex{re.Mul(x.Real, yre), im}, nil
			case UntypedComplex:
				return UntypedComplex{
					re.Sub(x.Real, y.Real),
					im.Sub(x.Imag, y.Imag),
				}, nil
			}
		}
	case token.Div:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x / y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x / y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x / y, nil
			}
		case complex64:
			switch y := y.(type) {
			case complex64:
				return x / y, nil
			}
		case complex128:
			switch y := y.(type) {
			case complex128:
				return x / y, nil
			}
		}
	case token.Rem:
	case token.Pow:
	case token.LogicalAnd, token.LogicalOr:
		panic("logical ops processed before binOp")
	case token.Equal:
		return valEq(x, y), nil
	case token.NotEqual:
		return !valEq(x, y), nil
	case token.Less:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x < y, nil
			}
		case int8:
			switch y := y.(type) {
			case int8:
				return x < y, nil
			}
		case int16:
			switch y := y.(type) {
			case int16:
				return x < y, nil
			}
		case int32:
			switch y := y.(type) {
			case int32:
				return x < y, nil
			}
		case int64:
			switch y := y.(type) {
			case int64:
				return x < y, nil
			}
		case uint:
			switch y := y.(type) {
			case uint:
				return x < y, nil
			}
		case uint8:
			switch y := y.(type) {
			case uint8:
				return x < y, nil
			}
		case uint16:
			switch y := y.(type) {
			case uint16:
				return x < y, nil
			}
		case uint32:
			switch y := y.(type) {
			case uint32:
				return x < y, nil
			}
		case uint64:
			switch y := y.(type) {
			case uint64:
				return x < y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x < y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x < y, nil
			}
		case *big.Int:
			switch y := y.(type) {
			case *big.Int:
				return x.Cmp(y) == -1, nil
			}
		case *big.Float:
			switch y := y.(type) {
			case *big.Float:
				return x.Cmp(y) == -1, nil
			}
		}
	case token.Greater:
		switch x := x.(type) {
		case int:
			switch y := y.(type) {
			case int:
				return x > y, nil
			}
		case int8:
			switch y := y.(type) {
			case int8:
				return x > y, nil
			}
		case int16:
			switch y := y.(type) {
			case int16:
				return x > y, nil
			}
		case int32:
			switch y := y.(type) {
			case int32:
				return x > y, nil
			}
		case int64:
			switch y := y.(type) {
			case int64:
				return x > y, nil
			}
		case uint:
			switch y := y.(type) {
			case uint:
				return x > y, nil
			}
		case uint8:
			switch y := y.(type) {
			case uint8:
				return x > y, nil
			}
		case uint16:
			switch y := y.(type) {
			case uint16:
				return x > y, nil
			}
		case uint32:
			switch y := y.(type) {
			case uint32:
				return x > y, nil
			}
		case uint64:
			switch y := y.(type) {
			case uint64:
				return x > y, nil
			}
		case float32:
			switch y := y.(type) {
			case float32:
				return x > y, nil
			}
		case float64:
			switch y := y.(type) {
			case float64:
				return x > y, nil
			}
		case *big.Int:
			switch y := y.(type) {
			case *big.Int:
				return x.Cmp(y) == 1, nil
			}
		case *big.Float:
			switch y := y.(type) {
			case *big.Float:
				return x.Cmp(y) == 1, nil
			}
		}
	}
	//return nil, fmt.Errorf("type mismatch Left: %T, Right: %T", x, y)
	panic(fmt.Sprintf("binOp type mismatch Left: %+v (%T), Right: %+v (%T) op: %v", x, x, y, y, op))
}

func typeConv(t reflect.Type, v reflect.Value) (res reflect.Value) {
	if v.Type() == t {
		return v
	}
	if v.Kind() == t.Kind() {
		// named type conversion
		res = reflect.New(t).Elem()
		switch v.Kind() {
		case reflect.Bool:
			res.SetBool(v.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			res.SetInt(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			res.SetUint(v.Uint())
		case reflect.Float32, reflect.Float64:
			res.SetFloat(v.Float())
		case reflect.Complex64, reflect.Complex128:
			res.SetComplex(v.Complex())
		case reflect.String:
			res.SetString(v.String())
		case reflect.Chan:
			res.Set(v)
		default:
			panic(interpPanic{fmt.Errorf("TODO typeConv same kind: %v", v.Kind())})
		}
		return res
	}
	switch t.Kind() {
	case reflect.Int:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(int(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(int(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(int(v.Float()))
		}
	case reflect.Int8:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(int8(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(int8(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(int8(v.Float()))
		}
	case reflect.Int16:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(int16(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(int16(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(int16(v.Float()))
		}
	case reflect.Int32:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(int32(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(int32(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(int32(v.Float()))
		}
	case reflect.Int64:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(int64(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(int64(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(int64(v.Float()))
		}
	case reflect.Uint:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(uint(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(uint(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(uint(v.Float()))
		}
	case reflect.Uint8:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(uint8(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(uint8(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(uint8(v.Float()))
		}
	case reflect.Uint16:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(uint16(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(uint16(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(uint16(v.Float()))
		}
	case reflect.Uint32:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(uint32(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(uint32(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(uint32(v.Float()))
		}
	case reflect.Uint64:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(uint64(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return reflect.ValueOf(uint64(v.Uint()))
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(uint64(v.Float()))
		}
	case reflect.Float64:
		if v.Kind() == reflect.Float64 {
			res = reflect.New(t).Elem()
			res.SetFloat(v.Float())
			return res
		}
		return reflect.ValueOf(float64(v.Int()))
	case reflect.Complex64:
		switch v.Kind() {
		case reflect.Complex64, reflect.Complex128:
			res = reflect.New(t).Elem()
			res.SetComplex(v.Complex())
			return res
		}
		return reflect.ValueOf(complex(float32(v.Int()), 0))
	case reflect.Complex128:
		switch v.Kind() {
		case reflect.Complex64, reflect.Complex128:
			res = reflect.New(t).Elem()
			res.SetComplex(v.Complex())
			return res
		}
		return reflect.ValueOf(complex(float64(v.Int()), 0))
	case reflect.Interface:
		return reflect.ValueOf(v.Interface())
	case reflect.String:
		switch src := v.Interface().(type) {
		case []byte:
			return reflect.ValueOf(string(src))
		case rune:
			return reflect.ValueOf(string(src))
		}
	}
	if t == reflect.TypeOf([]byte(nil)) {
		switch src := v.Interface().(type) {
		case string:
			return reflect.ValueOf([]byte(src))
		}
	}
	panic(interpPanic{fmt.Errorf("unknown type conv: %v <- %v", t, v.Type())})
}
