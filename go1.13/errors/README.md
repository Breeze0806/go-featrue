# errors
本篇主要记录errors库1.13的新增特性，如果需要使用1.13之前版本的erros包，
最好把原有库代码的errors库使用[golang.org/x/xerrors](https://github.com/golang/xerrors)

## 函数Unwrap
使用该函数的前提是要让error的结构体实现Unwrap() error方法才能返回原错误，否则将返回nil

## 函数Is
这个函数能判断error是否相等，如需自定义判断条件需要实现 Is(err error) bool方法才可以，
注意target不会进行Unwrap处理，errors在比较时会进行Unwrap处理，可以传入nil方法。
用法是当原先以下列方式进行比较时
```golang
if err == io.ErrUnexpectedEOF{...}
```
可以替换为
```golang
if errors.Is(err, io.ErrUnexpectedEOF){...}
```
## 函数As
这个函数能做类型断言或者类型转换，如需自定义判断条件需要实现 As(err interface{}) bool方法才可以，
注意target不会进行Unwrap处理，errors在比较时会进行Unwrap处理，target不能可以传入nil，必须传入
指向error类型的指针，否则会panic
用法是当原先以下列方式进行比较时
```golang
if e, ok := err.(*os.PathError); ok{...}
```
可以替换为
```golang
var e *os.PathError
if errors.As(err, &e){...}
```
## Errorf
采用%w对于实现Unwrap方法的error进行打印

## 参考文献
更多的请查看[ErrorValueFAQ](https://github.com/golang/go/wiki/ErrorValueFAQ)