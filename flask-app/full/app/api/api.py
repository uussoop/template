# making a class for a particular resource
# the get, post methods correspond to get and post requests
# they are automatically mapped by flask_restful.
# other methods include put, delete, etc.

from flask_restful import Resource, Api
from flask import Flask, jsonify, request

# from app.application import db
from sqlalchemy import select
from app.models.models import Profile

# from app import api, db
from app.models.database import db


class Hello(Resource):
    def get(self):
        print("hellooooooooooooooooooooooooooo")
        profile = db.session.execute(select(Profile).filter_by(id=1)).one()[0]
        print(profile)
        return jsonify({"messages": profile.first_name, "age": profile.age})

    # Corresponds to POST request
    def post(self):

        data = request.get_json()  # status code
        return jsonify({"data": data}), 201


#
# api.add_resource(Hello, "/")
#
