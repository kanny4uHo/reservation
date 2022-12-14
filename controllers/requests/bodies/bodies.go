package bodies

import (
	"barckend/utils"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/pkg/errors"
	"time"
)

type BaseBody struct{}

func Require(b any, body []byte) error {
	err := utils.ParseAndValid(b, body)
	if err != nil {
		return errors.Wrap(err, "Input parsing error")
	}
	return nil
}

type LoginForm struct {
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Password string  `json:"password" valid:"Required; MinSize(6); MaxSize(40)"`
}

func (form *LoginForm) Valid(validation *validation.Validation) {
	if form.Email == nil && form.Phone == nil {
		validation.AddError("login_check", "Phone or email required")
	}
	if form.Phone != nil {
		validation.Numeric(form.Phone, "phone")
		validation.MinSize(form.Phone, 11, "phone")
		validation.MaxSize(form.Phone, 11, "phone")
	} else {
		validation.MaxSize(form.Email, 30, "email")
		validation.Email(form.Email, "email")
	}
}

type RegisterAdmin struct {
	Email      string `json:"email" valid:"Required; Email; MaxSize(30)"`
	Phone      string `json:"phone" valid:"Required; MinSize(11); MaxSize(11); Numeric"`
	Password   string `json:"password" valid:"Required; MinSize(6); MaxSize(40)"`
	Name       string `json:"name" valid:"Required; MinSize(3); MaxSize(50)"`
	Surname    string `json:"surname" valid:"Required; MinSize(3); MaxSize(50)"`
	Patronymic string `json:"patronymic" valid:"Required; MinSize(3); MaxSize(50)"`
}

type RegisterOwner struct {
	Email      string `json:"email" valid:"Required; Email; MaxSize(30)"`
	Phone      string `json:"phone" valid:"Required; MinSize(11); MaxSize(11); Numeric"`
	Password   string `json:"password" valid:"Required; MaxSize(100)"`
	Name       string `json:"name" valid:"Required; MinSize(3); MaxSize(50)"`
	Surname    string `json:"surname" valid:"Required; MinSize(3); MaxSize(50)"`
	Patronymic string `json:"patronymic" valid:"Required; MinSize(3); MaxSize(50)"`
}

type CreateBarInfo struct {
	Email       string      `json:"email" valid:"Required; Email; MaxSize(30)"`
	Phone       string      `json:"phone" valid:"Required; MinSize(11); MaxSize(11); Numeric"`
	Name        string      `json:"name" valid:"Required; MinSize(3); MaxSize(50)"`
	Description string      `json:"description" valid:"Required; MaxSize(400)"`
	Address     string      `json:"address" valid:"Required; MaxSize(100)"`
	WorkHours   []WorkHours `json:"work_hours"  valid:"Required; MaxSize(7)"`
}

func (form *CreateBarInfo) Valid(validation *validation.Validation) {
	Validate(validation, form.WorkHours)
}

func Validate(validation *validation.Validation, workHours []WorkHours) {
	for _, wh := range workHours {
		isValid, err := validation.Valid(wh)
		if err != nil {
			validation.AddError("Internal error", "Internal error occurred while validating")
		} else if !isValid {
			validation.AddError("WorkHours", "Work hours are invalid")
		}
	}
}

type WorkHours struct {
	Weekday int8   `json:"weekday" valid:"Required; Range(1,7)"`
	From    string `json:"from" valid:"Required; Match(/^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$/)"`
	To      string `json:"to" valid:"Required; Match(/^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$/)"`
}

type UpdateProfile struct {
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Name       *string `json:"name,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
}

func (form *UpdateProfile) Valid(validation *validation.Validation) {
	if form.Phone != nil {
		validation.Numeric(form.Phone, "phone")
		validation.MinSize(form.Phone, 11, "phone")
		validation.MaxSize(form.Phone, 11, "phone")
	}
	if form.Email != nil {
		validation.Email(form.Email, "email")
		validation.MaxSize(form.Email, 30, "email")
	}
	if form.Surname != nil {
		validation.MinSize(form.Surname, 3, "surname")
		validation.MaxSize(form.Surname, 50, "surname")
	}
	if form.Name != nil {
		validation.MinSize(form.Name, 3, "name")
		validation.MaxSize(form.Name, 50, "name")
	}
	if form.Patronymic != nil {
		validation.MinSize(form.Patronymic, 3, "patronymic")
		validation.MaxSize(form.Patronymic, 50, "patronymic")
	}
}

type UpdateBar struct {
	Email       *string      `json:"email,omitempty"`
	Phone       *string      `json:"phone,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Address     *string      `json:"address,omitempty"`
	WorkHours   *[]WorkHours `json:"work_hours,omitempty"`
	IsVisible   *bool        `json:"is_visible_to_user,omitempty"`
}

func (form *UpdateBar) Valid(validation *validation.Validation) {
	if form.Phone != nil {
		validation.Numeric(form.Phone, "phone")
		validation.MinSize(form.Phone, 11, "phone")
		validation.MaxSize(form.Phone, 11, "phone")
	}
	if form.Email != nil {
		validation.Email(form.Email, "email")
		validation.MaxSize(form.Email, 30, "email")
	}
	if form.Name != nil {
		validation.MinSize(form.Name, 3, "name")
		validation.MaxSize(form.Name, 50, "name")
	}
	if form.Address != nil {
		validation.MaxSize(form.Name, 100, "address")
	}
	if form.Description != nil {
		validation.MaxSize(form.Name, 400, "description")
	}
	if form.WorkHours != nil {
		validation.MaxSize(form.WorkHours, 7, "work_hours")
		Validate(validation, *form.WorkHours)
	}
}

type CreateTable struct {
	BarId       int64   `json:"bar_id" valid:"Required; Min(1)"`
	Name        string  `json:"name" valid:"Required; MinSize(1); MaxSize(30)"`
	Capacity    int8    `json:"persons" valid:"Required; Range(1,70)"`
	Description *string `json:"description,omitempty"`
}

func (form *CreateTable) Valid(validation *validation.Validation) {
	if form.Description != nil {
		validation.MaxSize(form.Name, 400, "description")
	}
}

type UpdateTable struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
}

func (form *UpdateTable) Valid(validation *validation.Validation) {
	if form.Name != nil {
		validation.MinSize(form.Name, 1, "name")
		validation.MaxSize(form.Name, 30, "name")
	}
	if form.Capacity != nil {
		validation.Range(form.Capacity, 1, 70, "capacity")
	}
	if form.Description != nil {
		validation.MaxSize(form.Name, 400, "description")
	}
}

type CreateReservation struct {
	TableId     int64  `json:"table_id" valid:"Required; Min(1)"`
	FromTime    string `json:"from_time" valid:"Required; Match(/^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$/)"`
	ToTime      string `json:"to_time" valid:"Required; Match(/^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$/)"`
	Date        string `json:"date" valid:"Required"`
	PersonCount int    `json:"persons" valid:"Required; Range(1,70)"`
}

func (form *CreateReservation) Valid(validation *validation.Validation) {
	formattedFromTime := fmt.Sprintf("%s %s", form.Date, form.FromTime)
	_, err := time.Parse("02.01.2006 15:04", formattedFromTime)
	if err != nil {
		validation.AddError("datetime", "Date or timing has wrong format or value")
	}
	formattedToTime := fmt.Sprintf("%s %s", form.Date, form.ToTime)
	_, err = time.Parse("02.01.2006 15:04", formattedToTime)
	if err != nil {
		validation.AddError("datetime", "Date or timing has wrong format or value")
	}
}
