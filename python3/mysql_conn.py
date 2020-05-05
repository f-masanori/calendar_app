import mysql.connector
import config
import requests
import json

class MySQL():
    def __init__(self, config):
        """
        :param config: 接続設定を格納した辞書
        """
        self.config = config
        self.conn = None
        if config is not None:
            self.connect()

    def connect(self, config=None):
        """
        MySQLに接続
        """
        if config is None:
            config = self.config
        conn = mysql.connector.connect(**config)
        self.conn = conn
        return conn
