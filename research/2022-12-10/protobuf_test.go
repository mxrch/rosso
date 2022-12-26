package protobuf

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Unmarshal(t *testing.T) {
   buf, err := os.ReadFile("com.pinterest.txt")
   if err != nil {
      t.Fatal(err)
   }
   mes, err := unmarshal(buf)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   if err := enc.Encode(mes); err != nil {
      t.Fatal(err)
   }
}
