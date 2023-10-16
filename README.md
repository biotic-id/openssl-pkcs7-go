# openssl-pkcs7-go

Пакет оборачивает C-библиотеку в высокоуровневый интерфейс. Основная задача пакета – проверка ГОСТ подписей от Госключа и получение метаданных из них.

Пример кода

```go
package main

import (
	"fmt"
	Pkcs7 "github.com/biotic-id/openssl-pkcs7-go/interfaces"
	"os"
)

func main() {
	var signature []byte
	var data []byte

	signature, e := os.ReadFile("/Users/roman/Desktop/IMG_2627.jpeg.sig")
	if e != nil {
		fmt.Errorf("signature read failed %v", e)
	}
	data, e = os.ReadFile("/Users/roman/Desktop/IMG_2627.jpeg")
	if e != nil {
		fmt.Errorf("data read failed %v", e)
	}

	var pkcs7 = Pkcs7.FromBytes(signature)

	var retVal = pkcs7.Verify(Pkcs7.NotQualifiedCA, data)
	fmt.Printf("%v\n", retVal)

	fmt.Printf("CN = %s\n", pkcs7.GetAttr(Pkcs7.CN))
	fmt.Printf("SNILS = %s\n", pkcs7.GetAttr(Pkcs7.SNILS))
	fmt.Printf("Given name = %s\n", pkcs7.GetAttr(Pkcs7.GivenName))
	fmt.Printf("Surname = %s\n", pkcs7.GetAttr(Pkcs7.Surname))
}
```

