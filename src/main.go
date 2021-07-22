// Package main is the application entry point
package main

import (
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	/* This would match the route like the following:
		/sum/3/5
		/product/6/23
	*/
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	// Obtain the values of the route params defined in the route above
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))
	
	//Set the values for use in the template
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.gohtml"
	
	// Perform the calculation depending on the "operation" route param
	switch operation {
	case "sum":
		c.Data["result"] = add(num1, num2)
	case "product":
		c.Data["result"] = multiply(num1, num2)
	default:
		c.TplName = "invalid-route.gohtml"
	}
}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
