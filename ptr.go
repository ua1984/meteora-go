package meteora

// String returns a pointer to the string value.
func String(s string) *string {
	return &s
}

// Int returns a pointer to the int value.
func Int(i int) *int {
	return &i
}

// Bool returns a pointer to the bool value.
func Bool(b bool) *bool {
	return &b
}
