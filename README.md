# calendar_app
- もともと、インターン先で開発していた予定表アプリでしたが、サーバーレスになることが決定し自分の著作になりました。上司からのアドバイス等を頂きながら、１から自分で開発しました。現在はreactで開発しているカレンダーアプリの[サーバーサイド](https://github.com/f-masanori/calendar_react)として勉強のため開発を進めています
- モバイルアプリを対象とした開発だったため、APIサーバーとして開発しています。
- Docker上で開発をしており、クリーンアーキテクチャっぽい？アーキテクチャにしました。(ネットの記事を読んだりしながら見様見真似でやってみたので正しいと言えるのかは不明)
- DB(Mysql)もDocker上で動かしています。
## ブランチ
## なぜ

- なぜgolangか
  - 標準ライブラリだけで、ある程度のサーバー開発ができる
  - 静的型付け言語である(好み)
- なぜDockerか
  - ローカル環境を汚さない
  - 環境差異の減少(ex VPSレンタルしてから、git cloneからdocker-compose upですぐに立ち上げることができる)

## 開発環境
- mac mojave 10.14.6
- Docker version 19.03.4
- docker-compose version 1.24.1
## 実行
- /docker 内で docker-compose up
  - port 8080: calendar app
  - port 3306: mySQL

## API実装

- /registerUser(POST)
  - 説明 : SignInした際にユーザー初期登録としてのAPI
  - req :  headers={ 'Content-Type': "application/json"}  ,body= json("UID","Email")
  - res : (未実装[resに何が必要かを学んでから実装])
- /addEvent(POST)
  - 説明 : ユーザー別にEventを追加するAPI(jwt認証)
  - req :headers: { 'Content-Type': "application/json",  'Authorization': idToken }, body= json("EventID", "Date", "InputEvent")
  - res : (未実装[resに何が必要かを学んでから実装])
- /getEventsByUID(GET)
  - 説明 : ユーザー別に全てのEventsを取得
- /deleteEvent(POST)
  - 説明 : 指定されたEventを削除
- /getNextEventID(GET)
  - 説明 : ユーザーが次に作成するEventのIDを取得

##  追加・改善したいアイデア

- slackと接続して通知させたい
- Linebotから操作

## 追加・改善したい実装

- ~~eventの編集・削除のAPIを追加~~

- トランザクションの実装

- ~~DB設計を見直す(Eventのバックカラーなどが入るため)~~

- ヘッダにログインメールアドレス記載

  ______

- 現在、EventIDをクライアント側で管理しているが、EventをDB追加時にDB内のEventID情報を取得して追加でいい？

 ## これから勉強するところ

- Dockerを用いたデプロイ(今は単純に開発環境でデプロイしている)
- golangの詳しい仕様・標準ライブラリの詳しい仕様
  - DBの取扱いをもう少し詳しくなる
- interface(よく理解しないまま使用してしまっている)



## ハマった点などとその解決策

- Golangでファイルを読み込む時、なぜか読み込めない
  - docker開発していることを忘れてローカルでのパスを参照してしまっていた。
  - docker内でのパスに変更→できた！
- VPS上でDockerでmysqlサーバーを立ち上げ時DBに保存できない時がある(ローカルでは正しく動いた)
  - 調べてみると、日本語が入っている場合だけ保存ができていない
  - ぐぐるとどうやら、文字コードが関係しているらしい
  - https://kitigai.hatenablog.com/entry/2019/03/03/203310を参考にして、VPSのMysqlコンテナ中に/tmp/dockerdir/custom.cnfを作成し、文字コードを設定
  - DB作り直して再起動→できた！
- ローカルで立ち上げたサーバーから、APIを叩けない
  - 調べてみるとCORSが関係してるらしい
  - goliraというライブラリで対処した
  - しかしプリフライトエラー
  - 試行錯誤ののち、goliraでルーティングする際にHTTPのメソッド指定するとエラーが出ることが判明(プリフライトがOPTIONSでくるから?)
  - goliraのメソッド指定無くして、CORS対処のミドルウェアを挟む→できた！
- クリーンアーキテクチャの記事を読んだりしたがほとんど理解できなかった
  - 入門系の記事を参考にして写経して自分で実装してみると、なんとなくイメージが湧いてきた（習うより慣れろを実感）
- firebase authenticationの鍵がパスの指定ミスでgitignoreに正しく追加されていない状態でリモートにpushしてしまった
  - firebase authenticationのプロジェクトを新しくすることで対応
  - 大事な鍵が入っているディレクトリでgitを扱うときは、リモートpushする前にgit stateなどで確認してミスを起こさないようにする

## DB 

- #### 命名規則
  
  - テーブル名 ... 複数形
  - カラム ... 基本単数系
  - 基本全て小文字のスネークケース
  
- #### DB seedについて

  - DBにテストデータを入れるために作成した
  - 詳しくはコードを読む
  
- #### mockDBについて

  - ~~go-sqlmock~~
    - ↑ネット上に参考が少ないため使用しない
  - https://qiita.com/gold-kou/items/cb174690397f651e2d7f
## DB migration
- sql-migrateを使用
  
  -  https://github.com/rubenv/sql-migrate
- 参考

  - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 設定ファイル

  - docker/go/dbconfig.yml
  - DBの種類別の設定方法
    - https://github.com/rubenv/sql-migrate/blob/master/test-integration/dbconfig.yml
    - https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
- 基本操作 (コンテナ中のdbconfig.ymlのある場所で行う)
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
      - $ sql-migrate status
    

## テスト[1つ以外未実装]

##### 標準パッケージのtestingを使用する

- Services(アプリケーションロジック)のテスト

  1. dockerコンテナを立ち上げてexecとかで中に入る

  2. go test -run Get とかでGetと名のつく関数のテストを実行

     - ex)go test -run Get で func TestGetAllSuccessが実行

       



