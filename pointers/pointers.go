package pointers

// TODO: Add test
func StrPtr(str string) *string {
	return &str
}

func BoolPtr(bl bool) *bool {
	return &bl
}

func Float64Ptr(f float64) *float64 {
	return &f
}
func IntPtr(i int) *int {
	return &i
}
func Int64Ptr(i int64) *int64 {
	return &i
}
// TODO: Add test
func IsBlank(pStr *string) bool {
	return pStr == nil || *pStr == ""
}
