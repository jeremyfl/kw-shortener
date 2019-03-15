package main

import (
	fmt "fmt"

	"github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func DecodeProto(c echo.Context) error {
	reference := &ReferenceProto{
		Orig:        "https://www.youtube.com/watch?v=5dS-n8qZjPk",
		Destination: "https://www.kw.com/watch?v=5dS-n8qZjPk",
	}

	data, err := proto.Marshal(reference)

	if err != nil {
		panic("Marshalling failed!")
	}

	// Print Marshall Byte
	fmt.Println(data)

	// Marshalling to JSON String
	m := jsonpb.Marshaler{}
	result, _ := m.MarshalToString(reference)

	return c.JSON(200, result)

}
