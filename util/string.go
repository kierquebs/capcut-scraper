package util

import (
	"strings"
)

func StringReplace(s , old, new string, n int) string{
	res := strings.Replace(s, old,new,n)
   return res
}