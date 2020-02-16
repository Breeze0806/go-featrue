# 数字字面量

## 新增支持进制字面量

+ 2进制字面量：0b1011或者0B1011
+ 8进制字面量：0x1011或者0X1011
+ 16进制浮点字面量：0x1.0p-1021
+ i复数自字面量：0x2i - 1
+ 字面量分隔符： 1000_000

## 位移符号
<< 和 >>可以用于所有整形，不再局限于非负数

## go/scanner
对于上述的字面量都识别为1个整数

## text/scanner
对于上述的字面量都识别为1个整体

## math/big
+ new(big.Int).SetString,big.ParseFloat
,new(big.Float).Parse 将base设为0时，按照字面量对应进制数去解析
+ new(big.Float).SetString， new(big.Rat).SetString
按照字面量去解析

## strconv
+ strconv.ParseInt和strconv.ParseInt将base设为0时，按照字面量对应进制数去解析
+ strconv.ParseFloat按照字面量去解析会失败

## 参考文献
[字面量设计](https://github.com/golang/proposal/blob/master/design/19308-number-literals.md)