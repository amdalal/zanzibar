// Code generated by thriftrw v1.18.0. DO NOT EDIT.
// @generated

package bar

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Bar_ArgWithManyQueryParams_Args represents the arguments for the Bar.argWithManyQueryParams function.
//
// The arguments for argWithManyQueryParams are sent and received over the wire as this struct.
type Bar_ArgWithManyQueryParams_Args struct {
	AStr            string     `json:"aStr,required"`
	AnOptStr        *string    `json:"anOptStr,omitempty"`
	ABool           bool       `json:"aBool,required"`
	AnOptBool       *bool      `json:"anOptBool,omitempty"`
	AInt8           int8       `json:"aInt8,required"`
	AnOptInt8       *int8      `json:"anOptInt8,omitempty"`
	AInt16          int16      `json:"aInt16,required"`
	AnOptInt16      *int16     `json:"anOptInt16,omitempty"`
	AInt32          int32      `json:"aInt32,required"`
	AnOptInt32      *int32     `json:"anOptInt32,omitempty"`
	AInt64          int64      `json:"aInt64,required"`
	AnOptInt64      *int64     `json:"anOptInt64,omitempty"`
	AFloat64        float64    `json:"aFloat64,required"`
	AnOptFloat64    *float64   `json:"anOptFloat64,omitempty"`
	AUUID           UUID       `json:"aUUID,required"`
	AnOptUUID       *UUID      `json:"anOptUUID,omitempty"`
	AListUUID       []UUID     `json:"aListUUID,required"`
	AnOptListUUID   []UUID     `json:"anOptListUUID,omitempty"`
	AStringList     StringList `json:"aStringList,required"`
	AnOptStringList StringList `json:"anOptStringList,omitempty"`
	AUUIDList       UUIDList   `json:"aUUIDList,required"`
	AnOptUUIDList   UUIDList   `json:"anOptUUIDList,omitempty"`
	ATs             Timestamp  `json:"aTs,required"`
	AnOptTs         *Timestamp `json:"anOptTs,omitempty"`
}

