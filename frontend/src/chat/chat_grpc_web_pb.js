/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chat = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatClient =
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
proto.chat.ChatPromiseClient =
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
 *   !proto.chat.SubscribeRequest,
 *   !proto.chat.SubscribeResponse>}
 */
const methodDescriptor_Chat_Subscribe = new grpc.web.MethodDescriptor(
  '/chat.Chat/Subscribe',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.chat.SubscribeRequest,
  proto.chat.SubscribeResponse,
  /**
   * @param {!proto.chat.SubscribeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.SubscribeResponse.deserializeBinary
);


/**
 * @param {!proto.chat.SubscribeRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.SubscribeResponse>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.Chat/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Chat_Subscribe);
};


/**
 * @param {!proto.chat.SubscribeRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.SubscribeResponse>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatPromiseClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.Chat/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Chat_Subscribe);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.AddMessageRequest,
 *   !proto.chat.AddMessageResponse>}
 */
const methodDescriptor_Chat_AddMessage = new grpc.web.MethodDescriptor(
  '/chat.Chat/AddMessage',
  grpc.web.MethodType.UNARY,
  proto.chat.AddMessageRequest,
  proto.chat.AddMessageResponse,
  /**
   * @param {!proto.chat.AddMessageRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.AddMessageResponse.deserializeBinary
);


/**
 * @param {!proto.chat.AddMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.chat.AddMessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.AddMessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.addMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/AddMessage',
      request,
      metadata || {},
      methodDescriptor_Chat_AddMessage,
      callback);
};


/**
 * @param {!proto.chat.AddMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.AddMessageResponse>}
 *     Promise that resolves to the response
 */
proto.chat.ChatPromiseClient.prototype.addMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.Chat/AddMessage',
      request,
      metadata || {},
      methodDescriptor_Chat_AddMessage);
};


module.exports = proto.chat;

