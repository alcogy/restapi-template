FROM python:3.11
WORKDIR /app
COPY ./api/python/docker/requirements.txt .
RUN pip install -r requirements.txt

EXPOSE 8080

CMD ["python", "main.py"]