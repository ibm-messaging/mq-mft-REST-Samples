
## Sample application - for MFT REST service
#  MFT V9.0.4  provides support for REST srvice as below .
1. GET Agent Status
2. GET Transfer Status


## Steps to Configure and use the go lang sample to exercise the REST service.

1. Install golanguage.

https://golang.org/doc/install 

2. Set the environment variable "GOPATH" till the application is downloaded. 

 set GOPATH=C:\Users\IBM_ADMIN\goWorkspace

3. Build the application using the following command.

go build mftRESTsampleApplication.go 

The above command will generate an executable by name mftRESTsampleApplication.exe

4. All set to work with the sample application in go.

## Options: 

The above executable will accept hostName, portNumber and  serviceName (agent/transfer) as input arguments.

##Example: 

1.) Execute the command with default value

mftRESTSampleApplication.exe

2.) Execute the command by passing values

mftRESTsampleApplication.exe -host 9.23.34.45 -port 8989 -service transfer

3.) Execute the command by passing just required values

mftRESTsampleApplication.exe -service transfer
