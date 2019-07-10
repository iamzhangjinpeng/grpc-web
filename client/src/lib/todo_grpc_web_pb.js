/**
 * @fileoverview gRPC-Web generated client stub for todo
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


// var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.todo = require('./todo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.todoServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.todoServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.listTodoParams,
 *   !proto.todo.ListTodoResponse>}
 */
const methodInfo_todoService_listTodo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.ListTodoResponse,
  /** @param {!proto.todo.listTodoParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.ListTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.listTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.ListTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.ListTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todoServiceClient.prototype.listTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todoService/listTodo',
      request,
      metadata || {},
      methodInfo_todoService_listTodo,
      callback);
};


/**
 * @param {!proto.todo.listTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.ListTodoResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.todoServicePromiseClient.prototype.listTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todoService/listTodo',
      request,
      metadata || {},
      methodInfo_todoService_listTodo);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.addTodoParams,
 *   !proto.todo.addTodoResponse>}
 */
const methodInfo_todoService_addTodo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.addTodoResponse,
  /** @param {!proto.todo.addTodoParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.addTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.addTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.addTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.addTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todoServiceClient.prototype.addTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todoService/addTodo',
      request,
      metadata || {},
      methodInfo_todoService_addTodo,
      callback);
};


/**
 * @param {!proto.todo.addTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.addTodoResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.todoServicePromiseClient.prototype.addTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todoService/addTodo',
      request,
      metadata || {},
      methodInfo_todoService_addTodo);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.getTodoParams,
 *   !proto.todo.getTodoResponse>}
 */
const methodInfo_todoService_getTodo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.getTodoResponse,
  /** @param {!proto.todo.getTodoParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.getTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.getTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.getTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.getTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todoServiceClient.prototype.getTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todoService/getTodo',
      request,
      metadata || {},
      methodInfo_todoService_getTodo,
      callback);
};


/**
 * @param {!proto.todo.getTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.getTodoResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.todoServicePromiseClient.prototype.getTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todoService/getTodo',
      request,
      metadata || {},
      methodInfo_todoService_getTodo);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.deleteTodoParams,
 *   !proto.todo.deleteTodoResponse>}
 */
const methodInfo_todoService_deleteTodo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.deleteTodoResponse,
  /** @param {!proto.todo.deleteTodoParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.deleteTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.deleteTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.deleteTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.deleteTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todoServiceClient.prototype.deleteTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todoService/deleteTodo',
      request,
      metadata || {},
      methodInfo_todoService_deleteTodo,
      callback);
};


/**
 * @param {!proto.todo.deleteTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.deleteTodoResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.todoServicePromiseClient.prototype.deleteTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todoService/deleteTodo',
      request,
      metadata || {},
      methodInfo_todoService_deleteTodo);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.updateTodoParams,
 *   !proto.todo.updateTodoResponse>}
 */
const methodInfo_todoService_updateTodo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.updateTodoResponse,
  /** @param {!proto.todo.updateTodoParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.updateTodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.updateTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.updateTodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.updateTodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.todoServiceClient.prototype.updateTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.todoService/updateTodo',
      request,
      metadata || {},
      methodInfo_todoService_updateTodo,
      callback);
};


/**
 * @param {!proto.todo.updateTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.updateTodoResponse>}
 *     A native promise that resolves to the response
 */
proto.todo.todoServicePromiseClient.prototype.updateTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.todoService/updateTodo',
      request,
      metadata || {},
      methodInfo_todoService_updateTodo);
};


module.exports = proto.todo;
