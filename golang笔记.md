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

%q 占位符，用于格式化字符串时，将值用双引号包裹起来，并对特殊字符进行转义。这使得输出的字符串可以直接作为 Go 语言的字符串字面量使用。

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


8.0 Map
map 是一种特殊的数据结构：一种元素对 (pair) 的*无序*集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。

map 这种数据结构在其他编程语言中也称为字典 (Python) 、hash 和 HashTable 等。
令 v := map1[key1] 可以将 key1 对应的值赋值给 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。
mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。make去初始化就不会传字面量了
8.1.2 map 容量
和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype, cap)。例如：

map2 := make(map[string]float32, 100)
当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

new() 和 make() 的区别
看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 *T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel（参见第 8 章和第 13 章）。
换言之，new() 函数分配内存，make() 函数初始化；

make只能用于slice切片、map、channel

```go
package main

type Foo map[string]string
type Bar struct {
    thingOne string
    thingTwo int
}

func main() {
    // OK
    y := new(Bar)
    (*y).thingOne = "hello"
    (*y).thingTwo = 1

    // NOT OK
    z := make(Bar) // 编译错误：cannot make type Bar
    (*z).thingOne = "hello"
    (*z).thingTwo = 1

    // OK
    x := make(Foo)
    x["x"] = "goodbye"
    x["y"] = "world"

    // NOT OK
    u := new(Foo)
    (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
    (*u)["y"] = "world"
}


 ```

 可以看到结构体可以new，但不能make，直接无法编译，map则相反，不可以new，new会触发运行时错误，make是可以的


 结构体
```go
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
	innerS //anonymous field
}
```
 通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。

``` go
 type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	fmt.Println(p.getName()) // Tom
}
``` 
Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可。
实例化 Student后，强制类型转换为接口类型 Person。

但是删除 (*Worker).getName() 程序并不会报错，因为我们并没有在 main 函数中使用。这种情况下我们如何确保某个类型实现了某个接口的所有方法呢？一般可以使用下面的方法进行检测，如果实现不完整，编译期将会报错。

```go
var _ Person = (*Student)(nil)
var _ Person = (*Worker)(nil)

// 接口转化为实例
func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	stu := p.(*Student) // 接口转为实例
	fmt.Println(stu.getAge())
}

```
实例化后赋值给接口类型，查看该实例是否完全实现该接口


结构体：
有区别吗？
```go
type Base struct {
  Name string
}
type Derived struct {
  Base
  Age int
}
和
type Base struct {
  Name string
}

type Derived struct {
  *Base
  Age int
}
```



是的，这两个 Derived 结构体的定义是有区别的。

1. 嵌入类型（Embedded Type）

type Derived struct { Base Age int }

这里 Base 是一个嵌入类型。
Derived 结构体直接包含了 Base 结构体的所有字段和方法。
访问 Base 结构体的字段时，可以直接通过 derived.Name 访问，无需使用点号两次。
2. 嵌入类型指针（Embedded Type Pointer）

type Derived struct { *Base Age int }

这里 *Base 是一个指向 Base 结构体的指针。
Derived 结构体包含了一个指向 Base 结构体的指针。
访问 Base 结构体的字段时，需要通过指针来访问，例如 derived.Base.Name。
示例


```go
package main

import "fmt"

type Base struct {
        Name string
}

type Derived1 struct {
        Base
        Age  int
}

type Derived2 struct {
        *Base
        Age  int
}

func main() {
        d1 := Derived1{Base: Base{"Alice"}, Age: 30}
        fmt.Println(d1.Name) // 输出：Alice

        b := &Base{"Bob"}
        d2 := Derived2{Base: b, Age: 25}
        fmt.Println(d2.Base.Name) // 输出：Bob
}
```

区别总结

访问方式不同： 嵌入类型可以直接访问字段，嵌入类型指针需要通过指针来访问。
内存占用： 嵌入类型直接包含了 Base 结构体的数据，而嵌入类型指针只包含了一个指针，节省了内存空间，但增加了访问的复杂性。
修改影响：
嵌入类型：修改 Derived 的实例会直接影响到 Base 结构体中的字段。
嵌入类型指针：修改 Derived 的实例不会直接影响到原始的 Base 结构体。

Go 没有传统的类继承，但通过嵌套结构体，可以实现类似继承的效果。嵌套的 RouterGroup 结构体的所有方法和字段都可以被 Engine 结构体直接访问