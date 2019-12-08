//go:generate protoc -I ./echo --go_out=plugins=grpc,paths=source_relative:./echo ./echo/echo.proto

package proto
