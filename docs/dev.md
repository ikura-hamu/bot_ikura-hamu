# dev

## 必要な環境

- Docker
- Docker Compose
- Task(推奨)

### Docker Compose の構成

- `bot` アプリケーション。 air によるホットリロード
- `db` MongoDB
- `express` MongoDB の UI

## 環境変数

|KEY|DEFAULT VALUE|目的|
|-|-|-|
|`BOT_ACCESS_TOKEN`||traQ botのアクセストークン|
|`BOT_VERIFICATION_TOKEN`||traQ botのverificationトークン|
|`BOT_ID`||traQ bot のid|
|`BOT_USER_ID`||traQ bot の user id|
|`NS_MONGODB_DATABASE`|`bot`|MongoDBのデータベース|
|`NS_MONGODB_HOSTNAME`|`db`|MongoDBのホスト名|
|`NS_MONGODB_PASSWORD`|`password`|MongoDBのパスワード|
|`NS_MONGODB_PORT`|`27017`|MongoDBのポート番号|
|`NS_MONGO_USER`|`root`|MongoDBのユーザー|

## 引数

指定しない場合は本番環境。`dev`とすると開発環境。

開発環境になると、ログの出力レベルと、クライアントが変わる。

||開発環境|本番環境|
|-|-|-|
|ログ(zap)|`Debug`以上|`Info`以上|
|クライアント|標準出力|traQ|
