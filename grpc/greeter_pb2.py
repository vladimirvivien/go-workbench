# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeter.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='greeter.proto',
  package='protobuf',
  syntax='proto3',
  serialized_pb=_b('\n\rgreeter.proto\x12\x08protobuf\"\x1c\n\x0cGreetRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\" \n\rGreetResponse\x12\x0f\n\x07message\x18\x01 \x01(\t2E\n\x07Greeter\x12:\n\x05Greet\x12\x16.protobuf.GreetRequest\x1a\x17.protobuf.GreetResponse\"\x00\x62\x06proto3')
)




_GREETREQUEST = _descriptor.Descriptor(
  name='GreetRequest',
  full_name='protobuf.GreetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='protobuf.GreetRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=55,
)


_GREETRESPONSE = _descriptor.Descriptor(
  name='GreetResponse',
  full_name='protobuf.GreetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='protobuf.GreetResponse.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=89,
)

DESCRIPTOR.message_types_by_name['GreetRequest'] = _GREETREQUEST
DESCRIPTOR.message_types_by_name['GreetResponse'] = _GREETRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GreetRequest = _reflection.GeneratedProtocolMessageType('GreetRequest', (_message.Message,), dict(
  DESCRIPTOR = _GREETREQUEST,
  __module__ = 'greeter_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.GreetRequest)
  ))
_sym_db.RegisterMessage(GreetRequest)

GreetResponse = _reflection.GeneratedProtocolMessageType('GreetResponse', (_message.Message,), dict(
  DESCRIPTOR = _GREETRESPONSE,
  __module__ = 'greeter_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.GreetResponse)
  ))
_sym_db.RegisterMessage(GreetResponse)



_GREETER = _descriptor.ServiceDescriptor(
  name='Greeter',
  full_name='protobuf.Greeter',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=91,
  serialized_end=160,
  methods=[
  _descriptor.MethodDescriptor(
    name='Greet',
    full_name='protobuf.Greeter.Greet',
    index=0,
    containing_service=None,
    input_type=_GREETREQUEST,
    output_type=_GREETRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_GREETER)

DESCRIPTOR.services_by_name['Greeter'] = _GREETER

try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


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
          request_serializer=GreetRequest.SerializeToString,
          response_deserializer=GreetResponse.FromString,
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
            request_deserializer=GreetRequest.FromString,
            response_serializer=GreetResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'protobuf.Greeter', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaGreeterServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Greet(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaGreeterStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Greet(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    Greet.future = None


  def beta_create_Greeter_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('protobuf.Greeter', 'Greet'): GreetRequest.FromString,
    }
    response_serializers = {
      ('protobuf.Greeter', 'Greet'): GreetResponse.SerializeToString,
    }
    method_implementations = {
      ('protobuf.Greeter', 'Greet'): face_utilities.unary_unary_inline(servicer.Greet),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Greeter_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('protobuf.Greeter', 'Greet'): GreetRequest.SerializeToString,
    }
    response_deserializers = {
      ('protobuf.Greeter', 'Greet'): GreetResponse.FromString,
    }
    cardinalities = {
      'Greet': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'protobuf.Greeter', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
