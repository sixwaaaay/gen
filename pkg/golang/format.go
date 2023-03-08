package golang

import "go/format"

func FormatCode(code string) string {
	ret, err := format.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}
