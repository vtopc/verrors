# verrors

Yet another error wrapper package.
Contains next error types:
- InvalidArgument
- NotFound
- AlreadyExists

with mapping from PostgreSQL errors, and to HTTP and gRPC.

## TODO
* Add HTTP middleware.
* Add interceptor for wrapping errors into ToGRPCError().
