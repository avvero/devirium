

Latency + Processing Time = Response Time

- Latency = the time the message is in transit between two points (e.g. on the network, passing through gateways, etc.) AND time to wait for processing.
- Processing time = the time it takes for the message to be processed (e.g. translation between formats, enriched, or whatever)

«Latency and response time are often used synonymously, but they are not the same. The response time is what the client sees: besides the actual time to process the request (the service time), it includes network delays and queueing delays. Latency is the duration that a request is waiting to be handled—during which it is latent, awaiting service» [[Designing Data Intensive Applications]]

#latency 
#draft