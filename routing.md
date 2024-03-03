# Routes:
- In main.go
```
app:=goLF.New()
app.POST("/routes",handler)
app.GET("/routes",handler)
app.PATCH("/routes",handler)
app.PUT("/route",handler)
app.DELETE("/route',handler)

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
    ctx.Response("hii I am done")
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
- we need to declare some custom error from our side 
- Add default routes:
   - Metrics route 
   - swagger rendering route
   - health check route
   - route to check database health. 















