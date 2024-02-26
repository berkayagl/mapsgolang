# Hash Tablosu Nedir?

Hash tablosu, verilerin ilişkisel bir şekilde depolandığı bir veri yapısıdır.
Veriler, her anahtardan benzersiz bir değer üreten bir karma işlev tarafından dizi konumlarına eşlenir.

![alt text](https://miro.medium.com/v2/resize:fit:1400/1*F68yzTL2UVGg9LhsaeJJMA.png)

Hash tabloların en büyük avantajı verimli olmalarıdır.
Bir öğeye erişim süresi ortalama O(1)'dir, bu nedenle daha hızlı arama, ekleme ve silme işlemleri sunar.
Hash fonksiyonu anahtardan adresler üretir.
Anahtarları uygun konuma eşleme işlemine hashleme denir.

Go'daki map veri tipi, anahtarları değerlere eşlediği anlamına gelir.
Anahtarları değerlerle eşler ve verileri Go'da saklar.

# Golang Maps

![alt text](https://donofden.com/images/doc/golang-map-1.png)

Map türleri, reference veya slices gibi pointers türleridir.
Boş bir map Go'da nil veri türü gibi davranır,
oysa sadece bir nil map bildirirseniz, bunun gibi boş bir harita döndürecektir:

```go
package main

func main() {

	var m map[string]string
	// map

}

```

Eğer nil bir map yazmaya çalışırsak, çalışma zamanı panik hatasına neden olacaktır.

```go
package main

import "fmt"

func main() {

	var m map[string]string

	m["name"] = "Berkay"

	fmt.Println(m)
}
```

Hata şu şekilde olur :

```
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
	/X/X/X/X/GolangExample/maps.go:9 +0x32
exit status 2

```
Bu durumda, var keyword kullanarak bir map bildirdiğimiz için,
bu çalışma zamanı panik hatasını iletmek için yerleşik bir make işlevi kullanabiliriz.
Bu fonksiyonu kullanırken, bir hash map veri yapısı ayırır ve başlatır ve onu işaret eden bir map değeri döndürür.

```go
package main

import "fmt"

func main() {

	var m map[string]string

	m = make(map[string]string)

	m["name"] = "Berkay"

	fmt.Println(m)
}

// Çıktısı : map[name:Berkay]

```

Dolayısıyla, bir haritayı map literal kullanarak da başlatabiliriz.

```go

package main

import "fmt"

func main() {

	m := map[string]string{"name": "Berkay", "age": "20", "status": "Busy"}

	fmt.Println(m)
}

// Çıktısı : map[age:20 name:Berkay status:Busy]
```

Ayrıca, yerleşik make fonksiyonunu kullanmakla aynı olan map değişmezleri ile boş bir diziyi başlatabiliriz.

```go
package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}

	for key, value := range mapData {
		fmt.Println("\nKey:", key, "\nValue:", value)
	}

}

// Çıktısı

// Key: name 
// Value: Berkay

// Key: age 
// Value: 20

// Key: status 
// Value: Busy

```

Go'nun kodumuzda daha fazla işlem yapabilmemiz için bize sağladığı iki yerleşik fonksiyon daha vardır,
map uzunluğunu kontrol etmek için len'i kullanabiliriz ve ayrıca mevcut anahtarlardan bazılarını silmek istiyorsak,
değerleri argüman olarak geçirirken sadece delete fonksiyonunu uygulayabiliriz.
Dahası, delete fonksiyonu hiçbir şey döndürmez ve belirtilen anahtar mevcut değilse hiçbir şey yapmaz.
Bu nedenle, yukarıdaki mapData haritası için bu iki yerleşik işlevi deniyorum.

Lenght:

```go
package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}

	l := len(mapData)

	fmt.Println(l)
}

// Çıktısı : 3

```

Delete:

```go
package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}
	delete(mapData, "status")

	fmt.Println(mapData)
}

// Çıktısı : map[age:20 name:Berkay]

```

Görüşmek üzere <3
