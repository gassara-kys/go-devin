# AGENTS.md

このドキュメントは `github.com/gassara-kys/go-devin` リポジトリで作業するAIエージェント向けのガイドです。設計背景や運用ルールを共有し、以降のタスクを一貫したスタイルで進められるようにします。

## 1. アーキテクチャ概要
- **ルートパッケージ**: `github.com/gassara-kys/go-devin`。`NewClient` を直接公開し、`Sessions/Secrets/Knowledge/Playbooks` サービスにアクセスできる。
- **サービス実装**: `pkg/<domain>` ディレクトリにドメインごとのリクエスト/レスポンス構造体とメソッドを配置。エンドポイントごとにファイルを分割し、メソッド名は `List`, `Create`, `Update` などに統一。
- **共通HTTP基盤**: `internal/httpclient` が API 呼び出し・リトライ・ログを司る。ルートクライアントもここを利用。
- **バリデーション**: Gin の `binding.Validator` を利用し、構造体タグは `binding:"required"` などを使用する。独自validatorオプションは廃止。
- **テスト補助**: `internal/testtransport.RoundTripFunc` と `internal/testutil.NewExecutor` で HTTP 通信をモック。サービス単体テストが同じ実行経路を通れるようにする。

## 2. コーディング規約
- すべてのパッケージテストは **テーブルドリブン形式**＋`go-cmp` で差分確認する。
- 新しいエンドポイントを追加する場合は、`pkg/<domain>/xxx.go` でエンドポイント単位のファイルを作成し、対応する `_test.go` を隣に置く。
- ルートクライアントから内部の `internal/...` パッケージへ直接アクセスしてよいが、外部利用者は `devin.NewClient` のみ扱う設計を崩さない。
- module path / import はすべて `github.com/gassara-kys/go-devin` に統一する。

## 3. 例とドキュメント
- `examples/<domain>/<endpoint>` に `package main` の実行サンプルを配置。`DEVIN_API_KEY`（必要に応じて追加の ID）を環境変数から読み込む形式とする。
- README は英語版 (`README.md`) と日本語版 (`README_ja.md`) を用意。ライセンスは MIT。
- サンプルコードや README 中の import 例は必ずルートパッケージ (`github.com/gassara-kys/go-devin`) を使用する。

## 4. ビルド/テスト
- `make lint` → golangci-lint、`make test` → `go test ./...`、`make build` → `go build ./...`。
- ユニットテストはローカル HTTP サーバを立てず、`testtransport.RoundTripFunc` または `testutil.NewExecutor` で完結させる。
- 公式APIを直接叩くテストは用意しない。必要であれば別リポジトリや手動手順で行う。

## 5. 作業フローのヒント
- 新規エンドポイント追加時: 型定義 → サービスメソッド → 例 (examples) → README 追記 → テーブルドリブンテストの順に進めると効率的。
- 既存コードを参照する際は `pkg/<domain>` 内のファイル命名（`list_sessions.go` など）に従うと理解しやすい。
- 変更後は必ず `gofmt -w`、`go test ./...` を実行してから PR/コミットに進む。

このドキュメントは随時更新し、将来のエージェントが設計意図を把握しやすいようにしてください。
