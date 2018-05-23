package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
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

type Decoder = func(r io.Reader) (image.Image, error)
type Encoder = func(w io.Writer, img image.Image) error

var decoders = map[string]Decoder{
	"jpeg": jpeg.Decode,
	"jpg":  jpeg.Decode,
	"png":  png.Decode,
	"gif":  gif.Decode,
}

var inputFormat = flag.String("in", "jpg", "使い方")
var outputFormat = flag.String("out", "png", "使い方")

func main() {
	flag.Parse()
	path := flag.Args()[0]

	if len(flag.Args()) < 1 {
		os.Exit(1)
	}

	if path == "" {
		os.Exit(1)
	}

	fmt.Println(*inputFormat)

	decoder, ok := decoders[*inputFormat]
	if !ok {
		fmt.Println("invalid format")
		os.Exit(1)
	}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if ext := filepath.Ext(path); ext == "."+*inputFormat {
			f, err := os.Open(path)
			defer f.Close()
			if err != nil {
				return err
			}

			img, err := decoder(f)
			if err != nil {
				return err
			}

			pngFilePath := path[0:len(path)-len(ext)] + ".png"
			pngFile, err := os.Create(pngFilePath)
			defer pngFile.Close()
			if err != nil {
				return err
			}

			if err = png.Encode(pngFile, img); err != nil {
				return err
			}
			fmt.Println("created: " + pngFilePath)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
