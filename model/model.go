package models

import (
)

type User struct {
    ID       int64  `json:"id"`       // 列名为 `id`
    Username string `json:"username"` // 列名为 `username`
    Password string `json:"password"` // 列名为 `password`
}

