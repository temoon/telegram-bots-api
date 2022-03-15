package helpers

func Bool(v bool) *bool       { return &v }
func Int32(v int32) *int32    { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }

func BoolOrFalse(ref *bool) bool {
	if ref == nil {
		return false
	}

	return *ref
}

func Int32OrZero(ref *int32) int32 {
	if ref == nil {
		return 0
	}

	return *ref
}

func Int64OrZero(ref *int64) int64 {
	if ref == nil {
		return 0
	}

	return *ref
}

func StringOrEmpty(ref *string) string {
	if ref == nil {
		return ""
	}

	return *ref
}
