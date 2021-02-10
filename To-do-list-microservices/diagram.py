from diagrams import Cluster, Diagram
from diagrams.onprem.compute import Server
from diagrams.onprem.client import Users
from diagrams.onprem.container import Docker
from diagrams.generic.database import SQL


with Diagram("Todo microservice", show=False):
    auth = Server("auth")
    todo = Server("todo")
    users = Users("users")
    web = Server("web")
    auth_db = SQL("auth db")
    todo_db = SQL("todo db")

    web >> auth >> auth_db
    web >> todo >> todo_db
    users >> web

