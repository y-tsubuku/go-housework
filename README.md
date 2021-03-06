※作成中

# はじめに
こちらのソフトウェアは私のバックエンドスキルを証明するためのポートフォリオの1つです。

現時点では、フロントエンドやネイティブアプリなどのクライアントツールは存在せず、サービスの一般公開もしておりません。
ローカルでの動作確認までとなっております。あらかじめご了承ください。

# サービス概要
SmartHouseWork

ご家庭の家事を管理するソフトウェアです。

「誰が、いつ、何の家事」をやったかを管理することで、家事の負担を可視化することができます。

「自分ばかりが家事をしている」と悩んでいませんか？
まずは、「見える化」をしてみましょう！そこからパートナーへの相談に繋がります！

# IF
openapi/openapi.yamlをSwagger Editorなどでご確認ください。

# 動かしてみる
comming soon...

# このポートフォリオで証明できるスキル（現時点）
- サービスの企画
- 要件定義(機能要件のみ)
- REST-APIの設計(OpenAPI)
- RDBの論理設計、物理設計
- Goによる実装(UT含む)
- Dockerを使ったローカル環境の構築
- CIの構築
- Git操作
- 認証・認可(JWT)

# 技術スタック
Go/OpenAPI(Swagger)/Nginx/PostgreSQL/Docker/CircleCI/Gitなど

# 今後追加予定のもの
- CloudFormationを使ったインフラ構築
- デプロイパイプラインの構築
- React+TypeScriptを使ったフロントエンド開発
- ゲストユーザでのログイン機能（各種操作権限無し）
- パスワード変更機能
- スマートスピーカーを使ったUIの追加
- サービスの一般リリース

# develop info
## mockgenの例
```
$ cd app/server/repository
$ mockgen -source user.go -destination mock_user.go -package repository
```