# type List

片方向連結リストを提供するモジュールです。[type Node](Node.md)が使われています。  
連結順でのノード操作は、`Front()`で生成した[list.Iterator](Iterator.md) interfaceを利用します。

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## サンプルコード

[example](../examples/ex_single/ex_single.go)

## メソッド概略
### func NewList(free func(interface{})) \*List
双方向連結リスト(`*List`)を生成します。
free引数で、ノード削除時の処理を指定できます。消去時に特殊な処理が必要ない場合はnilを指定します。

### func (lst \*List) Move() \*List
全要素を譲渡した`*List`を生成し、この操作をしたリストを空にします。

主に、リストを他の処理へ譲渡する場合の、メモリ解放処理のコードを単純化するために使います。

`Move()`を使った後に譲渡することで、多重なメモリ解放が防げます。
これにより`defer`での、条件判断無しに`Clear()`呼び出しが可能になります。

### func (lst \*List) IsEmpty() bool
リストの要素が空の時にtrueを、それ以外の時はfalseをかえします。

### func (lst \*List) Clear()
保持しているノードを全て開放します。
NewList()で指定された初期化処理も合わせて実行します。

### func (lst \*List) PopFront() (interface{}, bool)
先頭の要素を取り出します。

### func (lst \*List) PushFront(val interface{}) \*Node
先頭に要素を追加します。

### func (lst \*List) Front() Iterator
先頭要素から始まる正方向の[Iterator interface](Iterator.md)を生成します。

