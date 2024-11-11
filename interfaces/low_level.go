package interfaces

/*
#cgo pkg-config: openssl
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lopenssl_pkcs7 -Wl,-rpath,/usr/local/lib
#include "openssl_pkcs7.h"
*/
import "C"
import (
	"unsafe"
)

const NID_INN int = (int)(C.NID_INN)
const NID_SNILS int = (int)(C.NID_SNILS)
const NID_CN int = (int)(C.NID_commonName)
const NID_givenName int = (int)(C.NID_givenName)
const NID_surname int = (int)(C.NID_surname)
const NID_OGRN = (int)(C.NID_OGRN)
const NID_OGRNIP = (int)(C.NID_OGRNIP)

func VerifyPkcs7(ca []byte, signature []byte, data []byte) bool {
	return C.verify_pkcs7(C.CBytes(ca), C.size_t(len(ca)), C.CBytes(signature), C.size_t(len(signature)), C.CBytes(data), C.size_t(len(data))) == 1
}

func GetPkcs7Attr(signature []byte, attr int) string {
	val := C.pkcs7_attr(C.CBytes(signature), C.int(len(signature)), C.int(attr))
	retVal := C.GoString(val)
	C.free(unsafe.Pointer(val))
	return retVal
}
