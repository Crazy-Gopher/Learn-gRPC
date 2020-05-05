import grpc

import user_pb2

import user_pb2_grpc

def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)
        response = stub.HealthCheck(user_pb2.HealthCheckRequest())
    print(response.msg)
run()