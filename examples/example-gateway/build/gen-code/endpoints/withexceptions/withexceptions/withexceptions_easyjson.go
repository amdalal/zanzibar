// Code generated by zanzibar
// @generated
// Checksum : cPD08VfXf6ZVgSJ/oQpHyw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package withexceptions

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(in *jlexer.Lexer, out *WithExceptions_Func1_Result) {
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
					out.Success = new(Response)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		case "e1":
			if in.IsNull() {
				in.Skip()
				out.E1 = nil
			} else {
				if out.E1 == nil {
					out.E1 = new(EndpointExceptionType1)
				}
				(*out.E1).UnmarshalEasyJSON(in)
			}
		case "e2":
			if in.IsNull() {
				in.Skip()
				out.E2 = nil
			} else {
				if out.E2 == nil {
					out.E2 = new(EndpointExceptionType2)
				}
				(*out.E2).UnmarshalEasyJSON(in)
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
func easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(out *jwriter.Writer, in WithExceptions_Func1_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Success).MarshalEasyJSON(out)
	}
	if in.E1 != nil {
		const prefix string = ",\"e1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.E1).MarshalEasyJSON(out)
	}
	if in.E2 != nil {
		const prefix string = ",\"e2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.E2).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithExceptions_Func1_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WithExceptions_Func1_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithExceptions_Func1_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WithExceptions_Func1_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc1(l, v)
}
func easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(in *jlexer.Lexer, out *WithExceptions_Func1_Args) {
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
func easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(out *jwriter.Writer, in WithExceptions_Func1_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithExceptions_Func1_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WithExceptions_Func1_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithExceptions_Func1_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WithExceptions_Func1_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptionsWithExceptionsFunc11(l, v)
}
func easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(in *jlexer.Lexer, out *Response) {
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
func easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions(l, v)
}
func easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(in *jlexer.Lexer, out *EndpointExceptionType2) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var Message2Set bool
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
		case "message2":
			out.Message2 = string(in.String())
			Message2Set = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !Message2Set {
		in.AddError(fmt.Errorf("key 'message2' is required"))
	}
}
func easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(out *jwriter.Writer, in EndpointExceptionType2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message2\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message2))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointExceptionType2) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointExceptionType2) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointExceptionType2) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointExceptionType2) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions1(l, v)
}
func easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(in *jlexer.Lexer, out *EndpointExceptionType1) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var Message1Set bool
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
		case "message1":
			out.Message1 = string(in.String())
			Message1Set = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !Message1Set {
		in.AddError(fmt.Errorf("key 'message1' is required"))
	}
}
func easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(out *jwriter.Writer, in EndpointExceptionType1) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message1\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message1))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointExceptionType1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointExceptionType1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD396a41aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointExceptionType1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointExceptionType1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD396a41aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsWithexceptionsWithexceptions2(l, v)
}
