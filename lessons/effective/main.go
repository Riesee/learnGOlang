package main

import (
	"fmt"
	"io"
	"log"
	"os"
)


type T struct {
    a int
    b float64
    c string
}

type ByteSize float64

type ByteSlice []byte




func main() {
// 	for i := 0; i < 5; i++ {
//     defer fmt.Printf("%d ", i)
// }

// var p *[]int = new([]int)     
// fmt.Println(p, "pppp")

// var v  []int = make([]int, 100) 
// fmt.Println(&v, "vvvv")

// q := make([]int, 100)
// fmt.Println(q, "qqqq")


// array := [...]float64{7.0, 8.5, 9.1}
// x := Sum(&array)
// fmt.Println(x)

// fmt.Printf("Hello %d\n", 23)
// fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
// fmt.Println("Hello", 23)
// fmt.Println(fmt.Sprint("Hello ", 23))


// t := &T{ 7, -2.35, "abc\tdef" }
// m := MyString("hello")
// fmt.Printf("%s\n", m)
// fmt.Printf("%v\n", t)
// fmt.Printf("%+v\n", t)
// fmt.Printf("%#v\n", t)
// fmt.Printf("%v\n", t)

// test := Min(32, 30, 3234, 23454, 123455)
// fmt.Println(test)



// fmt.Print(ByteSize(1e13), "\n")

// var b ByteSlice
// b'nin adresini (&b) gönderiyoruz çünkü Write metodu pointer bekliyor (*ByteSlice).
// fmt.Fprintf(&b, "Bu saatin %d günü var", 7)

// fmt.Println(string(b))

// b.Write([]byte("Merhaba "))
//     b.Write([]byte("Dünya!"))
    
//     fmt.Println(string(b)) 


fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err, "\ttest.go")
    }
    fmt.Println(io.Copy(os.Stdout, fd))

}

// func (bs *ByteSlice) Write(p []byte) (n int, err error) {
//     *bs = append(*bs, p...)
//     return len(p), nil
// }


// var (
// home   = os.Getenv("HOME")
// user   = os.Getenv("USER")
// gopath = os.Getenv("GOPATH")
// )

// func init() {
//     if user == "" {
//         log.Fatal("$USER not set")
//     }
//     if home == "" {
//         home = "/home/" + user
//     }
//     if gopath == "" {
//         gopath = home + "/go"
//     }
//     // gopath, komut satırındaki --gopath bayrağı ile geçersiz kılınabilir.
//     flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
// }

// func (b ByteSize) String() string {
// 	const (
//     _           = iota // boş tanımlayıcıya atayarak ilk değeri yok say
//     KB ByteSize = 1 << (10 * iota)
//     MB
//     GB
//     TB
//     PB
//     EB
//     ZB
//     YB
// )

//     switch {
//     case b >= YB:
//         return fmt.Sprintf("%.2fYB", b/YB)
//     case b >= ZB:
//         return fmt.Sprintf("%.2fZB", b/ZB)
//     case b >= EB:
//         return fmt.Sprintf("%.2fEB", b/EB)
//     case b >= PB:
//         return fmt.Sprintf("%.2fPB", b/PB)
//     case b >= TB:
//         return fmt.Sprintf("%.2fTB", b/TB)
//     case b >= GB:
//         return fmt.Sprintf("%.2fGB", b/GB)
//     case b >= MB:
//         return fmt.Sprintf("%.2fMB", b/MB)
//     case b >= KB:
//         return fmt.Sprintf("%.2fKB", b/KB)
//     }
//     return fmt.Sprintf("%.2fB", b)
// }

// func (t *T) String() string {
//     return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
// }

// type MyString string

// func (m MyString) String() string {
//     return fmt.Sprintf("MyString=%s", string(m))
// }

// func Min(a ...int) int {
//     min := int(^uint(0) >> 1)
//     for _, i := range a {
//         if i < min {
//             min = i
//         }
//     }
//     return min
// }


// func Sum(a *[3]float64) (sum float64) {
//     for _, v := range *a {
//         sum += v
//     }
//     return
// }
