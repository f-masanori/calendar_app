import mysql.connector 
import config
import requests
import json
import mysql_conn
import datetime
import schedule
import time
MySQLConfig = {
    'host': config.host,
    'user': config.user,
    'password': config.password,
    'database': 'app',
}
email = "test05051@test.com"
WEB_HOOK_URL = config.webhookURL

def getTodayStr():
    _today = datetime.date.today()
    print(_today)
    today = _today.strftime('%Y-%m-%d')
    return today
def main(): 
    mysql = mysql_conn.MySQL(MySQLConfig)
    connection = mysql.connect()
    # print(connection.is_connected())   
    date = getTodayStr()
    # print(type(date))
    cursor = connection.cursor()
    cursor.execute(
        'SELECT event,date FROM events AS E INNER JOIN users AS U ON E.uid=U.uid WHERE email=%s AND date=%s',
        (email, date))
    rows = cursor.fetchall()

    event=""
    for i in rows:
        event = i[0]
        print(i[0])
    # DB操作が終わったらカーソルとコネクションを閉じる
    cursor.close()
    connection.close()
    if event != "":   
        requests.post(WEB_HOOK_URL, data=json.dumps({
        'text': event,  # 通知内容
        'username': 'Bakira-Tech-Python-Bot',  # ユーザー名
        'icon_emoji': ':smile_cat:',  # アイコン
        'link_names': 1,  # 名前をリンク化
        }))


def job():
    main()

#AM11:00のjob実行を登録
schedule.every().day.at("20:00").do(job)
while True:
    schedule.run_pending()
    time.sleep(1)


