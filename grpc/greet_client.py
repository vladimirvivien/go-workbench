from __future__ import print_function

import grpc

import greeter_pb2 as pb

def run():
 channel = grpc.insecure_channel('localhost:50051')
 stub = pb.GreeterStub(channel)
 response = stub.Greet(pb.GreetRequest(name="Vladimir"))
 print(response.message)

if __name__ == '__main__':
  run()

  