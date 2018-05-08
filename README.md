# awx-client-go

[![Build Status](https://travis-ci.org/moolitayer/awx-client-go.svg?branch=master)](https://travis-ci.org/moolitayer/awx-client-go)

A golang client library for [AWX](https://github.com/ansible/awx) and [Ansible Tower](https://www.ansible.com/products/tower) REST API.

## Installation
Install awx-client-go using the "go-get" command:
```
go get github.com/golang/glog # Dependency
go get github.com/moolitayer/awx-client-go/awx
```

## Usage
### import
```go
import 	"github.com/moolitayer/awx-client-go/awx"
```

### Creating a connection:
```go
// Uses the builder pattern:
connection, err := awx.NewConnectionBuilder().
  Url("http://awx.example.com/api").          // Url is mandatory
  Username(username).
  Password(password).
  Token("TOKEN").
  Bearer("BEARER").
  CAFile("/etc/pki/tls/cert.pem").
  Insecure(insecure).
  Proxy(http://myproxy.example.com).
  Build()                                    // Create the client
if err != nil {
  panic(err)
}
defer connection.Close()                      // Don't forget to close the connection!
```

`Url()` points at an AWX server's root API endpoint (including the '/api' path) and is mandatory.  
`Proxy()` specifies a proxy server to use for all outgoing connection to the AWX server.
#### Authentication
Use one of:
- `Username()` and `Password()` specify Basic Auth for AWX API server.
- `Token()` uses the authtoken/ endpoint and works with AWX < 1.0.5 and Ansible tower < 3.3.
- `Bearer()` uses OAuth2 and works since AWX 1.0.5 and Ansible Tower 3.3.

When Username and Password are specified the client will attempt to acquire Token Or Bearer based on what the server supports.

#### TLS
`CAFile()` specifies path of a file containing PEM encoded CA certificates used to verify the AWX server. If no CAFile is provided, the default host trust store will be used. `CAFile()` can be used multiple times to specify a list of files.  
`Insecure(true)` can be specified to disable TLS verification.

### Supported resources
- Projects
- Jobs
- Job Templates

Please submit feature requests as Github [issues](https://github.com/moolitayer/awx-client-go/issues/new).

### Retrieving resources
```go
projectsResource := connection.Projects()

// Get a list of all Projects.
getProjectsRequest := projectsResource.Get()
getProjectsResponse, err := getProjectsRequest.Send()
if err != nil {
  panic(err)
}

// Print the results:
projects := getProjectsResponse.Results()
for _, project := range projects {
  fmt.Printf("%d: %s - %s\n", project.Id(), project.Name(), project.SCMURL())
}
```
#### Filtering
User `Filter()` on a request to filter lists:
```go
projectsResource := connection.Projects()

// Get a list of all Projects using git SCM.
getProjectsResponse, err := projectsResource.Get().
  Filter("scm_type", "git").
  Send()
getProjectsResponse, err := getProjectsRequest.Filter("scm_type", "git").Send()

```
#### Retrieving resource by id
Use `Id(...)` on a resource list to get a single resource
```go
// Get a resource managing a project with id=4
projectResource := connection.Projects().Id(4)

// Send the request to retrieve the project:
getProjectResponse, err := projectResource.Get().Send()
```

#### Launching a Job from a Template
```go
// Launch Job Template with id=8
launchResource := connection.JobTemplates().Id(8).Launch()

response, err := launchResource.Post().
  ExtraVars(map[string]string{"awx_environment": "staging"}).
  ExtraVar("instance", "example.com").
  Limit("master.example.com").
  Send()
if err != nil {
  return err
}
```
`ExtraVars()` Specifies a map passed to AWX as extra vars.  
`ExtraVar()` Specifies a single key value pair.  
`Limit()` is an Ansible host pattern.
See [Job Template](http://docs.ansible.com/ansible-tower/latest/html/userguide/job_templates.html)

## Examples

See [examples](examples).
