package reload

import (
	"strconv"
	"strings"
)

type Flags struct {
	up  string
	low string
	cap string
	bin string
	hex string
}

func (f *Flags) makeFlags() Flags {
	f.up = "(up)"
	f.low = "(low)"
	f.cap = "(cap)"
	f.bin = "(bin)"
	f.hex = "(hex)"
	return *f
}

func (f *Flags) upFlag(flag string) bool {
	return flag == "(up)"
}

func (f *Flags) lowFlag(flag string) bool {
	return flag == "(low)"
}

func (f *Flags) capFlag(flag string) bool {
	return flag == "(cap)"
}

func (f *Flags) binFlag(flag string) bool {
	return flag == "(bin)"
}

func (f *Flags) hexFlag(flag string) bool {
	return flag == "(hex)"
}

func (f *Flags) capSpecified(flag string) bool {
	if !strings.HasPrefix(flag, "(cap, ") || !strings.HasSuffix(flag, ")") {
		return false
	}

	content := flag[len("(cap, ") : len(flag)-1]

	trimmedContent := strings.TrimSpace(content)

	_, err := strconv.Atoi(trimmedContent)
	if err != nil  {
		return false
	}else{
		return true
	}
}

func (f *Flags) upSpecified(flag string) bool {
	if !strings.HasPrefix(flag, "(up, ") || !strings.HasSuffix(flag, ")") {
		return false
	}

	content := flag[len("(up, ") : len(flag)-1]

	trimmedContent := strings.TrimSpace(content)

	_, err := strconv.Atoi(trimmedContent)
	if err != nil  {
		return false
	}else{
		return true
	}
}

func (f *Flags) lowSpecified(flag string) bool {
	if !strings.HasPrefix(flag, "(low, ") || !strings.HasSuffix(flag, ")") {
		return false
	}

	content := flag[len("(low, ") : len(flag)-1]

	trimmedContent := strings.TrimSpace(content)

	_, err := strconv.Atoi(trimmedContent)
	if err != nil  {
		return false
	}else{
		return true
	}
}

func (f *Flags) isFlag(flag string) bool {
	checkFuncs := []func(string) bool{
		f.upFlag,
		f.lowFlag,
		f.capFlag,
		f.binFlag,
		f.hexFlag,
		f.capSpecified,
		f.upSpecified,
		f.lowSpecified,
	}

	for _, checkFunc := range checkFuncs {
		if checkFunc(flag) {
			return true
		}
	}
	return false
}

var flags Flags

func init() {
	flags.makeFlags()
}
