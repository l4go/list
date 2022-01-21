# type Node

片方向連結が可能なノードを提供するモジュールです。  

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## メソッド概略
### func (cur \*BiNode) InsertNext(val interface{}) \*BiNode
次の要素を追加します。

### func (cur \*BiNode) InsertPrev(val interface{}) \*BiNode
前の要素を追加します。

### func (cur \*BiNode) Value() interface{}
この要素の値を取得します。

### func (cur \*BiNode) Remove() interface{}
この要素を削除します。

### func AsForwardIterator(b, e \*BiNode) Iterator
このノードから始まる順方向の[list.Iterator interface](Iterator.md)を生成します。

### func AsReverseIterator(b, e \*BiNode) Iterator
このノードから始まる逆方向の[list.Iterator interface](Iterator.md)を生成します。
