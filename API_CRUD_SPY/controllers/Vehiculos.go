package controllers

import (
	"encoding/json"
	"errors"
	"github.com/Zseiru15/API_CRUD_SPY/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// VehiculosController operations for Vehiculos
type VehiculosController struct {
	beego.Controller
}

// URLMapping ...
func (c *VehiculosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Vehiculos
// @Param	body		body 	models.Vehiculos	true		"body for Vehiculos content"
// @Success 201 {int} models.Vehiculos
// @Failure 403 body is empty
// @router / [post]
func (c *VehiculosController) Post() {
	var v models.Vehiculos
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddVehiculos(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"succes": true,
				"status": 201,
				"message": "creacion generada correctamente",
				"data":    v}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Vehiculos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vehiculos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VehiculosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetVehiculosById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{
			"succes": true,
			"status": 200,
			"message": "consulta realizada correctamente",
			"data":    v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Vehiculos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Vehiculos
// @Failure 403
// @router / [get]
func (c *VehiculosController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllVehiculos(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{
			"succes": true,
			"status": 200,
			"message": "consulta realizada correctamente",
			"data":    l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Vehiculos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Vehiculos	true		"body for Vehiculos content"
// @Success 200 {object} models.Vehiculos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VehiculosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Vehiculos{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateVehiculosById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Vehiculos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *VehiculosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteVehiculos(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
