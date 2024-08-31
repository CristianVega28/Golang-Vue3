package database

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

func Helper_template_insert[T any](struct_type T) string {

	reflect_data := reflect.TypeOf(struct_type)

	template_struct := `(`

	for i := 0; i < reflect_data.NumField(); i++ {
		field_struct := reflect_data.Field(i)

		var text_result string
		if i != reflect_data.NumField()-1 {
			text_result = fmt.Sprintf("{{.%s}}, ", field_struct.Name)

			fmt.Println(text_result)

			if field_struct.Type.String() == "string" {
				text_result = fmt.Sprintf("\"{{.%s}}\", ", field_struct.Name)
			}

		} else {
			if field_struct.Type.String() == "string" {
				text_result = fmt.Sprintf("\"{{.%s}}\") ", field_struct.Name)
			} else {
				text_result = fmt.Sprintf("{{.%s}}) ", field_struct.Name)

			}
		}

		template_struct += text_result
	}

	return template_struct
}

func insert_data_in_template[T any](template_param string, struct_one T) string {

	var template_resultado bytes.Buffer
	t_tmpl := template.Must(template.New("message").Parse(template_param))
	if err := t_tmpl.Execute(&template_resultado, struct_one); err != nil {
		panic("Hubo un error y no se pudo crear el template")
	}
	message := template_resultado.String()

	return message
}

func Helper_insert_data_template[T any](struct_array []T) string {

	result := ""

	if len(struct_array) > 0 {
		template_create := Helper_template_insert(struct_array[0])

		for index, struct_one := range struct_array {
			if index != len(struct_array)-1 {
				result_insert_data_template := insert_data_in_template(template_create, struct_one)

				result += fmt.Sprintf("%s, ", result_insert_data_template)
			} else {
				result_insert_data_template := insert_data_in_template(template_create, struct_one)

				result += fmt.Sprintf("%s; ", result_insert_data_template)

			}
		}
	}

	return result
}
