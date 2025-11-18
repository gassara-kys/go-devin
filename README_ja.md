# go-devin

Devin API v1 を Go から扱うための SDK です。`devin` ルートパッケージでクライアントを生成し、`pkg/sessions` や `pkg/secrets` などのパッケージで各エンドポイントを型安全に利用できます。

## 特長

- `devin.NewClient` で API キーを渡すだけのシンプルな初期化。HTTP クライアントやリトライ設定、ロガーもオプションで差し替え可能。
- エンドポイントごとに `ListSessionsRequest` のような専用リクエスト／レスポンス構造体を定義し、公式ドキュメントと 1 対 1 に対応。
- `internal/httpclient` に共通の HTTP 実行処理を集約し、408/425/429/5xx の自動リトライや Bearer 認証、slog ログを提供。
- Gin 互換のバリデーションと `google/go-querystring` によるクエリ生成を採用。
- `examples/<domain>/<endpoint>` ディレクトリに `DEVIN_API_KEY=xxx go run .` でそのまま動くサンプルを多数同梱。

## インストール

```bash
go get github.com/gassara-kys/go-devin
```

## クイックスタート

```go
package main

import (
    "context"
    "fmt"
    "os"

    devin "github.com/gassara-kys/go-devin"
    "github.com/gassara-kys/go-devin/pkg/sessions"
)

func main() {
    apiKey := os.Getenv("DEVIN_API_KEY")
    if apiKey == "" {
        panic("DEVIN_API_KEY is not set")
    }

    client, err := devin.NewClient(apiKey)
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    summaries, err := client.Sessions.List(ctx, &sessions.ListSessionsRequest{Limit: 5})
    if err != nil {
        panic(err)
    }
    for _, sess := range summaries.Sessions {
        fmt.Println(sess.SessionID)
    }
}
```

## サンプル

```bash
DEVIN_API_KEY=xxx go run ./examples/sessions/list
DEVIN_API_KEY=xxx DEVIN_SESSION_ID=devin-123 go run ./examples/sessions/send_message
```

## 開発

```bash
make lint
make test
make build
```

## ライセンス

MIT
