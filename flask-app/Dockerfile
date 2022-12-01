FROM python:3.8-alpine

RUN mkdir -p /home/brain/app/
WORKDIR /home/brain/app/
COPY ./requirements.txt .
COPY full/ .
RUN pip install  -r requirements.txt
# RUN flask db init --no-cache-dir
# RUN flask db migrate
# RUN flask db upgrade
# COPY . .

