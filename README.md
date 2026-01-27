A front-end web-application that connects to 5 microservices:
 - Broker - optional single point of entry to micro-services;
 - Authentication - PG - based service to log users in;
 - Logger - loggs event of all the microservices, based upon MongoDB;
 - Mail - sends email using a template;
 - Listener - consumes messages in RMQ and starts a process.

Communication:

 - REST API with JSON as transport;
 - Sending and receiving using RPC and gRPC;
 - Event handling using AMQP.