// ToWire translates a Bar_ArgWithManyQueryParams_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_ArgWithManyQueryParams_Args) ToWire() (wire.Value, error) {
	var (
		fields [24]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.AStr), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.AnOptStr != nil {
		w, err = wire.NewValueString(*(v.AnOptStr)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	w, err = wire.NewValueBool(v.ABool), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.AnOptBool != nil {
		w, err = wire.NewValueBool(*(v.AnOptBool)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	w, err = wire.NewValueI8(v.AInt8), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	if v.AnOptInt8 != nil {
		w, err = wire.NewValueI8(*(v.AnOptInt8)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}

	w, err = wire.NewValueI16(v.AInt16), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 7, Value: w}
	i++
	if v.AnOptInt16 != nil {
		w, err = wire.NewValueI16(*(v.AnOptInt16)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 8, Value: w}
		i++
	}

	w, err = wire.NewValueI32(v.AInt32), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 9, Value: w}
	i++
	if v.AnOptInt32 != nil {
		w, err = wire.NewValueI32(*(v.AnOptInt32)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}

	w, err = wire.NewValueI64(v.AInt64), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 11, Value: w}
	i++
	if v.AnOptInt64 != nil {
		w, err = wire.NewValueI64(*(v.AnOptInt64)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 12, Value: w}
		i++
	}

	w, err = wire.NewValueDouble(v.AFloat64), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 13, Value: w}
	i++
	if v.AnOptFloat64 != nil {
		w, err = wire.NewValueDouble(*(v.AnOptFloat64)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 14, Value: w}
		i++
	}

	w, err = v.AUUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 15, Value: w}
	i++
	if v.AnOptUUID != nil {
		w, err = v.AnOptUUID.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 16, Value: w}
		i++
	}
	if v.AListUUID == nil {
		return w, errors.New("field AListUUID of Bar_ArgWithManyQueryParams_Args is required")
	}
	w, err = wire.NewValueList(_List_UUID_ValueList(v.AListUUID)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 17, Value: w}
	i++
	if v.AnOptListUUID != nil {
		w, err = wire.NewValueList(_List_UUID_ValueList(v.AnOptListUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 18, Value: w}
		i++
	}
	if v.AStringList == nil {
		return w, errors.New("field AStringList of Bar_ArgWithManyQueryParams_Args is required")
	}
	w, err = v.AStringList.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 19, Value: w}
	i++
	if v.AnOptStringList != nil {
		w, err = v.AnOptStringList.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.AUUIDList == nil {
		return w, errors.New("field AUUIDList of Bar_ArgWithManyQueryParams_Args is required")
	}
	w, err = v.AUUIDList.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 21, Value: w}
	i++
	if v.AnOptUUIDList != nil {
		w, err = v.AnOptUUIDList.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 22, Value: w}
		i++
	}

	w, err = v.ATs.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 23, Value: w}
	i++
	if v.AnOptTs != nil {
		w, err = v.AnOptTs.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 24, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _StringList_Read(w wire.Value) (StringList, error) {
	var x StringList
	err := x.FromWire(w)
	return x, err
}

func _UUIDList_Read(w wire.Value) (UUIDList, error) {
	var x UUIDList
	err := x.FromWire(w)
	return x, err
}

// FromWire deserializes a Bar_ArgWithManyQueryParams_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithManyQueryParams_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithManyQueryParams_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithManyQueryParams_Args) FromWire(w wire.Value) error {
	var err error

	aStrIsSet := false

	aBoolIsSet := false

	aInt8IsSet := false

	aInt16IsSet := false

	aInt32IsSet := false

	aInt64IsSet := false

	aFloat64IsSet := false

	aUUIDIsSet := false

	aListUUIDIsSet := false

	aStringListIsSet := false

	aUUIDListIsSet := false

	aTsIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.AStr, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				aStrIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.AnOptStr = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TBool {
				v.ABool, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				aBoolIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.AnOptBool = &x
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TI8 {
				v.AInt8, err = field.Value.GetI8(), error(nil)
				if err != nil {
					return err
				}
				aInt8IsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TI8 {
				var x int8
				x, err = field.Value.GetI8(), error(nil)
				v.AnOptInt8 = &x
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TI16 {
				v.AInt16, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				aInt16IsSet = true
			}
		case 8:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
				v.AnOptInt16 = &x
				if err != nil {
					return err
				}

			}
		case 9:
			if field.Value.Type() == wire.TI32 {
				v.AInt32, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				aInt32IsSet = true
			}
		case 10:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.AnOptInt32 = &x
				if err != nil {
					return err
				}

			}
		case 11:
			if field.Value.Type() == wire.TI64 {
				v.AInt64, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
				aInt64IsSet = true
			}
		case 12:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.AnOptInt64 = &x
				if err != nil {
					return err
				}

			}
		case 13:
			if field.Value.Type() == wire.TDouble {
				v.AFloat64, err = field.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
				aFloat64IsSet = true
			}
		case 14:
			if field.Value.Type() == wire.TDouble {
				var x float64
				x, err = field.Value.GetDouble(), error(nil)
				v.AnOptFloat64 = &x
				if err != nil {
					return err
				}

			}
		case 15:
			if field.Value.Type() == wire.TBinary {
				v.AUUID, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				aUUIDIsSet = true
			}
		case 16:
			if field.Value.Type() == wire.TBinary {
				var x UUID
				x, err = _UUID_Read(field.Value)
				v.AnOptUUID = &x
				if err != nil {
					return err
				}

			}
		case 17:
			if field.Value.Type() == wire.TList {
				v.AListUUID, err = _List_UUID_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				aListUUIDIsSet = true
			}
		case 18:
			if field.Value.Type() == wire.TList {
				v.AnOptListUUID, err = _List_UUID_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 19:
			if field.Value.Type() == wire.TList {
				v.AStringList, err = _StringList_Read(field.Value)
				if err != nil {
					return err
				}
				aStringListIsSet = true
			}
		case 20:
			if field.Value.Type() == wire.TList {
				v.AnOptStringList, err = _StringList_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 21:
			if field.Value.Type() == wire.TList {
				v.AUUIDList, err = _UUIDList_Read(field.Value)
				if err != nil {
					return err
				}
				aUUIDListIsSet = true
			}
		case 22:
			if field.Value.Type() == wire.TList {
				v.AnOptUUIDList, err = _UUIDList_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 23:
			if field.Value.Type() == wire.TI64 {
				v.ATs, err = _Timestamp_Read(field.Value)
				if err != nil {
					return err
				}
				aTsIsSet = true
			}
		case 24:
			if field.Value.Type() == wire.TI64 {
				var x Timestamp
				x, err = _Timestamp_Read(field.Value)
				v.AnOptTs = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !aStrIsSet {
		return errors.New("field AStr of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aBoolIsSet {
		return errors.New("field ABool of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aInt8IsSet {
		return errors.New("field AInt8 of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aInt16IsSet {
		return errors.New("field AInt16 of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aInt32IsSet {
		return errors.New("field AInt32 of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aInt64IsSet {
		return errors.New("field AInt64 of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aFloat64IsSet {
		return errors.New("field AFloat64 of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aUUIDIsSet {
		return errors.New("field AUUID of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aListUUIDIsSet {
		return errors.New("field AListUUID of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aStringListIsSet {
		return errors.New("field AStringList of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aUUIDListIsSet {
		return errors.New("field AUUIDList of Bar_ArgWithManyQueryParams_Args is required")
	}

	if !aTsIsSet {
		return errors.New("field ATs of Bar_ArgWithManyQueryParams_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithManyQueryParams_Args
// struct.
func (v *Bar_ArgWithManyQueryParams_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [24]string
	i := 0
	fields[i] = fmt.Sprintf("AStr: %v", v.AStr)
	i++
	if v.AnOptStr != nil {
		fields[i] = fmt.Sprintf("AnOptStr: %v", *(v.AnOptStr))
		i++
	}
	fields[i] = fmt.Sprintf("ABool: %v", v.ABool)
	i++
	if v.AnOptBool != nil {
		fields[i] = fmt.Sprintf("AnOptBool: %v", *(v.AnOptBool))
		i++
	}
	fields[i] = fmt.Sprintf("AInt8: %v", v.AInt8)
	i++
	if v.AnOptInt8 != nil {
		fields[i] = fmt.Sprintf("AnOptInt8: %v", *(v.AnOptInt8))
		i++
	}
	fields[i] = fmt.Sprintf("AInt16: %v", v.AInt16)
	i++
	if v.AnOptInt16 != nil {
		fields[i] = fmt.Sprintf("AnOptInt16: %v", *(v.AnOptInt16))
		i++
	}
	fields[i] = fmt.Sprintf("AInt32: %v", v.AInt32)
	i++
	if v.AnOptInt32 != nil {
		fields[i] = fmt.Sprintf("AnOptInt32: %v", *(v.AnOptInt32))
		i++
	}
	fields[i] = fmt.Sprintf("AInt64: %v", v.AInt64)
	i++
	if v.AnOptInt64 != nil {
		fields[i] = fmt.Sprintf("AnOptInt64: %v", *(v.AnOptInt64))
		i++
	}
	fields[i] = fmt.Sprintf("AFloat64: %v", v.AFloat64)
	i++
	if v.AnOptFloat64 != nil {
		fields[i] = fmt.Sprintf("AnOptFloat64: %v", *(v.AnOptFloat64))
		i++
	}
	fields[i] = fmt.Sprintf("AUUID: %v", v.AUUID)
	i++
	if v.AnOptUUID != nil {
		fields[i] = fmt.Sprintf("AnOptUUID: %v", *(v.AnOptUUID))
		i++
	}
	fields[i] = fmt.Sprintf("AListUUID: %v", v.AListUUID)
	i++
	if v.AnOptListUUID != nil {
		fields[i] = fmt.Sprintf("AnOptListUUID: %v", v.AnOptListUUID)
		i++
	}
	fields[i] = fmt.Sprintf("AStringList: %v", v.AStringList)
	i++
	if v.AnOptStringList != nil {
		fields[i] = fmt.Sprintf("AnOptStringList: %v", v.AnOptStringList)
		i++
	}
	fields[i] = fmt.Sprintf("AUUIDList: %v", v.AUUIDList)
	i++
	if v.AnOptUUIDList != nil {
		fields[i] = fmt.Sprintf("AnOptUUIDList: %v", v.AnOptUUIDList)
		i++
	}
	fields[i] = fmt.Sprintf("ATs: %v", v.ATs)
	i++
	if v.AnOptTs != nil {
		fields[i] = fmt.Sprintf("AnOptTs: %v", *(v.AnOptTs))
		i++
	}

	return fmt.Sprintf("Bar_ArgWithManyQueryParams_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Byte_EqualsPtr(lhs, rhs *int8) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I16_EqualsPtr(lhs, rhs *int16) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Double_EqualsPtr(lhs, rhs *float64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _UUID_EqualsPtr(lhs, rhs *UUID) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Timestamp_EqualsPtr(lhs, rhs *Timestamp) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Bar_ArgWithManyQueryParams_Args match the
// provided Bar_ArgWithManyQueryParams_Args.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithManyQueryParams_Args) Equals(rhs *Bar_ArgWithManyQueryParams_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.AStr == rhs.AStr) {
		return false
	}
	if !_String_EqualsPtr(v.AnOptStr, rhs.AnOptStr) {
		return false
	}
	if !(v.ABool == rhs.ABool) {
		return false
	}
	if !_Bool_EqualsPtr(v.AnOptBool, rhs.AnOptBool) {
		return false
	}
	if !(v.AInt8 == rhs.AInt8) {
		return false
	}
	if !_Byte_EqualsPtr(v.AnOptInt8, rhs.AnOptInt8) {
		return false
	}
	if !(v.AInt16 == rhs.AInt16) {
		return false
	}
	if !_I16_EqualsPtr(v.AnOptInt16, rhs.AnOptInt16) {
		return false
	}
	if !(v.AInt32 == rhs.AInt32) {
		return false
	}
	if !_I32_EqualsPtr(v.AnOptInt32, rhs.AnOptInt32) {
		return false
	}
	if !(v.AInt64 == rhs.AInt64) {
		return false
	}
	if !_I64_EqualsPtr(v.AnOptInt64, rhs.AnOptInt64) {
		return false
	}
	if !(v.AFloat64 == rhs.AFloat64) {
		return false
	}
	if !_Double_EqualsPtr(v.AnOptFloat64, rhs.AnOptFloat64) {
		return false
	}
	if !(v.AUUID == rhs.AUUID) {
		return false
	}
	if !_UUID_EqualsPtr(v.AnOptUUID, rhs.AnOptUUID) {
		return false
	}
	if !_List_UUID_Equals(v.AListUUID, rhs.AListUUID) {
		return false
	}
	if !((v.AnOptListUUID == nil && rhs.AnOptListUUID == nil) || (v.AnOptListUUID != nil && rhs.AnOptListUUID != nil && _List_UUID_Equals(v.AnOptListUUID, rhs.AnOptListUUID))) {
		return false
	}
	if !v.AStringList.Equals(rhs.AStringList) {
		return false
	}
	if !((v.AnOptStringList == nil && rhs.AnOptStringList == nil) || (v.AnOptStringList != nil && rhs.AnOptStringList != nil && v.AnOptStringList.Equals(rhs.AnOptStringList))) {
		return false
	}
	if !v.AUUIDList.Equals(rhs.AUUIDList) {
		return false
	}
	if !((v.AnOptUUIDList == nil && rhs.AnOptUUIDList == nil) || (v.AnOptUUIDList != nil && rhs.AnOptUUIDList != nil && v.AnOptUUIDList.Equals(rhs.AnOptUUIDList))) {
		return false
	}
	if !(v.ATs == rhs.ATs) {
		return false
	}
	if !_Timestamp_EqualsPtr(v.AnOptTs, rhs.AnOptTs) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithManyQueryParams_Args.
func (v *Bar_ArgWithManyQueryParams_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("aStr", v.AStr)
	if v.AnOptStr != nil {
		enc.AddString("anOptStr", *v.AnOptStr)
	}
	enc.AddBool("aBool", v.ABool)
	if v.AnOptBool != nil {
		enc.AddBool("anOptBool", *v.AnOptBool)
	}
	enc.AddInt8("aInt8", v.AInt8)
	if v.AnOptInt8 != nil {
		enc.AddInt8("anOptInt8", *v.AnOptInt8)
	}
	enc.AddInt16("aInt16", v.AInt16)
	if v.AnOptInt16 != nil {
		enc.AddInt16("anOptInt16", *v.AnOptInt16)
	}
	enc.AddInt32("aInt32", v.AInt32)
	if v.AnOptInt32 != nil {
		enc.AddInt32("anOptInt32", *v.AnOptInt32)
	}
	enc.AddInt64("aInt64", v.AInt64)
	if v.AnOptInt64 != nil {
		enc.AddInt64("anOptInt64", *v.AnOptInt64)
	}
	enc.AddFloat64("aFloat64", v.AFloat64)
	if v.AnOptFloat64 != nil {
		enc.AddFloat64("anOptFloat64", *v.AnOptFloat64)
	}
	enc.AddString("aUUID", (string)(v.AUUID))
	if v.AnOptUUID != nil {
		enc.AddString("anOptUUID", (string)(*v.AnOptUUID))
	}
	err = multierr.Append(err, enc.AddArray("aListUUID", (_List_UUID_Zapper)(v.AListUUID)))
	if v.AnOptListUUID != nil {
		err = multierr.Append(err, enc.AddArray("anOptListUUID", (_List_UUID_Zapper)(v.AnOptListUUID)))
	}
	err = multierr.Append(err, enc.AddArray("aStringList", (_List_String_Zapper)(([]string)(v.AStringList))))
	if v.AnOptStringList != nil {
		err = multierr.Append(err, enc.AddArray("anOptStringList", (_List_String_Zapper)(([]string)(v.AnOptStringList))))
	}
	err = multierr.Append(err, enc.AddArray("aUUIDList", (_List_UUID_Zapper)(([]UUID)(v.AUUIDList))))
	if v.AnOptUUIDList != nil {
		err = multierr.Append(err, enc.AddArray("anOptUUIDList", (_List_UUID_Zapper)(([]UUID)(v.AnOptUUIDList))))
	}
	enc.AddInt64("aTs", (int64)(v.ATs))
	if v.AnOptTs != nil {
		enc.AddInt64("anOptTs", (int64)(*v.AnOptTs))
	}
	return err
}

// GetAStr returns the value of AStr if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAStr() (o string) {
	if v != nil {
		o = v.AStr
	}
	return
}

// GetAnOptStr returns the value of AnOptStr if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptStr() (o string) {
	if v != nil && v.AnOptStr != nil {
		return *v.AnOptStr
	}

	return
}

// IsSetAnOptStr returns true if AnOptStr is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptStr() bool {
	return v != nil && v.AnOptStr != nil
}

// GetABool returns the value of ABool if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetABool() (o bool) {
	if v != nil {
		o = v.ABool
	}
	return
}

// GetAnOptBool returns the value of AnOptBool if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptBool() (o bool) {
	if v != nil && v.AnOptBool != nil {
		return *v.AnOptBool
	}

	return
}

// IsSetAnOptBool returns true if AnOptBool is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptBool() bool {
	return v != nil && v.AnOptBool != nil
}

// GetAInt8 returns the value of AInt8 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAInt8() (o int8) {
	if v != nil {
		o = v.AInt8
	}
	return
}

// GetAnOptInt8 returns the value of AnOptInt8 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptInt8() (o int8) {
	if v != nil && v.AnOptInt8 != nil {
		return *v.AnOptInt8
	}

	return
}

// IsSetAnOptInt8 returns true if AnOptInt8 is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptInt8() bool {
	return v != nil && v.AnOptInt8 != nil
}

// GetAInt16 returns the value of AInt16 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAInt16() (o int16) {
	if v != nil {
		o = v.AInt16
	}
	return
}

// GetAnOptInt16 returns the value of AnOptInt16 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptInt16() (o int16) {
	if v != nil && v.AnOptInt16 != nil {
		return *v.AnOptInt16
	}

	return
}

// IsSetAnOptInt16 returns true if AnOptInt16 is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptInt16() bool {
	return v != nil && v.AnOptInt16 != nil
}

// GetAInt32 returns the value of AInt32 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAInt32() (o int32) {
	if v != nil {
		o = v.AInt32
	}
	return
}

// GetAnOptInt32 returns the value of AnOptInt32 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptInt32() (o int32) {
	if v != nil && v.AnOptInt32 != nil {
		return *v.AnOptInt32
	}

	return
}

// IsSetAnOptInt32 returns true if AnOptInt32 is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptInt32() bool {
	return v != nil && v.AnOptInt32 != nil
}

// GetAInt64 returns the value of AInt64 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAInt64() (o int64) {
	if v != nil {
		o = v.AInt64
	}
	return
}

// GetAnOptInt64 returns the value of AnOptInt64 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptInt64() (o int64) {
	if v != nil && v.AnOptInt64 != nil {
		return *v.AnOptInt64
	}

	return
}

// IsSetAnOptInt64 returns true if AnOptInt64 is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptInt64() bool {
	return v != nil && v.AnOptInt64 != nil
}

