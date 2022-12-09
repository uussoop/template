from flask import Flask, jsonify, request
from flask_restful import Resource, Api
from flask_migrate import Migrate, migrate
from sqlalchemy import select
import logging
import os
from flask_apispec.extension import FlaskApiSpec


BASE_DIR = os.path.abspath(os.path.dirname(__file__))
# creating the flask app
app = Flask(__name__)
# creating an API object
apis = Api(app)
# adding configuration for using a sqlite database
docs = FlaskApiSpec(app)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///site.db"
if __name__ != "__main__":
    gunicorn_logger = logging.getLogger("gunicorn.error")
    app.logger.handlers = gunicorn_logger.handlers
    app.logger.setLevel(gunicorn_logger.level)
from app.models.database import db

migrate = Migrate(app, db)
from apispec import APISpec
from apispec.ext.marshmallow import MarshmallowPlugin

app.config.update(
    {
        "APISPEC_SPEC": APISpec(
            title="calender",
            version="v1",
            plugins=[MarshmallowPlugin()],
            openapi_version="2.0.0",
        ),
        "APISPEC_SWAGGER_URL": "/swagger",  # URI to access API Doc JSON
        "APISPEC_SWAGGER_UI_URL": "/swagger-ui",  # URI to access UI of API Doc
    }
)


# another resource to calculate the square of a number


from app.api.api import Hello

# adding the defined resources along with their corresponding urls
apis.add_resource(Hello, "/")
docs.register(Hello)


# driver function
if __name__ == "__main__":

    app.run(debug=True)
