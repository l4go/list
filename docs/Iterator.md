# Iterator interface

このライブラリで提供する各種連結リスト([type List](List.md)、[type List](DoubleList.md)、[type FifoList](List.md))に対して逐次操作行うためのinterfaceです。

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## サンプルコード

* [片方向リスト](../examples/ex_single/ex_single.go)
* [双方向リスト](../examples/ex_double/ex_double.go)
* [FIFOリスト](../examples/ex_fifo/ex_fifo.go)

## メソッド概略
### func (it \*Iterator) Clone() Iterator
Iteratorを複製します。複数の異なる操作を行いたいときに利用します。

### func (it \*Iterator) Next()
Iteratorの示すノードを次の位置に変更します。

### func (it \*Iterator) RemoveAndNext()
Iteratorの示すノードを削除しつつ、次の位置に移動します。

### func (it \*Iterator) Insert(val interface{})
Iteratorの示すノードの次の位置に、要素を追加します。

### func (it \*Iterator) IsEnd() bool
Iteratorを`Next()`でたどり終わったときにtrueを、それ以外の時はfalseを返します。

### func (it \*Iterator) Value() interface{}
Iteratorの示すノードの要素を返します。
