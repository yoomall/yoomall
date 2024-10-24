package execl

import "yoomall/core"

func TimeFormatter(v interface{}) interface{} {
	switch val := v.(type) {
	case string:
		return v
	case core.LocalTime:
		return val.Format("2006-01-02 15:04:05")
	default:
		return v
	}
}
