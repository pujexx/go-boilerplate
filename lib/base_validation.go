package lib

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
	"reflect"
)



func ValidateStruct(data interface{}) (bool, []ValidateError) {
	validate := validator.New()
	errs := []ValidateError{}
	if err := validate.Struct(data); err != nil {
		validatorErros := err.(validator.ValidationErrors)
		logrus.Println("validation", validatorErros)
		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, trans)
		for _, e := range validatorErros {
			errs = append(errs, ValidateError{
				Field: getJSONName(data,e.Field()),
				Error: e.Translate(trans),
			})
		}
		return false, errs
	}
	return true, errs
}

func ValidateVar(data interface{},name string, role string) (bool, []ValidateError) {
	validate := validator.New()
	errs := []ValidateError{}
	if err := validate.Var(data,role); err != nil {
		validatorErros := err.(validator.ValidationErrors)
		logrus.Println("validation", validatorErros)
		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, trans)
		for _, e := range validatorErros {
			errs = append(errs, ValidateError{
				Field: name,
				Error: fmt.Sprintf("%v%v",name,e.Translate(trans)),
			})
		}
		return false, errs
	}
	return true, errs
}


func getJSONName(src interface{}, f string) string{
	st := reflect.TypeOf(src)
	field,_ := st.FieldByName(f)
	name := field.Tag.Get("json")
	if name != "" {
		return name
	}else {
		return f
	}
}