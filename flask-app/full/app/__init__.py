# rawww flask template
# from flask import Flask
# from flask import jsonify

# # import logging
# # logger = logging.getLogger(__name__)
# # logger.addHandler(logging.StreamHandler(stream=sys.stderr))
# # logger.setLevel(logging.DEBUG)


# app = Flask(__name__)


# @app.route("/")
# def main():

#     return jsonify({"message": "somestuff"})


# if __name__ == "__main__":

#     app.run(debug=True)

#####################################################################################
# using flask_restful
from flask import Flask, jsonify, request
from flask_restful import Resource, Api

# from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate, migrate
from sqlalchemy import select
import logging

# from api.api import Hello
import os

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
# creating the flask app
app = Flask(__name__)
# creating an API object
apis = Api(app)
# adding configuration for using a sqlite database
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///site.db"
if __name__ != "__main__":
    gunicorn_logger = logging.getLogger("gunicorn.error")
    app.logger.handlers = gunicorn_logger.handlers
    app.logger.setLevel(gunicorn_logger.level)
from app.models.database import db

# Creating an SQLAlchemy instance
# db = SQLAlchemy(app)
migrate = Migrate(app, db)


# another resource to calculate the square of a number


from app.api.api import Hello

# adding the defined resources along with their corresponding urls
apis.add_resource(Hello, "/")


# driver function
if __name__ == "__main__":

    app.run(debug=True)
