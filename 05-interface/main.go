package main

import (
	"fmt"
	"unsafe"
)

type controller interface {
	// インターフェースの中でメソッドを定義できる．
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed       int
	enginePower int
}
type bicycle struct {
	speed      int
	humanPower int
}

// vehicle オブジェクトの中身の値を変更したいので，ポインタレシーバーを使う．
func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

// この時点で，vehicle のデータ型は，controller インターフェース内のメソッドを全て実装したので，
// controller インターフェースを実装したと自動で見なされる．

func (b *bicycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bicycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}

// bicycleも同様

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

// vehicle で String メソッドを実装すると，標準パッケージの Stringer インターフェースを実装したことになる．
func (v vehicle) String() string {
	// Sprintf は formatで整えた文章をStringとして返す関数．標準出力はしない．
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func main() {
	// vehicle, bicycle の実体を作る．
	// v, b はそれぞれの実体の先頭アドレスを持つ．
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bicycle{0, 5}
	speedUpAndDown(b)
	// v はstinger インターフェースを実装しているのでこれでStringメソッドで実装したものが出力される．
	fmt.Println(v)

	// i1, i2 は一緒．
	var i1 interface{}
	var i2 any
	// サイズは16Byte
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)

}
func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
