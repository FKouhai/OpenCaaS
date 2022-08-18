# OpenCaaS (WIP)

OpenCaaS is a backend platform, that will be used to automatically deploy containers on a set of nodes(could be VMs or bare metal as long as they are running a linux os)
- will use the Docker SDK for go to do the different tasks such as creating containers, scaling a container, removing a container
- need to have an etcd database instantiated where the information of the nodes and containers will be stored
- will consists on 2 binaries one for the server and one for the client both will be a webserver to allow an easy of communication and an ease of automation for the deployment and container OPS
- The server will connect to the etcd db to read and write information in it, such as container name, number of replicas, node where its deployed, etc...

## Architecture
This is the application architecture on a high level standpoint
 
![](OpenCaaS.png)