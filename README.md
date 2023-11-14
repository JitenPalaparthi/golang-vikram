# golang

- to create go application, create a directory use. The name of the module is virtual root path of the project.
```
go mod init <name>
```

# build 

```
go build -o app main.go
```
# stripe down the build

```
go build -ldflags "-w" -o app_striped main.go
```

## keywords

- there are 25 keywords in Golang

- break,case,const,continue,default,else,fallthrough,for,func,goto,if,import,map,package,range,return,switch,var 

## builtin funcs
- append, cap, clear, copy, delete, len, print, println
