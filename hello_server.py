from concurrent import futures
import logging

import grpc

import hello_pb2
import hello_pb2_grpc

"""
from google.protobuf import json_format
req = json.loads(json_format.MessageToJson(request)) # Convert Request message into Json/Dictionary
json_format.Parse(json.dumps(response), user_pb2.GetUserResponse(), ignore_unknown_fields=False) # Convert Json/Dictionary into Response message

import "google/protobuf/struct.proto";
google.protobuf.Struct user = 2;
"""

class Greeter(hello_pb2_grpc.GreeterServicer):

    def SayHello(self, request, context):
        print("Received %s"% request.name)
        return hello_pb2.HelloReply(message='Hello, %s!' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    hello_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('localhost:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    print("Server listening on port 50051")
    serve()
    
    
