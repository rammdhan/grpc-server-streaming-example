# grpc-server-streaming-example
Just a simple app that consists of 3 services (try to implement event-driven but on simple way LOL). 
- Auth service handle the user login
- User service (server) will get the user login info, then send it to notification service (client)
- Notification service (client) will inform the user (only print actually) the info from user service (server)

- This program using:
  --
   - gRPC server streaming that send the user login info to client
   - Redis PubSub to send/receive the user login info
