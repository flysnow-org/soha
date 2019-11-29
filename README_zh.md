# SOHA
SOHA是一个Go语言(Golang)模板的增强库，它通过自定义函数的方式扩展了Golang自带的模板引擎能力，让我们可以更好的开发Go Web应用。

[English Doc](README.md)

## 要求

1. Go version 1.13+
2. Using Go Modules

## 快速开始

```
# 例子在 example/hello.go 文件里
$ cat example/hello.go 
```

```go
package main

import (
	"github.com/flysnow-org/soha"
	"html/template"
	"log"
	"os"
)
func main() {

	sohaFuncMap := soha.CreateFuncMap()
	const templateText = `{{md5 .}}`

	tmpl, err := template.New("titleTest").Funcs(sohaFuncMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(os.Stdout, 123456)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
```

```
➜  soha go run example/hello.go  
e10adc3949ba59abbe56e057f20f883e
```

## SOHA 可以使用的函数介绍

#### int
基于给定的参数值，创建一个`int`
```
int INPUT
```

#### string
基于给定的参数值，创建一个`string`
```
string INPUT
```

#### float
基于给定的参数值，创建一个`float64`
```
float INPUT
```

---

#### first
截取集合的前N个元素为一个新的切片
```
first LIMIT COLLECTION
```

#### last
截取集合的后N个元素为一个新的切片
```
last INDEX COLLECTION
```

#### after
截取集合第N个元素之后的元素为一个新的切片
```
after INDEX COLLECTION
```

#### complement
返回集合里不存在的其他元素
```
COLLECTION | complement COLLECTION [COLLECTION]...
```

#### symdiff
返回两个集合的对称差集(也就是去掉重复元素后的集合)
```
COLLECTION | symdiff COLLECTION
```

#### dict
基于KV键值对创建一个字典(map)
```
dict KEY VALUE [KEY VALUE]...
```

#### echoParam
打印字典Key对应的值，如果Key不存在则打印空字符串
```
echoParam DICTIONARY KEY
```

#### in
检测一个元素是否在数组或者切片中，或者一个子字符串是否在一个字符串中，返回一个boolean
```
in SET ITEM
```

#### intersect
返回两个集合的合集
```
intersect SET1 SET2
```

#### union
返回两个集合(数组、切片)的并集
```
union SET1 SET2
```

#### isset
检测集合是否存在给定的索引INDEX，Map是否存在给定的Key
```
isset COLLECTION INDEX

isset COLLECTION KEY
```

#### querify
把给定的KV键值对转化为编码后的URL Query Parameter
```
querify KEY VALUE [KEY VALUE]...
```

#### shuffle
对给定的数组或者切片进行随机排序
```
shuffle COLLECTION
```

#### where
和我们SQL的where差不多，用于根据给定的字段值过滤出一个新的集合。这里需要注意的是集合必须为Array、Slice和Map，
如果是Map的话，那么Map对应的Value必须是Array、Slice，而这些Array、Slice中存储的是一个个DataType集合。

`OPERATOR`是一个可选的操作符，不写的话是`=`，也可以是其他`!=`,`>` `in`等操作符
```
where COLLECTION KEY [OPERATOR] MATCH
```

#### append
添加一个或者多个元素到Slice，并且返回一个新Slice
```
COLLECTION | append VALUE [VALUE]...

COLLECTION | append COLLECTION
```

#### seq
创建一个`int`类型的队列(slice)
```
seq LAST

seq FIRST LAST

seq FIRST INCREMENT LAST
```

#### uniq
去除array或者slice中重复的元素，返回一个新的去重后的slice
```
uniq SET
```

---

#### md5
返回MD5散列后的哈希值
```
md5 INPUT
```

#### sha1
返回SHA1散列后的哈希值
```
sha1 INPUT
```

#### sha256
返回SHA256散列后的哈希值
```
sha256 INPUT
```

---

#### base64
`base64Encode` and `base64Decode` ,base64位编码和解码。
```
base64Decode INPUT

base64Encode INPUT
```

#### jsonify
把一个对象编码为JSON.
```
jsonify INPUT
```

---

#### numFmt
数字格式化。`PRECISION`是精度，`OPTIONS`的组成格式为`<negative> <decimal> <grouping>`
```
numFmt PRECISION NUMBER [OPTIONS [DELIMITER]]
```

#### add
加法运算
```
add INPUT1 INPUT2
```

#### sub
减法运算
```
sub INPUT1 INPUT2
```

#### mul
乘法运算
```
mul INPUT1 INPUT2
```

#### div
除法运算
```
div INPUT1 INPUT2
```

#### mod
取模运算,返回 `INPUT1 % INPUT2`.
```
div INPUT1 INPUT2
```

