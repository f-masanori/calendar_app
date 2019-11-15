# mynikki_dev

## 開発環境
- mac mojave 10.14.6
- Docker version 19.03.4
- docker-compose version 1.24.1
## 実行
- /mynikki_dev/docker 内で docker-compose up
- ポート8080
## API
- / (GET)
  
    - DB情報を表示
    
- /users (GET)
  
    - usersテーブルから全データを持ってくる(json)
    
- /app (POST)
    - userdテーブルのnameを追加
    
    - 送信方法はjsonでキーは "name"
    
    - 例 {"name":"nmasanori"}
    
## DB命名規則

- テーブル名 ... 複数形
- カラム...基本単数系

​## DB 



##### 動いてるコンテナへの入り方
- $ docker exec -i -t コンテナ名 bash
- run との違いは？
#### migrationツール
- sql-migrateを使用

  -  https://github.com/rubenv/sql-migrate
- 参考

  - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 設定ファイル

  - docker/go/dbconfig.yml
  - DBの種類別の設定方法
    - https://github.com/rubenv/sql-migrate/blob/master/test-integration/dbconfig.yml
    - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 基本操作
    - migrationファイルの新規作成
    - $ sql-migrate new create_テーブル名(複数形);
      - ex) $ sql-migrate new create_users;
      - docker/go/migrations に  作成日付時間-create_users.sql という雛形sqlファイルが作成されるためその中に書き込む
    - migrationの実行
      - $ sql-migrate up
    - migration ドライラン
      - $ sql-migrate up -dryrun
    - マイグレーションのロールバック
      - $ sql-migrate down
    - マイグレーションの実行状態確認
      -  $ sql-migrate status
##### mysql接続
- $ mysql -u root -p 
- $ mysql -h 127.0.0.1 -P 3306 -u root -p mysql ローカルからの入り方
    - mysqlサーバーに入ってから　env　コマンドで環境変数確認
    - $ show columns from TABLENAME; テーブル構成確認

- $ mysql -h db -P 3306 -u root -p 違うコンテナからの入りかた

#### 設定
- mysql_docker/mysql/initdb 内のsqlファイルを初回のみ実行
- ↑migrationツールを入れるならば必要ない。
- config.ymlにDBなどの設定を記入(追加したときはconfig.yml,conf.goにそれぞれ追加)
##### 二つの違いは？
- $ docker-compose build
- $ docker-compose up 

#### アーキテクチャメモ
- infrastructure/router で ルーティング
- infrastructure/router の　userHandler := handlers.NewUserHandler(database.NewSqlHandler())　でuserHandlerの実体作成.userHandlerをレシーバーとするメソッドがそれぞれのハンドラー(コントローラ)

