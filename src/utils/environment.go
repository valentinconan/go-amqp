package utils

import "os"

func Getenv(value string, defaultvalue string) string {
   result := defaultvalue
   if value != "" && os.Getenv(value) != "" {
      result = os.Getenv(value)
   }
   return result
}