// GetAFloat64 returns the value of AFloat64 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAFloat64() (o float64) {
	if v != nil {
		o = v.AFloat64
	}
	return
}

// GetAnOptFloat64 returns the value of AnOptFloat64 if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptFloat64() (o float64) {
	if v != nil && v.AnOptFloat64 != nil {
		return *v.AnOptFloat64
	}

	return
}

// IsSetAnOptFloat64 returns true if AnOptFloat64 is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptFloat64() bool {
	return v != nil && v.AnOptFloat64 != nil
}

// GetAUUID returns the value of AUUID if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAUUID() (o UUID) {
	if v != nil {
		o = v.AUUID
	}
	return
}

// GetAnOptUUID returns the value of AnOptUUID if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptUUID() (o UUID) {
	if v != nil && v.AnOptUUID != nil {
		return *v.AnOptUUID
	}

	return
}

// IsSetAnOptUUID returns true if AnOptUUID is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptUUID() bool {
	return v != nil && v.AnOptUUID != nil
}

// GetAListUUID returns the value of AListUUID if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAListUUID() (o []UUID) {
	if v != nil {
		o = v.AListUUID
	}
	return
}

// IsSetAListUUID returns true if AListUUID is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAListUUID() bool {
	return v != nil && v.AListUUID != nil
}

