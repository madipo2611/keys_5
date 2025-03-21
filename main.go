package main

import (
	"html/template"
	"os"
)

type Page struct {
	OrgName  string
	OrgInfo  string
	OrgStile []string
}

func main() {
	orgName := "ООО \"РЕГ.РУ\""
	orgInfo := "ООО \"РЕГ.РУ\" — один из ведущих регистраторов доменных имен в России. Наша компания предоставляет широкий спектр услуг в области регистрации доменов, хостинга и интернет-решений для бизнеса. Мы помогаем клиентам создать и развивать их онлайн-присутствие, предлагая надежные и доступные услуги."
	orgStile := []string{"Impact", "Times New Roman", "Georgia", "Courier New", "Verdana"}

	page := Page{
		OrgName:  orgName,
		OrgInfo:  orgInfo,
		OrgStile: orgStile,
	}

	// Парсинг HTML-шаблона
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	// Генерация html файла
	fileName := "org.html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, page)
	if err != nil {
		panic(err)
	}

	println("HTML-страница успешно создана:", fileName)
}
