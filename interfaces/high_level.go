package interfaces

type AttrType int

const (
	GivenName AttrType = AttrType(NID_givenName)
	Surname   AttrType = AttrType(NID_surname)
	SNILS     AttrType = AttrType(NID_SNILS)
	INN       AttrType = AttrType(NID_INN)
	CN        AttrType = AttrType(NID_CN)
	OGRN      AttrType = AttrType(NID_OGRN)
	OGRNIP    AttrType = AttrType(NID_OGRNIP)
)

type Pkcs7 struct {
	data []byte
}

func FromBytes(bytes []byte) *Pkcs7 {
	return &Pkcs7{
		data: bytes,
	}
}

func (p *Pkcs7) Verify(ca []byte, data []byte) bool {
	return VerifyPkcs7(ca, p.data, data)
}

func (p *Pkcs7) GetAttr(attrType AttrType) string {
	return GetPkcs7Attr(p.data, int(attrType))
}
