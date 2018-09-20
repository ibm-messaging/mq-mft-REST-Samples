# mq-mft-REST-golang-sample

## Compaling Sample ApplicationSteps to Configure and use the golang sample 

```go build $GOPATH/src/github.com/ibm-messaging/mq-mft-REST-Samples/mftRESTsampleApplication.go```

The above command will generate an executable by name mftRESTsampleApplication.exe.

##About sample application.

This is a sample application to demonstrate the REST API services provided to check the agent and transfer status.

##Agent REST Service

The responsce will havea the details of agents in a  JSON array.

The default attributes are  -

```name```: Name of the agent. 

```state```: Type of the agent. Value will be any of the following.

•	standard   - non-Bridge agent

•	protocolBridge – Protocale bridge type agent.

•	connectDirectBridge – Connect Direct Bridge type agent 

```type```:

•	active - Agent is running and is sending or receiving files. The agent is publishing its status at regular intervals. The last update was received within the expected time

•	ready -Agent is running, but is not sending or receiving files. The agent is publishing its status at regular intervals. The last update was received within the expected time period.

•	starting - Agent is starting, but is not yet ready to perform transfers.

•	unreachable - Agent status updates were not received at the expected time intervals. The agent might have stopped running due to an error, or have been shut down abruptly, or be running but experiencing communication problems

•	stopped -Agent has been stopped. It was shut down in a controlled manner.

•	endedUnexpectedly - Agent has ended unexpectedly. The agent will be automatically restarted, unless there have been more than maxRestartCount restart within the maxRestaratInterval time and the maxRestartDelay value is less than or equal to 0.

•	noInformation - If agent version is 7.0.2 or earlier. Then agent will not publish updates in a form that this command can process.

•	unknown - The status of the agent cannot be determined. It might have published a status which is not recognized by this tool. 

•	problem - Agent command handler might not be working. The agent is publishing status messages, but these status messages are out of date.


##Transfer REST Service

```destinationAgent``` : Name of the destination agent.

```sourceAgent``` : Name of the source agent.

```id```:  Transfer ID.

```status```: Status of the transfer.

```statistics``` : Statistical details about the transfer which includes the following -

     • endTime
     
     •`numberOfFileFailures
     
     • numberOfFileFailures
     
     • numberOfFileSuccesses
     
     • numberOfFileWarnings
     
     • numberOfFiles
     
     • startTime

```Originator```: Details about the transfer source, which includes the following information.
     
     • host
     • userId

## Options: 

The above executable will accept hostName, portNumber and  serviceName (agent/transfer) as input arguments.

1.) An attempt to execute with out any optional value will take the 'localhost' and the port '9080' and execute the command.

```mftRESTSampleApplication.exe```

> **NOTICE**: It is expected to work if the REST service is hosted in same server where exe is avaliable.  

(by default, it attempts to execute agent status REST service)

2.) An attempt to execute the REST service  which is hosted on a different server.

#Transfer Service

```mftRESTsampleApplication.exe -host 9.23.34.45 -port 8989 -service transfer```

#Agent Service

```mftRESTsampleApplication.exe -host 9.23.34.45 -port 8989 -service agent```

## Sample output: 

```mftRESTsampleApplication.exe -host 9.23.34.45 -port 8989 -service transfer```

==============================================================
```
{"transfer": [

  {
  
    "destinationAgent": {"name": "DA"},
    
    "id": "414D51204D465444454D4F20202020203FD7A05B21388809",
    
    "originator": {
    
      "host": "9.199.194.125",
      
      "userId": "Gandhimathy"
      
    },
    
    "sourceAgent": {"name": "SA"},
    
    "statistics": {
    
      "endTime": "2018-09-20T09:01:04.014Z",
      
      "numberOfFileFailures": 2,
      
      "numberOfFileSuccesses": 0,
      
      "numberOfFileWarnings": 0,
      
      "numberOfFiles": 2,
      
      "startTime": "2018-09-20T09:01:03.640Z"
      
    },
    
    "status": {"state": "failed"}
    
  },
  
  {
  
    "destinationAgent": {"name": "DA"},
    
    "id": "414D51204D465444454D4F20202020203FD7A05B21386914",
   
    "originator": {
    
     "host": "9.199.194.125",
     
      "userId": "Gandhimathy"
    
     },
     
     "sourceAgent": {"name": "SA"},
     
     "statistics": {
     
     "endTime": "2018-09-20T09:00:56.495Z",
     
     "numberOfFileFailures": 2,
     
     "numberOfFileSuccesses": 0,
     
     "numberOfFileWarnings": 0,
     
     "numberOfFiles": 2,
     
     "startTime": "2018-09-20T09:00:56.183Z"
    
    },
    
    "status": {"state": "failed"}
  
  }
 
 ]

} 
```

==============================================================
  
  ```mftRESTsampleApplication.exe -host 9.23.34.45 -port 8989 -service transfer```

==============================================================
```
{"agent": [

  {
  
    "name": "DA",
    
    "state": "ready",
    
    "type": "standard"
    
  },
  
  {
  
    "name": "SA",
    
    "state": "ready",
    
    "type": "standard"
    
  },
  
  {
  
    "name": "CDAGENT",
    
    "state": "stopped",
    
    "type": "connectDirectBridge"
    
  },
  
  {
  
    "name": "PROTOCOLBRIDGE",
    
    "state": "stopped",
    
    "type": "protocolBridge"
    
  }

]

}
```
==============================================================
