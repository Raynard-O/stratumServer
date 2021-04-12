// the server

The server is started using "go run . 8080"      (if no port is added then is is routed to 8080)


// Client

The client is started using " go run ."

// Authorise
The Authorise methos is called using the stdIn of the client.
sending "authorise".
a response of true is returned if the user is authorized by server.


// Subscribe
The Authorise methos is called using the stdIn of the client.
sending "subscribe args".
if arg is less than 1, then the server returns a new job,
else the sever checks for a job with the Extranonce1 number attached in the args i.e ( subscribe uu76d27dxbxt76r2d6gx2tgg)



// Notify (not completed)
The Authorise methos is called using the stdIn of the sever.
sending "notify jobId/Extranonce1".
the server broadcast a jobID to all connected client and waits for a client to accept and returns a true to the allocated client