# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import greeter_pb2 as greeter__pb2


class GreeterStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Greet = channel.unary_unary(
        '/protobuf.Greeter/Greet',
        request_serializer=greeter__pb2.GreetRequest.SerializeToString,
        response_deserializer=greeter__pb2.GreetResponse.FromString,
        )


class GreeterServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Greet(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_GreeterServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Greet': grpc.unary_unary_rpc_method_handler(
          servicer.Greet,
          request_deserializer=greeter__pb2.GreetRequest.FromString,
          response_serializer=greeter__pb2.GreetResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'protobuf.Greeter', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
