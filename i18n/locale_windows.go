package i18n

import (
	"syscall"
	"unsafe"
)

const LocaleNameMaxLength uint32 = 85

func getLocale() string {
	var locale string

	for _, proc := range [...]string{"GetUserDefaultLocaleName", "GetSystemDefaultLocaleName"} {
		if locale := getWinFromLocaleProc(proc); locale != "" {
			return locale
		}
	}

	return locale
}

func getWinFromLocaleProc(call string) string {
	dll, err := syscall.LoadDLL("kernel32")
	if err != nil {
		return ""
	}
	defer dll.Release()

	proc, err := dll.FindProc(call)
	if err != nil {
		return ""
	}

	buffer := make([]uint16, LocaleNameMaxLength)
	len, _, _ := proc.Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(LocaleNameMaxLength))

	if len == 0 {
		return ""
	}

	return syscall.UTF16ToString(buffer)
}
