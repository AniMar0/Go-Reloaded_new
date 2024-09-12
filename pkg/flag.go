package reload

import (
	"strings"
	"unicode"
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
	return strings.HasPrefix(flag, "(cap, ") && strings.HasSuffix(flag, ")")
}

func (f *Flags) upSpecified(flag string) bool {
	return strings.HasPrefix(flag, "(up, ") && strings.HasSuffix(flag, ")")
}

func (f *Flags) lowSpecified(flag string) bool {
	return strings.HasPrefix(flag, "(low, ") && strings.HasSuffix(flag, ")")
}

func (f *Flags) skipPonc(word string) bool {
	for _, char := range word {
		if unicode.IsLetter(rune(char)) {
			return false
		}
	}
	return true
}

func (f *Flags) checkFlag(flag string) bool {
	checkFuncs := []func(string) bool{
		f.upFlag,
		f.lowFlag,
		f.capFlag,
		f.binFlag,
		f.hexFlag,
		f.capSpecified,
		f.upSpecified,
		f.lowSpecified,
		f.skipPonc,
	}

	for _, checkFunc := range checkFuncs {
		if checkFunc(flag) {
			return true
		}
	}
	return false
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
