import os
from os.path import join, dirname
from dotenv import load_dotenv

dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)

host = os.environ.get("HOST")
port = os.environ.get("PORT")
user = os.environ.get("USER")
password = os.environ.get("PASSWORD")

webhookURL = os.environ.get("SLACK_WEBHOOK_TEST_SLACKBOT")


