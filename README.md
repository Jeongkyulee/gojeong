# gojeong
Simple use to private function

# Install

> --------------------------
```Go
go get github.com/Jeongkyulee/gojeong
```
> ----------------------------


### ErrCheck
```Go
if err != nil {
  fmt.Println(err)
}
```
--------------------

### PathRead
```Go
getwd, lastPath := PathRead(dirtoryname string)
reflect.Typeof(getwd) == string
reflect.Typeof(lastPath) == string
```

### ListFile
```Go
[]fileList := ListFile(Path string, fileexted string)
```

