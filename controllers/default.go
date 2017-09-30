package controllers

import (
	"github.com/astaxie/beego"
)

//MainController main controller
type MainController struct {
	beego.Controller
}

//Get index function
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Blance check user blance in game
func (mc *MainController) Blance() {
}

//Settle game coin settle
func (mc *MainController) Settle() {
}

//Register register
func (mc *MainController) Register() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}

//Recharge recharge to game function
func (mc *MainController) Recharge() {
}
