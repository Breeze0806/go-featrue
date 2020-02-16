# 其他包特性
这里主要罗列了其他包的相关特性

## bytes
ToValidUTF8会返回不合乎utf8规范的字符数组

## strings
同bytes

## database/sql
实现了NullInt32以及NullTime，另外可以借鉴这个修改我们的转换函数

## log
Writer返回的是os.Stderr,并非其用英文表述的os.Stdout

## os
UserHomeDir返回的是用户默认路径
在文件以追加写的方式返回时，WriteAt方法都会返回错误

##runtime
Caller，Callers输出函数时仅仅输出包名.函数名

## reflect
Value.IsZero,判断数值是否是零值，如int零值为0，string零值为"",chan零值为nil等,
**注意nil会大致panic**,MakeFunc具体见例子

## sync（todo）
+ 读写锁，锁效率提升10%，Once函数效率提升两倍
，这些提升是通过内联的方式实现的，目前还未实现对应的压测实例
+ Pool不再数次提升stop the world的时间，减少了gc压力

## time
parse, format对于时间支持输出本年度的第几天
新增Duration的方法Microseconds和Milliseconds