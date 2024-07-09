package util

import (
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	hwmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"
	"fmt"
)


func SendIoTCommand(client *iotda.IoTDAClient, deviceID string, commandParams map[string]interface{}) (*hwmodel.CreateCommandResponse, error) {
	// 将 commandParams 转换为 *interface{}
	paras := mapToInterface(commandParams)

	request := &hwmodel.CreateCommandRequest{
		DeviceId: deviceID,
		Body: &hwmodel.DeviceCommandRequest{
			Paras:       &paras,
		},
	}

	response, err := client.CreateCommand(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(response)
	return response, nil
}

// mapToInterface 将 map[string]interface{} 转换为 interface{}
func mapToInterface(m map[string]interface{}) interface{} {
	return m
}