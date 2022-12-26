package ascii85

import (
   "os"
   "testing"
)

func Test_Encode(t *testing.T) {
   src, err := os.Open("apple-audio.m3u8")
   if err != nil {
      t.Fatal(err)
   }
   defer src.Close()
   dst, err := os.Create("apple-audio.txt")
   if err != nil {
      t.Fatal(err)
   }
   defer dst.Close()
   if err := Encode(dst, src); err != nil {
      t.Fatal(err)
   }
}
