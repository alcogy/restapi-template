from models import database

# 全件取得
def get_all_samples():
  con = database.con()
  cursor = con.cursor()
  query = 'SELECT * FROM my_table;'
  cursor.execute(query)
  data = cursor.fetchall()
  cursor.close()
  con.close()
  return data

# ID指定の１件取得
def get_sample_by_id(id):
  con = database.con()
  cursor = con.cursor()
  query = "SELECT * FROM my_table WHERE ID = %s;"
  cursor.execute(query, tuple(id))
  data = cursor.fetchall()
  cursor.close()
  con.close()
  return data

# 登録
def post_sample(data):
  con = database.con()
  cursor = con.cursor()
  query = "INSERT INTO my_table VALUES (0, %s, %s, now());"
  data = (data["name"], data["value"])
  cursor.execute(query, data)
  last_incert_id = cursor.lastrowid
  con.commit()
  cursor.close()
  con.close()
  return last_incert_id

# 更新
def put_sample(data):
  con = database.con()
  cursor = con.cursor()
  query = "UPDATE my_table SET name = %s, value = %s WHERE id = %s;"
  data = (data["name"], data["value"], data["id"])
  cursor.execute(query, data)
  con.commit()
  cursor.close()
  con.close()

# 削除
def delete_sample(id):
  con = database.con()
  cursor = con.cursor()
  query = "DELETE FROM my_table WHERE id = %s;"
  cursor.execute(query, tuple(id))
  con.commit()
  cursor.close()
  con.close()