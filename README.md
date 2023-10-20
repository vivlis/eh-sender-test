# eh-sender-test
Event Hub sender test

Send events to specified event hub 

Event hub connection string is teaken from "EHProducer" environmental variable 
like "Endpoint=sb://****.servicebus.windows.net/;SharedAccessKeyName=SendAndListenPolicy;SharedAccessKey=****;EntityPath=***",

Other env variables:

"BatchSize" - Batch size is number of events in batch. Default 1.
"NoEvents" - Number of produces events. Default 1. 
"MaxNoPart" - Number of partition of consumer EH, Default is 0. If specified partition in batch will be set to random number between 0 - specified nummber  