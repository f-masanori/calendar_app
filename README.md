# mynikki_dev

## 開発環境
- mac mojave 10.14.6
- Docker version 19.03.4
- docker-compose version 1.24.1
## 実行
- /mynikki_dev/docker 内で docker-compose up
- ポート8080
## API 
- /users (GET)
    - req : なし
    - res : usersテーブルから全データを持ってくる(json)
- /app (POST)
    - userテーブルのnameを追加
    - req : json("name")
      - 例 {"name":"nmasanori"}
    - res : json("uuid","name")
    
- /nikkis (GET)
    - req : なし
    - res : nikkisテーブルから全データを持ってくる(json)
- /nikkis (POST)
    - req : json("UserId", "Date", "Title", "Content")
    - res : json("Id", "UserId", "Date", "Title", "Content")
## DB命名規則

- テーブル名 ... 複数形
- カラム ... 基本単数系

## DB 



##### 動いてるコンテナへの入り方
- $ docker exec -i -t コンテナ名 bash
##### 動いているコンテナ内でのコマンドの実行
- $ docker exec -i  63b7de01ee21 /bin/bash -c "cd ./dbseedgo && ls"
  - このようにすれば、コンテナ内のどこのディレクトリ内でもコマンドの実行ができる
### DB migration
- sql-migrateを使用

  -  https://github.com/rubenv/sql-migrate
- 参考

  - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 設定ファイル

  - docker/go/dbconfig.yml
  - DBの種類別の設定方法
    - https://github.com/rubenv/sql-migrate/blob/master/test-integration/dbconfig.yml
    - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 基本操作 (dbconfig.ymlのある場所で行う)
    - migrationファイルの新規作成
      - $ sql-migrate new create_テーブル名(複数形);
        - ex) $ sql-migrate new create_users; 
        - docker/go/migrations に  作成日付時間-create_users.sql という雛形sqlファイルが作成されるためその中SQL文を書き込む
    - migrationの実行
      - $ sql-migrate up
    - migration ドライラン
      - $ sql-migrate up -dryrun
    - migrationのroll back(ロールバック)
      - $ sql-migrate down
    - migrationの実行状態確認
      -  $ sql-migrate status
##### mysql接続
- $ mysql -u root -p 
- $ mysql -h 127.0.0.1 -P 3306 -u root -p mysql ローカルからの入り方
    - mysqlサーバーに入ってから　env　コマンドで環境変数確認
    - $ show columns from TABLENAME; テーブル構成確認

- $ mysql -h db -P 3306 -u root -p 違うコンテナからの入りかた

#### 設定
- ~~mysql_docker/mysql/initdb 内のsqlファイルを初回のみ実行~~
  - ↑migrationツールを入れるならば必要ない。
- ~~config.ymlにDBなどの設定を記入(追加したときはconfig.yml,conf.goにそれぞれ追加)~~
  - ↑構造体に直接記入

##### 二つの違いは？
- $ docker-compose build → imageの構築
- $ docker-compose up → image・コンテナの構築& コンテナの起動

#### アーキテクチャメモ

- クリーンアーキテクチャ使用(正しい構成なのかはわからない)

- infrastructure/router で ルーティング
- infrastructure/router の　userHandler := handlers.NewUserHandler(database.NewSqlHandler())　でuserHandlerの実体作成.userHandlerをレシーバーとするメソッドがそれぞれのハンドラー(コントローラ)

#### mockDBについて
- ~~go-sqlmock~~
  - ↑ネット上に参考が少ないため使用しない
- https://qiita.com/gold-kou/items/cb174690397f651e2d7f

### テスト
##### 標準パッケージのtestingを使用する

-   Services(アプリケーションロジック)のテスト
  1. dockerコンテナを立ち上げてexecとかで中に入る
  2. go test -run Get とかでGetと名のつく関数のテストを実行
    - ex)go test -run Get で func TestGetAllSuccessが実行
#### メモ
- gitでcommit メッセージを間違えた時
  - git commit --amend -m "書き直しメッセージ"
  - これで直前のcommitしたメッセージを変更できる
  - 参考(https://www.granfairs.com/blog/staff/git-commit-fix)
  
- gitでcommit を間違えた時
  
  - git reset --hard HEAD^ で直前のcommitを削除
  
- git で commit のログ確認
  
- git log
  
- go modはいづれ導入

  _________


- Go で int64 を int に変換するには int という関数を使う。

  ```
  b = int(a)
  ```

  これで int64 の a を値そのままで int 型の b に変換できる。
  
- 構造体に&(アドレス演算子）を使って初期化したり、newキーワードを使用すると**ポインタ型**で受け取ることができます。(関数の返しの時などの型に注意する)


