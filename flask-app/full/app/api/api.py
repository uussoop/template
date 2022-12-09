from app.models.models import Profile
from flask_restful import Resource, Api
from flask import Flask, jsonify, request
from sqlalchemy import select
from app.models.models import Event
from app.models.database import db
from flask_apispec import marshal_with, doc, use_kwargs
from flask_apispec.views import MethodResource
from marshmallow import Schema, fields


# class EventSchema(Schema):
#     message = fields.Str(default="Success")


class HelloSchema(Schema):
    message = fields.Str(required=False)
    error = fields.Dict(required=False)


class Hello(Resource, MethodResource):
    @doc(description="this is where u can manipulate and get events", tags=["event"])
    @marshal_with(HelloSchema)
    def get(self):
        try:
            event = db.session.execute(select(Profile).filter_by(id=1)).one()[0]
            return {"message": "Hello, World!"}, 200
        except Exception as e:
            return {"error": e, "error": {"code": 500, "trace": "some random"}}, 200

    # Corresponds to POST request
    @use_kwargs(HelloSchema, location=("json"))
    @marshal_with(HelloSchema)
    def post(self, **kwargs):

        # data = request.get_json()  # status code
        return {"message": "sth", "error": {"code": 500, "trace": "sorandom"}}, 200
