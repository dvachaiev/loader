## Compile loader
```
go build ./cmd/loader
```

## Compile plugin
```
go build -buildmode=plugin ./plugins/test
```

## Run simple HTTP server
```
python3 -m http.server 8000
```

## Run loader
```
./loader http://localhost:8000/test.so
```