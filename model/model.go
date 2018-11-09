package model

import (
    "github.com/jinzhu/gorm"
    // "strconv"
)

type DemoOrder struct {
    gorm.Model
    Amount         float64
    Order_id       string
    User_name      string
    Status         string
    File_url       string
}

// func (d_o DemoOrder) String() string {
//   return "Demo_Order: {" + "\torder_id: " + d_o.Order_id + "\tuser_name" + d_o.User_name + 
//                         "\tamount: " + strconv.FormatFloat(d_o.Amount, 'E', -1, 64) + "\tstatus: " + d_o.Status +
//                         "\tfile_url: " + d_o.File_url + "}"
// }

func CreateDemoOrder(amount float64, orderid string, username string, status string, fileurl string) *DemoOrder {
    var do *DemoOrder = &DemoOrder{Amount:amount, Order_id:orderid, User_name:username, Status:status, File_url:fileurl}
    return do
}
