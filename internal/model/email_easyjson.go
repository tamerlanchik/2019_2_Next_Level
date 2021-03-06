// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonFc263e4eDecode20192NextLevelInternalModel(in *jlexer.Lexer, out *Email) {
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
		case "Header":
			easyjsonFc263e4eDecode(in, &out.Header)
		case "Id":
			out.Id = int(in.Int())
		case "IsRead":
			out.IsRead = bool(in.Bool())
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
func easyjsonFc263e4eEncode20192NextLevelInternalModel(out *jwriter.Writer, in Email) {
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
		const prefix string = ",\"Header\":"
		out.RawString(prefix)
		easyjsonFc263e4eEncode(out, in.Header)
	}
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix)
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"IsRead\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRead))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Email) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc263e4eEncode20192NextLevelInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Email) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc263e4eEncode20192NextLevelInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Email) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc263e4eDecode20192NextLevelInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Email) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc263e4eDecode20192NextLevelInternalModel(l, v)
}
func easyjsonFc263e4eDecode(in *jlexer.Lexer, out *struct {
	From         string
	To           []string
	Subject      string
	ReplyTo      []string
	WhenReceived time.Time
}) {
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
			if in.IsNull() {
				in.Skip()
				out.To = nil
			} else {
				in.Delim('[')
				if out.To == nil {
					if !in.IsDelim(']') {
						out.To = make([]string, 0, 4)
					} else {
						out.To = []string{}
					}
				} else {
					out.To = (out.To)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.To = append(out.To, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Subject":
			out.Subject = string(in.String())
		case "ReplyTo":
			if in.IsNull() {
				in.Skip()
				out.ReplyTo = nil
			} else {
				in.Delim('[')
				if out.ReplyTo == nil {
					if !in.IsDelim(']') {
						out.ReplyTo = make([]string, 0, 4)
					} else {
						out.ReplyTo = []string{}
					}
				} else {
					out.ReplyTo = (out.ReplyTo)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ReplyTo = append(out.ReplyTo, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "WhenReceived":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.WhenReceived).UnmarshalJSON(data))
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
func easyjsonFc263e4eEncode(out *jwriter.Writer, in struct {
	From         string
	To           []string
	Subject      string
	ReplyTo      []string
	WhenReceived time.Time
}) {
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
		if in.To == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.To {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Subject\":"
		out.RawString(prefix)
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"ReplyTo\":"
		out.RawString(prefix)
		if in.ReplyTo == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ReplyTo {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"WhenReceived\":"
		out.RawString(prefix)
		out.Raw((in.WhenReceived).MarshalJSON())
	}
	out.RawByte('}')
}