// GetAnOptListUUID returns the value of AnOptListUUID if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptListUUID() (o []UUID) {
	if v != nil && v.AnOptListUUID != nil {
		return v.AnOptListUUID
	}

	return
}

// IsSetAnOptListUUID returns true if AnOptListUUID is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptListUUID() bool {
	return v != nil && v.AnOptListUUID != nil
}

// GetAStringList returns the value of AStringList if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAStringList() (o StringList) {
	if v != nil {
		o = v.AStringList
	}
	return
}

// IsSetAStringList returns true if AStringList is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAStringList() bool {
	return v != nil && v.AStringList != nil
}

// GetAnOptStringList returns the value of AnOptStringList if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptStringList() (o StringList) {
	if v != nil && v.AnOptStringList != nil {
		return v.AnOptStringList
	}

	return
}

// IsSetAnOptStringList returns true if AnOptStringList is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptStringList() bool {
	return v != nil && v.AnOptStringList != nil
}

// GetAUUIDList returns the value of AUUIDList if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAUUIDList() (o UUIDList) {
	if v != nil {
		o = v.AUUIDList
	}
	return
}

// IsSetAUUIDList returns true if AUUIDList is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAUUIDList() bool {
	return v != nil && v.AUUIDList != nil
}

// GetAnOptUUIDList returns the value of AnOptUUIDList if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptUUIDList() (o UUIDList) {
	if v != nil && v.AnOptUUIDList != nil {
		return v.AnOptUUIDList
	}

	return
}

