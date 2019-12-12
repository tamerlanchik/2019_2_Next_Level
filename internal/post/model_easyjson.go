// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package post

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

func easyjsonC80ae7adDecode20192NextLevelInternalPost(in *jlexer.Lexer, out *Email) {
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
		case "From":
			out.From = string(in.String())
		case "To":
			out.To = string(in.String())
		case "Body":
			out.Body = string(in.String())
		case "Subject":
			out.Subject = string(in.String())
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
func easyjsonC80ae7adEncode20192NextLevelInternalPost(out *jwriter.Writer, in Email) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"From\":"
		out.RawString(prefix[1:])
		out.String(string(in.From))
	}
	{
		const prefix string = ",\"To\":"
		out.RawString(prefix)
		out.String(string(in.To))
	}
	{
		const prefix string = ",\"Body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"Subject\":"
		out.RawString(prefix)
		out.String(string(in.Subject))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Email) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncode20192NextLevelInternalPost(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Email) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncode20192NextLevelInternalPost(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Email) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecode20192NextLevelInternalPost(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Email) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecode20192NextLevelInternalPost(l, v)
}
