# 1. Start all the 3 servers in background to allow the request

## 1.1 set the port and run them 

  $env:PORT=8081
  go run server/main.go

  $env:PORT=8082
  go run server/main.go

  $env:PORT=8083
  go run server/main.go


# 2. start the load balancer 
   go run main.go


# 3. Test it by using either curl or ab(apache benchmark)

  use for sending multiple request 

  for ($i=0; $i -lt 10; $i++) { curl http://localhost:8080/ }

