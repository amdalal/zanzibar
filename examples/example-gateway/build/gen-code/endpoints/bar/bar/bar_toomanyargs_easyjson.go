// Code generated by zanzibar
// @generated
// Checksum : 4/KpvSnYuEMTO4gmsqd7zA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/base/base"
	foo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/foo"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(in *jlexer.Lexer, out *Bar_TooManyArgs_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in, &*out.Success)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar1(in, &*out.BarException)
			}
		case "fooException":
			if in.IsNull() {
				in.Skip()
				out.FooException = nil
			} else {
				if out.FooException == nil {
					out.FooException = new(foo.FooException)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo(in, &*out.FooException)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(out *jwriter.Writer, in Bar_TooManyArgs_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out, *in.Success)
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar1(out, *in.BarException)
		}
	}
	if in.FooException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooException\":")
		if in.FooException == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo(out, *in.FooException)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs(l, v)
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo(in *jlexer.Lexer, out *foo.FooException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var TeapotSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "teapot":
			out.Teapot = string(in.String())
			TeapotSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !TeapotSet {
		in.AddError(fmt.Errorf("key 'teapot' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo(out *jwriter.Writer, in foo.FooException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"teapot\":")
	out.String(string(in.Teapot))
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar1(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar1(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
	var BinaryFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[UUID]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int32
					v2 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.MapIntWithRange {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Int32(int32(v4Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v5First := true
		for v5Name, v5Value := range in.MapIntWithoutRange {
			if !v5First {
				out.RawByte(',')
			}
			v5First = false
			out.String(string(v5Name))
			out.RawByte(':')
			out.Int32(int32(v5Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"binaryField\":")
	out.Base64Bytes(in.BinaryField)
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(in *jlexer.Lexer, out *Bar_TooManyArgs_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar2(in, &*out.Request)
			}
			RequestSet = true
		case "foo":
			if in.IsNull() {
				in.Skip()
				out.Foo = nil
			} else {
				if out.Foo == nil {
					out.Foo = new(foo.FooStruct)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo1(in, &*out.Foo)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(out *jwriter.Writer, in Bar_TooManyArgs_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar2(out, *in.Request)
	}
	if in.Foo != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"foo\":")
		if in.Foo == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo1(out, *in.Foo)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarTooManyArgs1(l, v)
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo1(in *jlexer.Lexer, out *foo.FooStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var FooStringSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fooString":
			out.FooString = string(in.String())
			FooStringSet = true
		case "fooI32":
			if in.IsNull() {
				in.Skip()
				out.FooI32 = nil
			} else {
				if out.FooI32 == nil {
					out.FooI32 = new(int32)
				}
				*out.FooI32 = int32(in.Int32())
			}
		case "fooI16":
			if in.IsNull() {
				in.Skip()
				out.FooI16 = nil
			} else {
				if out.FooI16 == nil {
					out.FooI16 = new(int16)
				}
				*out.FooI16 = int16(in.Int16())
			}
		case "fooDouble":
			if in.IsNull() {
				in.Skip()
				out.FooDouble = nil
			} else {
				if out.FooDouble == nil {
					out.FooDouble = new(float64)
				}
				*out.FooDouble = float64(in.Float64())
			}
		case "fooBool":
			if in.IsNull() {
				in.Skip()
				out.FooBool = nil
			} else {
				if out.FooBool == nil {
					out.FooBool = new(bool)
				}
				*out.FooBool = bool(in.Bool())
			}
		case "fooMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.FooMap = make(map[string]string)
				} else {
					out.FooMap = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 string
					v8 = string(in.String())
					(out.FooMap)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(base.Message)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooBaseBase(in, &*out.Message)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !FooStringSet {
		in.AddError(fmt.Errorf("key 'fooString' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooFoo1(out *jwriter.Writer, in foo.FooStruct) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fooString\":")
	out.String(string(in.FooString))
	if in.FooI32 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI32\":")
		if in.FooI32 == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.FooI32))
		}
	}
	if in.FooI16 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI16\":")
		if in.FooI16 == nil {
			out.RawString("null")
		} else {
			out.Int16(int16(*in.FooI16))
		}
	}
	if in.FooDouble != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooDouble\":")
		if in.FooDouble == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.FooDouble))
		}
	}
	if in.FooBool != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooBool\":")
		if in.FooBool == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.FooBool))
		}
	}
	if len(in.FooMap) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooMap\":")
		if in.FooMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.FooMap {
				if !v9First {
					out.RawByte(',')
				}
				v9First = false
				out.String(string(v9Name))
				out.RawByte(':')
				out.String(string(v9Value))
			}
			out.RawByte('}')
		}
	}
	if in.Message != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"message\":")
		if in.Message == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooBaseBase(out, *in.Message)
		}
	}
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooBaseBase(in *jlexer.Lexer, out *base.Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var BodySet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "body":
			out.Body = string(in.String())
			BodySet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !BodySet {
		in.AddError(fmt.Errorf("key 'body' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsFooBaseBase(out *jwriter.Writer, in base.Message) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"body\":")
	out.String(string(in.Body))
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar2(in *jlexer.Lexer, out *BarRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var BoolFieldSet bool
	var BinaryFieldSet bool
	var TimestampSet bool
	var EnumFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "boolField":
			out.BoolField = bool(in.Bool())
			BoolFieldSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
			TimestampSet = true
		case "enumField":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EnumField).UnmarshalJSON(data))
			}
			EnumFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !BoolFieldSet {
		in.AddError(fmt.Errorf("key 'boolField' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
	if !TimestampSet {
		in.AddError(fmt.Errorf("key 'timestamp' is required"))
	}
	if !EnumFieldSet {
		in.AddError(fmt.Errorf("key 'enumField' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar2(out *jwriter.Writer, in BarRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"boolField\":")
	out.Bool(bool(in.BoolField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"binaryField\":")
	out.Base64Bytes(in.BinaryField)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"timestamp\":")
	out.Raw((in.Timestamp).MarshalJSON())
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"enumField\":")
	out.Raw((in.EnumField).MarshalJSON())
	out.RawByte('}')
}
