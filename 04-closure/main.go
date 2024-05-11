package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	// 複数 defer 分があるときは，下から順に実行される
	defer fmt.Println("main func final-finish") // 最後
	defer fmt.Println("main func semi-finish")  // 2番目
	fmt.Println("hello world")                  // 最初
}

func trimExtension(files ...string) []string {
	// 可変長のstring型のスライスを引数として受け取り，stringのスライス型を返す関数．
	out := make([]string, 0, len(files))
	// range に stringのスライスを渡すと，インデックスと要素を返す．
	for _, f := range files {
		// TrimSuffix は第二引数に一致する末尾の文字列を削除した文字列を返す関数．
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	// 複数の返り値を設定したいときは，返り値の型を()で括る．
	// string を受け取り，stringとerror型の値を返す関数．

	// 標準パッケージのos.Openを使う．
	// os.Open はファイルが存在する場合は，ファイルオブジェクトへの参照が f に，nil が err に渡される．
	// os.Open はファイルが存在しない場合は，error が err に渡される．
	f, err := os.Open(name)
	if err != nil {
		// Go の errors というパッケージを利用して，エラー文を返す．
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func addExt(f func(file string) string, name string) {
	// 第一引数に無名関数を受け取る．
	fmt.Println(f(name))
}

func multiply() func(int) int {
	// 返り値が無名関数である関数
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	// 関数内で使う変数(クロージャー)
	// 外からアクセスできない．返り値にすることができる．
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

// global変数を使うと他の場所で書き換えられてしまう．
//var count int

func main() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))
	name1, err1 := fileChecker("main.go")
	// name2, err2 := fileChecker("file.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(name1)
	// 例外処理を書いておらず，実行が終わってしまうのでコメントアウトする
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return
	// }
	// fmt.Println(name2)

	// 無名関数を即座に実行する場合
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)

	// 無名関数を任意の場所で実行する場合．
	// f1が呼び出されたタイミングで無名関数が実行される．
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, "file1")

	// f3 には無名関数が入る．
	f3 := multiply()
	fmt.Println(f3(2))

	// f4には無名関数が入る．
	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}
