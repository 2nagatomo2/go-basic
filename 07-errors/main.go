package main

import (
	"errors"
	"fmt"
	"os"
)

// 新たなerrorを定義する．（標準パッケージで最初がErrで始まるものは，センチネルエラーと呼ばれる．）
var ErrCustom = errors.New("not found")

func main() {
	// errors パッケージの New 関数を使って新しいエラーを作成できる．
	// New 関数は作成したerrorString 構造体の実体の先頭アドレスを返す．
	// 引数に文字列を与えることで，エラーメッセージを追加できる．
	err01 := errors.New("something wrong")
	err02 := errors.New("something wrong")
	// %p はポインタを表示する．
	// %v でerr01の値を出力すると，エラーメッセージが帰ってくる．
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	fmt.Println(err01)
	// err01 と err02 は別のアドレスに作られるので異なるもの
	fmt.Println(err01 == err02)

	// %w , fmt.Errorf でwrap
	// エラーメッセージに付加情報を追加できる．
	err0 := fmt.Errorf("add info: %w", errors.New("original error"))
	// wrapされたerrorは *errors.wraperror 型になる．
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	// errors.Unwrapメソッドで，wrapされる前のオリジナルエラーが得られる．
	fmt.Println(errors.Unwrap(err0))
	fmt.Printf("%T\n", errors.Unwrap(err0))

	// %v を使うと，*errors.errorString型のままになる．
	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)
	// errorString構造体にはunwrap関数は実装されていない
	// unwrapの実装されていない構造体にUnwrap関数を適用した場合，nilが返される．
	fmt.Println(errors.Unwrap(err1))

	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	// 付加情報は累積する．
	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)

	// Unwrap したものを比較する関数
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)
	if err3 != nil {
		// どこかの階層にos.ErrNotExist が存在するかを調べる．
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}

}
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
