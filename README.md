# Config poc using go
This is an example of how to implement a configuration in go. First it tries to use the env variables, then tries to get the values from the *.toml file. If there are still some missing field, it uses a default value.

Run:
```
go run main.go
```