new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

等价于

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```


go 数据类型

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

这些类型的长度都是根据运行程序所在的操作系统类型所决定的：

int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
uintptr 的长度被设定为足够存放一个指针即可。

Go 语言中没有 float 类型。（Go 语言中只有 float32 和 float64）没有 double 类型。

与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：

整数：

int8（-128 -> 127）
int16（-32768 -> 32767）
int32（-2,147,483,648 -> 2,147,483,647）
int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
无符号整数：

uint8（0 -> 255）
uint16（0 -> 65,535）
uint32（0 -> 4,294,967,295）
uint64（0 -> 18,446,744,073,709,551,615）
浮点型（IEEE-754 标准）：

float32（+- 1e-45 -> +- 3.4 * 1e38）
float64（+- 5 1e-324 -> 107 1e308）



字符类型

byte类型==uint8

rune == int32 表示Unicode(UTF-8)码点

格式化说明符

在格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字），%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），%0d 用于规定输出定长的整数，其中开头的数字 0 是必须的。

%n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00。

fmt.Println跟println调用后的打印结果不同

```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

func main() {
    var shape Shape
    shape = Circle{5}
    fmt.Println(shape.Area())
}
```

在这个例子中：

Shape 是一个接口，定义了 Area() 方法。
Circle 结构体实现了 Shape 接口。
main 函数中，我们通过接口 Shape 来操作 Circle 对象。