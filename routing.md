# Routes:
- In main.go
```
app:=goLF.New()
app.POST("/route",handler)
app.GET("/route",handler)
app.PATCH("/route",handler)
app.PUT("/route",handler)
app.DELETE("/route',handler)
app.HEAD("/route",handler)
app.OPTIONS("/route",handler)

err:=app.Run()
if err!=nil{
   app.Logger.Errorf("Error in while running service, Err: %v,err)
}
```
- In handler format:

```
func handler(ctx *goLF.Context){
   if an error exists {
     ctx.Error(goLF Error)
    }

    // on success response
    ctx.Response("hii I am done",statusCode (optional))
}
```

- what ctx contains:
```
configs
logger
database
Metrics 
tracing 
request
response

```
- we need to declare some custom errors from our side 
- Add default routes:
   - Metrics route 
   - swagger rendering route
   - health check route
   - route to check database health. 



### Points to remember:

- speed
- error-prone
- Bench Marking testing is very important and note down the results.











