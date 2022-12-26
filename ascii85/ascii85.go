package ascii85

import (
   "bytes"
   "encoding/ascii85"
   "io"
)

func Encode(dst io.Writer, src io.Reader) error {
   buf := new(bytes.Buffer)
   _, err := io.Copy(ascii85.NewEncoder(buf), src)
   if err != nil {
      return err
   }
   for {
      _, err := io.CopyN(dst, buf, 80)
      if _, err := io.WriteString(dst, "\n"); err != nil {
         return err
      }
      if err == io.EOF {
         return nil
      } else if err != nil {
         return err
      }
   }
}
