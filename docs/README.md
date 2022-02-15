# l4go/list ライブラリ

可変長の連結リストを提供するライブラリーです。  
状況によらず、メモリ開放処理を一貫性した書き方で実装できるので、メモリ解放をプログラマが管理しやすい作りになっています。

* [list.List](List.md)
	* 片方向連結リストを提供するモジュールです。
	* [list.Node](Node.md)が使われています。
* [list.DoubleList](List.md)
	* 双方向連結リストを提供するモジュールです。
	* [list.BiNode](BiNode.md)が使われています。
* [list.FifoList](FifoList.md)
	* FIFO操作が可能な連結リストを提供するモジュールです。  
	* [list.Node](Node.md)が使われています。
* [list.Node](Node.md)
	* 片方向連結なノードを提供するモジュールです。
* [list.BiNode](BiNode.md)
	* 双方向連結なノードを提供するモジュールです。
* [list.Iterator interface](Iterator.md)
	* このライブラリで提供する各種連結リストに対して逐次操作行うためのinterfaceです。
* [list.Value interface](Value.md)
	* このライブラリで提供する各種構造から要素の取得操作をするためのinterfaceです。
