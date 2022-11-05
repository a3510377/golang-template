package i18n

import (
	/*
		#cgo CFLAGS: -x objective-c
		#cgo LDFLAGS: -framework Foundation
		#import <Foundation/Foundation.h>

		const char *getLocaleIdentifier() {
			NSString *cs = [[NSLocale currentLocale] localeIdentifier];
			const char *cstr = [cs UTF8String];
			return cstr;
		}
	*/
	"C"
)

func getLocale() string {
	if locale := getLocalFromEnv(); locale != "" {
		return locale
	}

	return C.GoString(C.getLocaleIdentifier())
}
