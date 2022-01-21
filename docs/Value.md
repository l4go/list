# Value interface

このライブラリで提供する各種構造([Itorator interface](Iterator.md)、[type Node](Node.md)、[type BiNode](BiNode.md))から要素の取得操作をするためのinterfaceです。

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## メソッド概略
### func (it \*Iterator) Value() interface{}
Iteratorの示すノードの要素を返します。
