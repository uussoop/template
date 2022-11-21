FROM python:3.8-alpine

RUN mkdir -p /home/brain/app/
WORKDIR /home/brain/app/
COPY ./requirements.txt .
COPY . .
RUN pip install --no-cache-dir -r requirements.txt
# RUN flask db init
# RUN flask db migrate
# RUN flask db upgrade
# COPY . .

