# type DoubleList

双方向連結リストを提供するモジュールです。[type BiNode](BiNode.md)が使われています。  
連結順でのノード操作は、`Front()`および`Back()`で生成した[list.Iterator](Iterator.md) interfaceを利用します。

## import
```go
import "github.com/l4go/list"
```
vendoringして使うことを推奨します。

## サンプルコード

[example](../examples/ex_double/ex_double.go)

## メソッド概略
### func NewDoubleList(free func(interface{})) \*DoubleList
双方向連結リスト(`*DoubleList`)を生成します。
free引数で、ノード削除時の処理を指定できます。消去時に特殊な処理が必要ない場合はnilを指定します。

### func (lst \*DoubleList) Move() \*DoubleList
全要素を譲渡した`*DoubleList`を生成し、この操作をしたリストを空にします。

主に、リストを他の処理へ譲渡する場合の、メモリ解放処理のコードを単純化するために使います。

`Move()`を使うことで、リストを他の処理に譲渡した場合の、`Clear()`による多重なメモリ解放が防げます。
これにより`defer`などでの、無条件な`Clear()`呼び出しが可能になります。(結果的にソースコードをシンプルにすることが出来ます。)

### func (lst \*DoubleList) IsEmpty() bool
リストの要素が空の時にtrueを、それ以外の時はfalseをかえします。

### func (lst \*DoubleList) Clear()
保持しているノードを全て開放します。
NewList()で指定された初期化処理も合わせて実行します。

### func (lst \*DoubleList) PopFront() (interface{}, bool)
先頭の要素を取り出します。

### func (lst \*DoubleList) PopBack() (interface{}, bool)
末尾の要素を取り出します。

### func (lst \*DoubleList) PushFront(val interface{}) \*BiNode
先頭に要素を追加します。

### func (lst \*DoubleList) PushBack(val interface{}) \*BiNode
末尾に要素を追加します。

### func (lst \*DoubleList) Front() Iterator
先頭要素から始まる末尾方向へ辿る[Iterator interface](Iterator.md)を生成します。

### func (lst \*DoubleList) Back() Iterator
末尾要素から始まる船長方向へ辿る[Iterator interface](Iterator.md)を生成します。


