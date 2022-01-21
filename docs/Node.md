# type Node

片方向連結が可能なノードを提供するモジュールです。  

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## メソッド概略
### func (cur \*Node) InsertNext(val interface{}) \*Node
次の要素を追加します。

### func (cur \*Node) RemoveNext() (interface{}, bool)
次の要素を削除します。

### func (cur \*Node) Value() interface{}
要素の値を取得します。

### func AsIterator(prev \*Node) Iterator
該当のノードから始まる正方向の[list.Iterator interface](Iterator.md)を生成します。
