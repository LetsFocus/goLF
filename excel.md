# Excel:

```
g:=goLF.New() // routing/ metrics...etc

excel:=goLF.NewExcel(addresses) // eg:Excel_Link=https://docs.google.com/spreadsheets/d/1QWI4LCz4kaw76z1BF1A6X6WFKIR8Lk5uaOurDI7w2mk/edit?usp=sharing  
excel.Post()
excel.Get()
excel.Delete()
excel.Put()
excel.Patch()
excel.GetColumns()

g.POST("/data",handler)

g.Run()
```

-- 3 Layer
- handler
- service
- store  ---- excel