## メモ(アーキテクチャ)
- クリーンアーキテクチャ使用(正しい構成なのかはわからない)

- infrastructure/router で ルーティング
- infrastructure/router の　userHandler := handlers.NewUserHandler(database.NewSqlHandler())　でuserHandlerの実体作成.userHandlerをレシーバーとするメソッドがそれぞれのハンドラー(コントローラ)

## メモ(Git)
- commit メッセージを間違えた時
  - git commit --amend -m "書き直しメッセージ"
  - これで直前のcommitしたメッセージを変更できる
  - 参考(https://www.granfairs.com/blog/staff/git-commit-fix)
  
- commit を間違えた時
  
  - git reset --hard HEAD~ で直前のcommitを削除
  
  - git reset --hard HEAD~2 で2個まえの ....
  
-  commit のログ確認

  - git log

- リモートリポジトリ(master)と同期

  - git pull origin master
  
-  git: ローカルをリモートブランチで上書き

  現状のローカルの状態はいらない時

  ```
$ git fetch origin
  $ git reset --hard origin/ブランチ名(masterとかmasa-devとか)
  ```

  
  
- ~~go modはいづれ導入したい~~した

  _________

## メモ(Golang)
- Go で int64 を int に変換するには int という関数を使う。

```
  b = int(a)
```

  これで int64 の a を値そのままで int 型の b に変換できる。

- 構造体に&(アドレス演算子）を使って初期化したり、newキーワードを使用すると**ポインタ型**で受け取ることができます。(関数の返しの時などの型に注意する)

______

- database/ sql 

  -  Query と Exec に分けて考える。 Query は副作用のない `SELECT`、Exec は副作用のある `INSERT` や、`UPDATE`、`DELETE` に当たる
  - トランザクションにてprepareを使うときは少し気をつけるhttps://precure-3dprinter.hatenablog.jp/entry/2018/11/22/Golang%E3%81%A7%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%82%92%E4%BD%BF%E3%81%86%E8%A9%B1](https://precure-3dprinter.hatenablog.jp/entry/2018/11/22/Golangでトランザクションを使う話)
  - トランザクションの冗長さを無くせるらしいhttps://qiita.com/miya-masa/items/316256924a1f0d7374bb

  ______

- go のstring int の互換

  - import("strconv")して
    
  - ```
    var i int
    var s string="123"
    i, _ = strconv.Atoi(s)
    fmt.Println(i) // -> 123
    ```
```

  - ```
    var i int=321
    var s string
    s = strconv.Itoa(i)
    fmt.Println(s) // -> "123"
```

______

- go の型情報取得
  
  - import("reflect")して
    
  - ```go
    fmt.Println(reflect.TypeOf(調べたい型)) 
    ```
  
_____


- docker上で開発する際にローカルのVScodeでコードの編集をしてもvscodeの拡張機能のフォーマッタが機能しないことがよくある
  
  - そのためターミナルで　gofmt -s -w ./ これでその配下にあるgoファイルを全て整形して保存してくれる
  - https://qiita.com/suin/items/9f9bdaa0cb9cb80cf752

### メモ(Linuxコマンド)

- 名前の変更

  - mv (旧ファイル名) (新ファイル名)

- ディレクトリの削除（警告なし）

  - rm -rf ()

- ディレクトリの送信

  - ```
    scp -r -i ~/.ssh/id_rsa ./$(DIR) user@109.222.000.000:$(PASS)
    ```

### メモ(docker)

- リアルタイムでログを見る
  
  - docker logs -f docker-id
  
- ##### 動いてるコンテナへの入り方

  - $ docker exec -i -t コンテナ名 bash

  ##### 動いているコンテナ内でのコマンドの実行

  - $ docker exec -i  63b7de01ee21 /bin/bash -c "cd ./dbseedgo && ls"
    - このようにすれば、コンテナ内のどこのディレクトリ内でもコマンドの実行ができる
  
- ##### mysql接続

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

### メモ(Mysql)

- mysql> CREATE DATABASE app; # DB作成
- mysql> show databases; #データベース一覧を表示
- 文字化けや日本語を保存できない時
  - 文字セット？に問題あり？

### メモ(JWTとは)

- JWT (Json Web Token)
- Tokenとはユーザーを識別するための認証情報。つまりJWTとは、JavaScriptのオブジェクトの形をした認証情報のことです。（https://techblog.roxx.co.jp/entry/2019/03/13/135739）
- JWT のメリット
  - 電子署名(勉強する)
- JWTの構成
  - JWTは大きく3つの要素で構成される。
    - ヘッダー
    - クレーム情報
    - 署名

### メモ(Base64とは)

- base64とは、64進数を意味する言葉で、すべてのデータをアルファベット(`a~z`, `A~z`)と数字(`0~9`)、一部の記号(`+`,`/`)の**64文字**で表すエンコード方式(https://qiita.com/PlanetMeron/items/2905e2d0aa7fe46a36d4)
- 

