from concurrent import futures
import logging

import grpc

import user_pb2
import user_pb2_grpc

"""
from google.protobuf import json_format
req = json.loads(json_format.MessageToJson(request)) # Convert Request message into Json/Dictionary
json_format.Parse(json.dumps(response), user_pb2.GetUserResponse(), ignore_unknown_fields=False) # Convert Json/Dictionary into Response message

import "google/protobuf/struct.proto";
google.protobuf.Struct user = 2;
"""

class UserServicer(user_pb2_grpc.UserServiceServicer):

    def HealthCheck(self, request, context):
        print("Request Received ")
        return user_pb2.HealthCheckResponse(msg='Hello, Kapil!')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    user_pb2_grpc.add_UserServiceServicer_to_server(UserServicer(), server)
    server.add_insecure_port('localhost:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    print("Server listening on port 50051")
    serve()
    
    
