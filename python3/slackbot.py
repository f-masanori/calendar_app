import mysql.connector 
import config
 
try:
    connection = mysql.connector.connect(
    host=config.host,
    port=config.port,
    user=config.user,
    password=config.password,
    database='app'
    )
    print(connection.is_connected())
    cursor = connection.cursor()


    cursor.execute('select * from users')
    rows = cursor.fetchall()
    for i in rows:
        print(i)
    # DB操作が終わったらカーソルとコネクションを閉じる
    cursor.close()
    connection.close()
except mysql.connector.Error as err:
    print("Something went wrong: {}".format(err))
