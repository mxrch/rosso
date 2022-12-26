package quotedprintable

import (
   "testing"
)

type URL_test struct {
   decode rune
   encode string
}

var tests = []URL_test{
   {'%', "%25"},
   {'\uFFFD', "%EF%BF%BD"},
   // graphic 1
   {'+', "+"},
   // graphic 2
   {'¶', "¶"},
   // graphic 3
   {'☺', "☺"},
   // graphic 4
   {'😀', "😀"},
   // not graphic 1
   {'\x1F', "%1F"},
   // not graphic 2
   {'\x80', "%C2%80"},
   // not graphic 3
   {'\u082e', "%E0%A0%AE"},
   // not graphic 4
   {'\U0001000C', "%F0%90%80%8C"},
}

func Test_Encode(t *testing.T) {
   for _, test := range tests {
      s := Encode(string(test.decode))
      if s != test.encode {
         t.Fatalf("%x %v", test.decode, s)
      }
   }
}

func Test_Decode(t *testing.T) {
   for _, test := range tests {
      s, err := Decode(test.encode)
      if err != nil {
         t.Fatal(err)
      }
      if s != string(test.decode) {
         t.Fatal(s)
      }
   }
}
