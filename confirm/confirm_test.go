package confirm

import "testing"

// *_test.go ファイル がテストファイルとなる。

var Debug bool = false

// TestIsOne テストしたい関数名の前にTestつけている。
func TestIsOne(t *testing.T) {
	t.Log("これからテストするよ！！")

	i := 1
	if Debug {
		t.Skip("スキップ")
	}
	// 同階層であれば関数読み込む (confirm.go参照)
	v := IsOne(i)
	if v != true {
		t.Errorf("%v != %v", i, 1)
	}
}

// testは該当処理と同じ階層に作成することが多い
