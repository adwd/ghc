package main

import (
	"os"
	"fmt"
	"path/filepath"
)

/**
https://docs.google.com/presentation/d/1-20sWsmtmoFrHbmsuqIzSkfQKZn7Cdj1fuNmV1Btq0o/edit#slide=id.g37c8886736_0_1079

次の仕様を満たすコマンドを作って下さい
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる

以下を満たすように開発してください
- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
*/

func main() {
	path := os.Args[1]

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	if path == "" {
		os.Exit(1)
	}

	filepath.Walk(path, func (path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if ext := filepath.Ext(path); ext == ".jpg" || ext == ".jpeg" {
				fmt.Println(path)
			}
		}

		return nil
	})
}
