package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPort(inputPort string) (string, error) {
    slicedInput:=strings.Split(inputPort, "=")
    if len(slicedInput) != 2 {
        return "", fmt.Errorf("Unexpected parameter, expected --PORT=#### but received: %s", inputPort)
    }

    portKey:=slicedInput[0]

    if !strings.EqualFold(portKey, "--port"){
        return "", fmt.Errorf("Unexpected parameter, expected --PORT=#### but received: %s", inputPort)
    }

    portValue:=slicedInput[1]

    if portValue==""{
    return "", fmt.Errorf("Empty PORT")
    }

    val, err := strconv.Atoi(portValue)
    if err != nil {
        return "", fmt.Errorf("The PORT must be a number ranging from 1,024 to 49,151, got: %s", portValue)
    }
    if val < 1024 {
        return "", fmt.Errorf("Can't bind lower ports -- Received: %s", portValue)
    }
    if val > 49151 {
        return "", fmt.Errorf("Can't bind private ports -- Received: %s", portValue)
    }
    return portValue, nil
}
