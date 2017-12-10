package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CabinetOrderRecord struct {
	Id              int    `orm:"column(order_no);pk" description:"内部生成的订单号"`
	CustomerId      string `orm:"column(customer_id);size(255);null" description:"顾客id 微信 openid 支付宝？"`
	PayType         int8   `orm:"column(pay_type)" description:"1 微信 2支付宝 3？"`
	ThirdOrderNo    string `orm:"column(third_order_no);size(255);null" description:"第三方支付id"`
	CabinetDetailId int    `orm:"column(cabinet_detail_id)"`
	Fee             int    `orm:"column(fee)" description:"钱数 单位分"`
	CreateDate      int    `orm:"column(create_date)"`
	PayDate         int    `orm:"column(pay_date);null"`
	IsPay           int8   `orm:"column(is_pay)" description:"是否支付 0 未支付 1已经支付"`
}

func (t *CabinetOrderRecord) TableName() string {
	return "cabinet_order_record"
}

func init() {
	orm.RegisterModel(new(CabinetOrderRecord))
}

// AddCabinetOrderRecord insert a new CabinetOrderRecord into database and returns
// last inserted Id on success.
func AddCabinetOrderRecord(m *CabinetOrderRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCabinetOrderRecordById retrieves CabinetOrderRecord by Id. Returns error if
// Id doesn't exist
func GetCabinetOrderRecordById(id int) (v *CabinetOrderRecord, err error) {
	o := orm.NewOrm()
	v = &CabinetOrderRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCabinetOrderRecord retrieves all CabinetOrderRecord matches certain condition. Returns empty list if
// no records exist
func GetAllCabinetOrderRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CabinetOrderRecord))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CabinetOrderRecord
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCabinetOrderRecord updates CabinetOrderRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateCabinetOrderRecordById(m *CabinetOrderRecord) (err error) {
	o := orm.NewOrm()
	v := CabinetOrderRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCabinetOrderRecord deletes CabinetOrderRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCabinetOrderRecord(id int) (err error) {
	o := orm.NewOrm()
	v := CabinetOrderRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CabinetOrderRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}