// IsSetAnOptUUIDList returns true if AnOptUUIDList is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptUUIDList() bool {
	return v != nil && v.AnOptUUIDList != nil
}

// GetATs returns the value of ATs if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetATs() (o Timestamp) {
	if v != nil {
		o = v.ATs
	}
	return
}

// GetAnOptTs returns the value of AnOptTs if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Args) GetAnOptTs() (o Timestamp) {
	if v != nil && v.AnOptTs != nil {
		return *v.AnOptTs
	}

	return
}

// IsSetAnOptTs returns true if AnOptTs is not nil.
func (v *Bar_ArgWithManyQueryParams_Args) IsSetAnOptTs() bool {
	return v != nil && v.AnOptTs != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "argWithManyQueryParams" for this struct.
func (v *Bar_ArgWithManyQueryParams_Args) MethodName() string {
	return "argWithManyQueryParams"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_ArgWithManyQueryParams_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_ArgWithManyQueryParams_Helper provides functions that aid in handling the
// parameters and return values of the Bar.argWithManyQueryParams
// function.
var Bar_ArgWithManyQueryParams_Helper = struct {
	// Args accepts the parameters of argWithManyQueryParams in-order and returns
	// the arguments struct for the function.
	Args func(
		aStr string,
		anOptStr *string,
		aBool bool,
		anOptBool *bool,
		aInt8 int8,
		anOptInt8 *int8,
		aInt16 int16,
		anOptInt16 *int16,
		aInt32 int32,
		anOptInt32 *int32,
		aInt64 int64,
		anOptInt64 *int64,
		aFloat64 float64,
		anOptFloat64 *float64,
		aUUID UUID,
		anOptUUID *UUID,
		aListUUID []UUID,
		anOptListUUID []UUID,
		aStringList StringList,
		anOptStringList StringList,
		aUUIDList UUIDList,
		anOptUUIDList UUIDList,
		aTs Timestamp,
		anOptTs *Timestamp,
	) *Bar_ArgWithManyQueryParams_Args

	// IsException returns true if the given error can be thrown
	// by argWithManyQueryParams.
	//
	// An error can be thrown by argWithManyQueryParams only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for argWithManyQueryParams
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// argWithManyQueryParams into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by argWithManyQueryParams
	//
	//   value, err := argWithManyQueryParams(args)
	//   result, err := Bar_ArgWithManyQueryParams_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from argWithManyQueryParams: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponse, error) (*Bar_ArgWithManyQueryParams_Result, error)

	// UnwrapResponse takes the result struct for argWithManyQueryParams
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if argWithManyQueryParams threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_ArgWithManyQueryParams_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_ArgWithManyQueryParams_Result) (*BarResponse, error)
}{}

func init() {
	Bar_ArgWithManyQueryParams_Helper.Args = func(
		aStr string,
		anOptStr *string,
		aBool bool,
		anOptBool *bool,
		aInt8 int8,
		anOptInt8 *int8,
		aInt16 int16,
		anOptInt16 *int16,
		aInt32 int32,
		anOptInt32 *int32,
		aInt64 int64,
		anOptInt64 *int64,
		aFloat64 float64,
		anOptFloat64 *float64,
		aUUID UUID,
		anOptUUID *UUID,
		aListUUID []UUID,
		anOptListUUID []UUID,
		aStringList StringList,
		anOptStringList StringList,
		aUUIDList UUIDList,
		anOptUUIDList UUIDList,
		aTs Timestamp,
		anOptTs *Timestamp,
	) *Bar_ArgWithManyQueryParams_Args {
		return &Bar_ArgWithManyQueryParams_Args{
			AStr:            aStr,
			AnOptStr:        anOptStr,
			ABool:           aBool,
			AnOptBool:       anOptBool,
			AInt8:           aInt8,
			AnOptInt8:       anOptInt8,
			AInt16:          aInt16,
			AnOptInt16:      anOptInt16,
			AInt32:          aInt32,
			AnOptInt32:      anOptInt32,
			AInt64:          aInt64,
			AnOptInt64:      anOptInt64,
			AFloat64:        aFloat64,
			AnOptFloat64:    anOptFloat64,
			AUUID:           aUUID,
			AnOptUUID:       anOptUUID,
			AListUUID:       aListUUID,
			AnOptListUUID:   anOptListUUID,
			AStringList:     aStringList,
			AnOptStringList: anOptStringList,
			AUUIDList:       aUUIDList,
			AnOptUUIDList:   anOptUUIDList,
			ATs:             aTs,
			AnOptTs:         anOptTs,
		}
	}

	Bar_ArgWithManyQueryParams_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Bar_ArgWithManyQueryParams_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_ArgWithManyQueryParams_Result, error) {
		if err == nil {
			return &Bar_ArgWithManyQueryParams_Result{Success: success}, nil
		}

		return nil, err
	}
	Bar_ArgWithManyQueryParams_Helper.UnwrapResponse = func(result *Bar_ArgWithManyQueryParams_Result) (success *BarResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_ArgWithManyQueryParams_Result represents the result of a Bar.argWithManyQueryParams function call.
//
// The result of a argWithManyQueryParams execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_ArgWithManyQueryParams_Result struct {
	// Value returned by argWithManyQueryParams after a successful execution.
	Success *BarResponse `json:"success,omitempty"`
}

// ToWire translates a Bar_ArgWithManyQueryParams_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_ArgWithManyQueryParams_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_ArgWithManyQueryParams_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ArgWithManyQueryParams_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithManyQueryParams_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithManyQueryParams_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithManyQueryParams_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_ArgWithManyQueryParams_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithManyQueryParams_Result
// struct.
func (v *Bar_ArgWithManyQueryParams_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Bar_ArgWithManyQueryParams_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgWithManyQueryParams_Result match the
// provided Bar_ArgWithManyQueryParams_Result.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithManyQueryParams_Result) Equals(rhs *Bar_ArgWithManyQueryParams_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithManyQueryParams_Result.
func (v *Bar_ArgWithManyQueryParams_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithManyQueryParams_Result) GetSuccess() (o *BarResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_ArgWithManyQueryParams_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "argWithManyQueryParams" for this struct.
func (v *Bar_ArgWithManyQueryParams_Result) MethodName() string {
	return "argWithManyQueryParams"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_ArgWithManyQueryParams_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
