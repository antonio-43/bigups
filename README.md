# bigups
#### Resolve domains by using fresh DNS servers

## Usage
Simple as:
```bash
echo "example.com" | bigups -r path/resolvers/list.txt
```
Or:
```bash
cat some_domains.txt | bigups -r path/resolvers/list.txt
```
Or even:
```bash
cat some_domains.txt | bigups -v -r path/resolvers/list.txt
```
for verbose mode.

## Requirements
- Go

## Installation
First clone the repo...
```bash
cd bigups/
go build -o bigups main.go
```

## Be happy and don't do nothing you shound't
