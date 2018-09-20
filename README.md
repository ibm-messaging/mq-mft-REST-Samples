## MFT REST API Service

IBM MQ MFT V9.0.5 support REST API service to query the details of Agent and Transfer.  
This repository includs a sample application to demonstrates MFT REST service API written in the Go language.

## Pre-req 

1.) Please ensure that you MFT is installed.

```https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_9.0.0/com.ibm.mq.helphome.v90.doc/WelcomePagev9r0.html```

2.) Configure MFT REST service.

```https://developer.ibm.com/messaging/2018/03/19/mq-9-0-5-quick-start-mft-rest-apis/```

3.) Follow the following to get more insight on AgentStatus and TransferStatus REST services.

```https://developer.ibm.com/messaging/2018/03/21/mqv9-0-5-mft-rest-api-agent-status-details/```

```https://developer.ibm.com/messaging/2018/03/29/mq-v9-0-5-mft-rest-api-transfer-status/```

## Getting started

If you are unfamiliar with Go, the following steps can help create a working environment with source code in a suitable tree. Initial setup tends to be platform-specific, but subsequent steps are independent of the platform.

### Linux

* Install the Go runtime and compiler. On Linux, the packaging may vary but a typical directory for the code is `/usr/lib/golang`.

* Create a working directory. For example, ```mkdir $HOME/gowork```

* Set environment variables. Based on the previous lines,

  ```export GOROOT=/usr/lib/golang```

  ```export GOPATH=$HOME/gowork```

* If using a version of Go from after 2017, you must set environment variables to permit some compile/link flags. This is due to a security fix in the compiler.

  ```export CGO_LDFLAGS_ALLOW="-Wl,-rpath.*"```

* Install the git client

### Windows

* Install the Go runtime and compiler. On Windows, the common directory is `c:\Go`
* Ensure you have a gcc-based compiler, for example from the Cygwin distribution. I use the mingw variation, to ensure compiled code can be used on systems without Cygwin installed
* Create a working directory. For example, ```mkdir c:\Gowork```
* Set environment variables. Based on the previous lines,

  ```set GOROOT=c:\Go```

  ```set GOPATH=c:\Gowork```

  ```set CC=x86_64-w64-mingw32-gcc.exe```

* The `CGO_LDFLAGS_ALLOW` variable is not needed on Windows
* Install the git client

### Common

* Make sure your PATH includes routes to the Go compiler (`$GOROOT/bin`), the Git client, and the C compiler.
* Change directory to the workspace you created earlier. (`cd $GOPATH`)
* Use git to get a copy of the MQ components into a new directory in the workspace.

  ```git clone https://github.com/ibm-messaging/mq-mft-REST-Samples.git src/github.com/ibm-messaging/mq-mft-REST-Samples```

* Follow the instructions in the [mq-metric-samples repository](https://github.com/ibm-messaging/mq-metric-samples) to compile the sample programs you are interested in.

At this point, you should have a compiled copy of the code in `$GOPATH/bin`.

## Issues and Contributions

For feedback and issues relating specifically to this package, please use the [GitHub issue tracker](https://github.com/ibm-messaging/mq-mft-REST-Samples/issues).

Contributions to this package can be accepted under the terms of the IBM Contributor License Agreement,
found in the [CLA file](CLA.md) of this repository. When submitting a pull request, you must include a statement stating
you accept the terms in the CLA.

## Copyright

© Copyright IBM Corporation 2018, 2019
