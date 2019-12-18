// Code generated by zanzibar
// @generated
// Checksum : fjmocJQWZ9RtHD6c7hXZig==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package serverless

import (
	json "encoding/json"

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

func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(in *jlexer.Lexer, out *Serverless_ServerlessArgWithHeaders_Result) {
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(out *jwriter.Writer, in Serverless_ServerlessArgWithHeaders_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Success).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_ServerlessArgWithHeaders_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_ServerlessArgWithHeaders_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_ServerlessArgWithHeaders_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_ServerlessArgWithHeaders_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(in *jlexer.Lexer, out *Serverless_ServerlessArgWithHeaders_Args) {
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(out *jwriter.Writer, in Serverless_ServerlessArgWithHeaders_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_ServerlessArgWithHeaders_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_ServerlessArgWithHeaders_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_ServerlessArgWithHeaders_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_ServerlessArgWithHeaders_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessServerlessArgWithHeaders1(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(in *jlexer.Lexer, out *Serverless_EmptyServerlessRequest_Result) {
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(out *jwriter.Writer, in Serverless_EmptyServerlessRequest_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_EmptyServerlessRequest_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_EmptyServerlessRequest_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_EmptyServerlessRequest_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_EmptyServerlessRequest_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(in *jlexer.Lexer, out *Serverless_EmptyServerlessRequest_Args) {
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
		case "testString":
			if in.IsNull() {
				in.Skip()
				out.TestString = nil
			} else {
				if out.TestString == nil {
					out.TestString = new(string)
				}
				*out.TestString = string(in.String())
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(out *jwriter.Writer, in Serverless_EmptyServerlessRequest_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if in.TestString != nil {
		const prefix string = ",\"testString\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.TestString))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_EmptyServerlessRequest_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_EmptyServerlessRequest_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_EmptyServerlessRequest_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_EmptyServerlessRequest_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessEmptyServerlessRequest1(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(in *jlexer.Lexer, out *Serverless_Beta_Result) {
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(out *jwriter.Writer, in Serverless_Beta_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Success).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_Beta_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_Beta_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_Beta_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_Beta_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(in *jlexer.Lexer, out *Serverless_Beta_Args) {
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(Request)
				}
				(*out.Request).UnmarshalEasyJSON(in)
			}
		case "alpha":
			if in.IsNull() {
				in.Skip()
				out.Alpha = nil
			} else {
				if out.Alpha == nil {
					out.Alpha = new(string)
				}
				*out.Alpha = string(in.String())
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(out *jwriter.Writer, in Serverless_Beta_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Request != nil {
		const prefix string = ",\"request\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Request).MarshalEasyJSON(out)
	}
	if in.Alpha != nil {
		const prefix string = ",\"alpha\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Alpha))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Serverless_Beta_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Serverless_Beta_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Serverless_Beta_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Serverless_Beta_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerlessServerlessBeta1(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(in *jlexer.Lexer, out *Response) {
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
		case "firstName":
			if in.IsNull() {
				in.Skip()
				out.FirstName = nil
			} else {
				if out.FirstName == nil {
					out.FirstName = new(string)
				}
				*out.FirstName = string(in.String())
			}
		case "lastName1":
			if in.IsNull() {
				in.Skip()
				out.LastName1 = nil
			} else {
				if out.LastName1 == nil {
					out.LastName1 = new(string)
				}
				*out.LastName1 = string(in.String())
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FirstName != nil {
		const prefix string = ",\"firstName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.FirstName))
	}
	if in.LastName1 != nil {
		const prefix string = ",\"lastName1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.LastName1))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless(l, v)
}
func easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(in *jlexer.Lexer, out *Request) {
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
		case "firstName":
			if in.IsNull() {
				in.Skip()
				out.FirstName = nil
			} else {
				if out.FirstName == nil {
					out.FirstName = new(string)
				}
				*out.FirstName = string(in.String())
			}
		case "lastName":
			if in.IsNull() {
				in.Skip()
				out.LastName = nil
			} else {
				if out.LastName == nil {
					out.LastName = new(string)
				}
				*out.LastName = string(in.String())
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
func easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FirstName != nil {
		const prefix string = ",\"firstName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.FirstName))
	}
	if in.LastName != nil {
		const prefix string = ",\"lastName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.LastName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE550ff0eEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE550ff0eDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsServerlessServerless1(l, v)
}