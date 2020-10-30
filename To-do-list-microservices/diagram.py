from diagrams import Cluster, Diagram
from diagrams.onprem.compute import Server
from diagrams.onprem.client import Users
from diagrams.onprem.container import Docker
from diagrams.onprem.database import MongoDB


with Diagram("Todo microservice", show=False):
    api = Server("auth")
    users = Users("users")
    web = Server("web")
    auth_db = MongoDB("auth db")

    users >> api >> auth_db

    users >> web