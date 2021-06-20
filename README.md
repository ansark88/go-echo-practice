## 概要

GoのWebAPIの練習。Echoを使用。

過去に参加した[Go言語 初心者向けハンズオン](https://techdo.connpass.com/event/100306/)の課題をEchoでやってみる。

## ビルド時の注意

wireを使っている都合でwire_gen.goを含める必要がある。

`$ go run main.go wire_gen.go `

## 環境

- Go v1.16
- [Echo](https://echo.labstack.com/) v4.2.1
- [Air](https://github.com/cosmtrek/air) v1.12.1 (ホットリロード)
- [validator](https://github.com/go-playground/validator) v9.31.0 (リクエストの検証)
- [wire](https://github.com/google/wire) v0.5.0 (DI)

## 参考

以下を参考にプロジェクト構成を作成。ありがとうございます。

- [今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう。（golang）](https://qiita.com/tono-maron/items/345c433b86f74d314c8d)
- [【Go + レイヤードアーキテクチャー】DDDを意識してWeb APIを実装してみる](https://yyh-gl.github.io/tech-blog/blog/go_web_api/)
- [HON9LIN/go-echo-boilerplate](https://github.com/HON9LIN/go-echo-boilerplate)