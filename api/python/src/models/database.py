import mysql.connector
import os

def con():
  return mysql.connector.connect(
    host = os.environ['DB_HOST'],
    user = os.environ['DB_USER'],
    password = os.environ['DB_PASSWORD'],
    database = os.environ['DB_DATABASE'],
    port = os.environ['DB_PORT'],
  )