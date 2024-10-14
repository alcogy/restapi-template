from flask import *
from dotenv import load_dotenv
from services import sample

load_dotenv()

app = Flask(__name__)
app.json.enxure_ascii = False

@app.get("/")
def check():
  return jsonify({"status": "ok"})

@app.get("/sample")
def select():
  data = sample.get_all_samples()
  return jsonify(data)

@app.get("/sample/<id>")
def select_one(id):
  data = sample.get_sample_by_id(id)
  return jsonify(data)

@app.post("/sample")
def insert():
  content = request.json
  id = sample.post_sample(content)
  return jsonify({"insertID": id})

@app.put("/sample")
def update():
  content = request.json
  sample.put_sample(content)
  return jsonify({"message": "ok"})

@app.delete("/sample/<id>")
def delete(id):
  sample.delete_sample(id)
  return jsonify({"message": "ok"})


if __name__ == '__main__':
  app.run(host='0.0.0.0', port=8080, debug=True)
