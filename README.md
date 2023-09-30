# test-container-runner
This CLI application represents to creation of a docker container by cli.

Firstly you should build Golang application then it will create a standalone executable file. Then you can try following command

````
./test-container-runner start --containerUrl registry.hub.docker.com/library/busybox --exposePorts 8080,8081,8082
````

