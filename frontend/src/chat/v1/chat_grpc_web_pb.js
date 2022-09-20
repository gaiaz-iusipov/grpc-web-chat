/**
 * @fileoverview gRPC-Web generated client stub for chat.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chat = {};
proto.chat.v1 = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.v1.ChatClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.v1.ChatPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.v1.Subscribe.Request,
 *   !proto.chat.v1.Subscribe.Response>}
 */
const methodDescriptor_Chat_Subscribe = new grpc.web.MethodDescriptor(
  '/chat.v1.Chat/Subscribe',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.chat.v1.Subscribe.Request,
  proto.chat.v1.Subscribe.Response,
  /**
   * @param {!proto.chat.v1.Subscribe.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.v1.Subscribe.Response.deserializeBinary
);


/**
 * @param {!proto.chat.v1.Subscribe.Request} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.v1.Subscribe.Response>}
 *     The XHR Node Readable Stream
 */
proto.chat.v1.ChatClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.v1.Chat/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Chat_Subscribe);
};


/**
 * @param {!proto.chat.v1.Subscribe.Request} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.v1.Subscribe.Response>}
 *     The XHR Node Readable Stream
 */
proto.chat.v1.ChatPromiseClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.v1.Chat/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Chat_Subscribe);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.v1.AddMessage.Request,
 *   !proto.chat.v1.AddMessage.Response>}
 */
const methodDescriptor_Chat_AddMessage = new grpc.web.MethodDescriptor(
  '/chat.v1.Chat/AddMessage',
  grpc.web.MethodType.UNARY,
  proto.chat.v1.AddMessage.Request,
  proto.chat.v1.AddMessage.Response,
  /**
   * @param {!proto.chat.v1.AddMessage.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.v1.AddMessage.Response.deserializeBinary
);


/**
 * @param {!proto.chat.v1.AddMessage.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.chat.v1.AddMessage.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.v1.AddMessage.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.v1.ChatClient.prototype.addMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.v1.Chat/AddMessage',
      request,
      metadata || {},
      methodDescriptor_Chat_AddMessage,
      callback);
};


/**
 * @param {!proto.chat.v1.AddMessage.Request} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.v1.AddMessage.Response>}
 *     Promise that resolves to the response
 */
proto.chat.v1.ChatPromiseClient.prototype.addMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.v1.Chat/AddMessage',
      request,
      metadata || {},
      methodDescriptor_Chat_AddMessage);
};


module.exports = proto.chat.v1;

