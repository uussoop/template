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
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate, migrate
from sqlalchemy import select

# creating the flask app
app = Flask(__name__)
# creating an API object
api = Api(app)
# adding configuration for using a sqlite database
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///site.db"

# Creating an SQLAlchemy instance
db = SQLAlchemy(app)
migrate = Migrate(app, db)
# Models
class Profile(db.Model):
    # Id : Field which stores unique id for every row in
    # database table.
    # first_name: Used to store the first name if the user
    # last_name: Used to store last name of the user
    # Age: Used to store the age of the user
    id = db.Column(db.Integer, primary_key=True)
    first_name = db.Column(db.String(20), unique=False, nullable=False)
    last_name = db.Column(db.String(20), unique=False, nullable=False)
    age = db.Column(db.Integer, nullable=False)

    # repr method represents how one object of this datatable
    # will look like
    def __repr__(self):
        return f"Name : {self.first_name}, Age: {self.age}"


# making a class for a particular resource
# the get, post methods correspond to get and post requests
# they are automatically mapped by flask_restful.
# other methods include put, delete, etc.


class Hello(Resource):

    # corresponds to the GET request.
    # this function is called whenever there
    # is a GET request for this resource
    def get(self):
        profile = db.session.execute(select(Profile).filter_by(id=1)).one()[0]
        print(profile)
        return jsonify({"messages": profile.first_name, "age": profile.age})

    # Corresponds to POST request
    def post(self):

        data = request.get_json()  # status code
        return jsonify({"data": data}), 201


# another resource to calculate the square of a number


# adding the defined resources along with their corresponding urls
api.add_resource(Hello, "/")


# driver function
if __name__ == "__main__":

    app.run(debug=True)
