# pdf-tiff-document-text-detection
- GCPのCloudVisionAPIの一機能である[PDF/TIFF ドキュメントテキスト検出](https://cloud.google.com/vision/docs/pdf?hl=ja#vision-web-detection-gcs-go)を利用するサンプル

# 使い方
GCPのcredentialsを通している前提 : 
`go run main.go`

http://localhost:8091/ocr?pdf={PDF名}&bucket={BUCKET名}

http://localhost:8091/show?pdf={PDF名}&bucket={BUCKET名}

# result
## 1.json
生産性向上特別措置法【生産性革命法】及び産業競争力強化法等の一部を改正する法律の概要
https://www.meti.go.jp/policy/jigyou_saisei/seisanseisochihoukyoukahou/pdf/gaiyou-1.pdf  
のOCR化結果のjson一部

## 1-full.txt
1.jsonの中から`fullTextAnnotation`の文字列のみを抽出したもの

## 2.json
経済産業省: 生産性向上特別措置法【生産性革命法】及び産業競争力強化法等の一部を改正する法律の概要

https://www.meti.go.jp/policy/jigyou_saisei/seisanseisochihoukyoukahou/pdf/gaiyou-1.pdf
のOCR化結果のjson一部

## 2-full.txt
2.jsonの中から`fullTextAnnotation`の文字列のみを抽出したもの