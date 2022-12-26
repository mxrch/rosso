package quotedprintable

import (
   "bytes"
   "io"
   "mime/quotedprintable"
   "unicode"
   "unicode/utf8"
)

const upper_hex = "0123456789ABCDEF"

func should_escape(r rune) bool {
   if r == '=' {
      return true
   }
   if r == utf8.RuneError {
      return true
   }
   if unicode.IsControl(r) {
      return true
   }
   return unicode.IsMark(r)
}

///////////////////////////////////

func Decode(b []byte) ([]byte, error) {
   r := bytes.NewReader(b)
   return io.ReadAll(quotedprintable.NewReader(r))
}

func Encode(p []byte) []byte {
   var buf bytes.Buffer
   for len(p) >= 1 {
      r, size := utf8.DecodeRune(p)
      if should_escape(r) {
         for _, b := range p[:size] {
            buf.WriteByte('=')
            buf.WriteByte(upper_hex[b>>4])
            buf.WriteByte(upper_hex[b&0xF])
         }
      } else {
         buf.Write(p[:size])
      }
      p = p[size:]
   }
   return buf.Bytes()
}
