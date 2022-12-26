package quotedprintable

import (
   "bytes"
   "net/url"
   "unicode"
   "unicode/utf8"
)

func Decode(s string) (string, error) {
   return url.PathUnescape(s)
}

func Encode(s string) string {
   var buf bytes.Buffer
   for len(s) >= 1 {
      r, size := utf8.DecodeRuneInString(s)
      value := s[:size]
      if should_escape(r) {
         buf.WriteString(url.PathEscape(value))
      } else {
         buf.WriteString(value)
      }
      s = s[size:]
   }
   return buf.String()
}

func Encode_Bytes(p []byte) []byte {
   var buf bytes.Buffer
   for len(p) >= 1 {
      r, size := utf8.DecodeRune(p)
      value := p[:size]
      if should_escape(r) {
         buf.WriteString(url.PathEscape(string(value)))
      } else {
         buf.Write(value)
      }
      p = p[size:]
   }
   return buf.Bytes()
}

func should_escape(r rune) bool {
   if r == '%' {
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
