all: access access_c

access_c: access.c
	gcc access.c -o access_c

access: access.go
	go build access.go