#### modBool
是否可以整除
```
modBool INT1 INT2
```

#### math.Ceil
向上取整
```
math.Ceil FLOAT
```

#### math.Floor
向下取整
```
math.Floor FLOAT
```

#### math.Round
四舍五入，取最接近的整数
```
math.Round FLOAT
```

#### math.Log
对数函数
```
math.Log FLOAT
```

---

#### path.Base
返回路径中最后的元素
```
path.Base PATH
```

#### path.Dir
返回路径中除了最后一个元素之外的所有
```
path.Dir PATH
```

#### path.Ext
返回一个路径中文件扩展名
```
path.Ext PATH
```

#### path.Join
多个路径拼接为一个
```
path.Join ELEMENT...
```

#### path.Split
把路径拆分为目录和文件，返回一个`DirFile`
```
path.Split PATH
```

---

#### reflect.IsMap
判断是否是Map类型
```
reflect.IsMap INPUT
```

#### reflect.IsSlice
判断是否是Slice类型
```
reflect.IsSlice INPUT
```

#### safeCSS
转换为安全的CSS字符串
```
safeCSS INPUT
```

#### safeHTML
转换为安全的HTML字符串
```
safeHTML INPUT
```

#### safeHTMLAttr
转换为安全的HTML标签属性字符串
```
safeHTMLAttr INPUT
```

#### safeJS
转换为安全的JS字符串
```
safeJS INPUT
```

---

#### chomp
删除所有结尾的换行符
```
chomp INPUT
```

#### findRE
根据正则查找字符串，返回匹配的字符串列表
```
findRE PATTERN INPUT [LIMIT]
```

#### hasPrefix
字符串是否有某个前缀
```
hasPrefix STRING PREFIX
```

#### lower
转换所有的字母为小写
```
lower INPUT
```

#### upper
转换所有的字母为大写
```
upper INPUT
```

#### replace
替换所有匹配的字符串为新的字符串
```
replace INPUT OLD NEW
```

#### replaceRE
根据正则进行字符串替换
```
replaceRE PATTERN REPLACEMENT INPUT
```

#### split
根据定界符对字符串拆分，返回一个字符串数组
```
split STRING DELIM
```

#### substr
截取字符串
```
substr STRING START [LENGTH]
```

#### trim
去掉字符串前后指定的字符
```
trim INPUT CUTSET
```

#### title
转换所有字符串为标题风格的字符串
```
title INPUT
```

#### truncate
比较智能的截取字符串：不会截断单词
```
truncate SIZE INPUT
```

#### strings.HasSuffix
判断一个字符串的后缀是否是某个字符串
```
strings.HasSuffix STRING SUFFIX
```

#### strings.Repeat
返回字符串INPUT被重复COUNT次组成的字符串
```
strings.Repeat INPUT COUNT
```

#### strings.RuneCount
返回字符串的大小，更准确，因为是按RUNE类型计算的，可以计算包含中文的字符串
```
strings.RuneCount INPUT
```

#### strings.TrimLeft
去掉字符串左边的`CUTSET`,然后返回新的字符串
```
strings.TrimLeft CUTSET STRING
```

#### strings.TrimPrefix
和`strings.TrimLeft`不同，`strings.TrimPrefix`要求是前缀全字符串匹配的，而不是`strings.TrimLeft`一个个的字符匹配
```
strings.TrimPrefix PREFIX STRING
```

#### strings.TrimRight
去掉字符串右边的`CUTSET`,然后返回新的字符串
```
strings.TrimRight CUTSET STRING
```

#### strings.TrimSuffix
和`strings.TrimRight`不同，`strings.TrimSuffix`要求是后缀全字符串匹配的，而不是`strings.TrimRight`一个个的字符匹配
```
strings.TrimSuffix SUFFIX STRING
```

---

#### dateFormat
日期时间格式化，把时间（字符串）格式化为指定的格式
```
dateFormat LAYOUT INPUT
```

#### now
返回当前时间
```
now
```

#### time
把一个时间戳字符串转换为`time.Time`
```
time INPUT
```

#### duration
把给定的数字转换为`time.Duration`
`Unit`可以是以下几种： nanosecond/ns, microsecond/us/µs, millisecond/ms, second/s, minute/m or hour/h.
```
duration UNIT NUMBER
```

---

#### htmlEscape
HTML中的特殊字符转义
```
htmlEscape INPUT
```

#### htmlUnescape
HTML中的特殊字符 反-转义
```
htmlUnescape INPUT
```

## 示例

SOHA使用示例参考 [SOHA 示例](example/)

## 致谢

特别感谢 [HUGO](https://github.com/gohugoio/hugo),HUGO内置了很多优秀的函数库，SOHA抽取了他们并进行了修剪和增强。

## License

SOHA is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.