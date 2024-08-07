package mystrings


// only uper case name are exported
func Reverse(s string) string {
	result := ""
	for _,v := range s{
		result = string(v)+result
	}
	return result
}