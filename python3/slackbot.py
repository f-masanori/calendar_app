import mysql.connector 
import config
import requests
import json
import mysql_conn
try:
    MySQLConfig = {'host': config.host,
              'user': config.user,
              'password': config.password,
              'database': 'app',
              }
    mysql = mysql_conn.MySQL(MySQLConfig)
    connection = mysql.connect()
    print(connection.is_connected())
    cursor = connection.cursor()
    cursor.execute(
        'SELECT event,date FROM events AS E INNER JOIN users AS U ON E.uid=U.uid WHERE email="test05051@test.com"')
    rows = cursor.fetchall()

    event=""
    for i in rows:
        event = i[0]
        print(i[0])
    # DB操作が終わったらカーソルとコネクションを閉じる
    cursor.close()
    connection.close()
    # if event != "":
    #     WEB_HOOK_URL = config.webhookURL
    #     requests.post(WEB_HOOK_URL, data=json.dumps({
    #     'text': event,  # 通知内容
    #     'username': 'Bakira-Tech-Python-Bot',  # ユーザー名
    #     'icon_emoji': ':smile_cat:',  # アイコン
    #     'link_names': 1,  # 名前をリンク化
    #     }))
except mysql.connector.Error as err:
    print("Something went wrong: {}".format(err))
