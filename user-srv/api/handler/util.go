package handler

import "encoding/json"

type respBody struct {
    Status  int8        `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func ResponseBody(status int8, msg string, data interface{}) (string, error) {
    body := respBody{
        Status:  status,
        Message: msg,
        Data:    data,
    }

    b, err := json.Marshal(body)
    if err != nil {
        return "", err
    }

    return string(b), nil
}
