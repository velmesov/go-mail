package main

import (
	"GoProjects/mail/error"
	"encoding/base64"
	"flag"
	"net/smtp"
)

var (
	to      string
	subject string
	msg     string
	name    string
	login   string
	pass    string
	host    string
	port    string
)

func init() {
	// Устанавливаем флаги
	flag.StringVar(&to, "to", "", "Получатель сообщения")
	flag.StringVar(&subject, "subject", "", "Тема сообщения")
	flag.StringVar(&msg, "msg", "", "Текст сообщения")
	flag.StringVar(&name, "name", "", "Имя отправителя")
	flag.StringVar(&login, "login", "", "Логин почтового аккаунта")
	flag.StringVar(&pass, "pass", "", "Пароль почтового аккаунта")
	flag.StringVar(&host, "host", "", "Хост SMTP сервера")
	flag.StringVar(&port, "port", "", "Порт SMTP сервера")
}

func main() {
	// Парсим флаги
	flag.Parse()

	// Подготовка данных для отправки
	preparationData()
}

// Подготовка данных для отправки
func preparationData() {
	error.CheckFlag(to, "Не указан получатель")
	error.CheckFlag(subject, "Не указана тема сообщения")
	error.CheckFlag(msg, "Пустое сообщение")
	error.CheckFlag(name, "Не указано имя отправителя")
	error.CheckFlag(login, "Не указан логин почтового аккаунта")
	error.CheckFlag(pass, "Не указан пароль почтового аккаунта")
	error.CheckFlag(host, "Не указан хост SMTP сервера")
	error.CheckFlag(port, "Не указан порт SMTP сервера")

	// Получатель
	byteTo, err := base64.StdEncoding.DecodeString(to)
	error.CheckNil(err)
	to := string(byteTo)

	// Тема сообщения
	byteSubject, err := base64.StdEncoding.DecodeString(subject)
	error.CheckNil(err)
	subject := string(byteSubject)

	// Текст сообщения
	byteMsg, err := base64.StdEncoding.DecodeString(msg)
	error.CheckNil(err)
	msg := string(byteMsg)

	// Имя отправителя
	byteName, err := base64.StdEncoding.DecodeString(name)
	error.CheckNil(err)
	name := string(byteName)

	// Логин
	byteLogin, err := base64.StdEncoding.DecodeString(login)
	error.CheckNil(err)
	login := string(byteLogin)

	// Пароль
	bytePass, err := base64.StdEncoding.DecodeString(pass)
	error.CheckNil(err)
	pass := string(bytePass)

	// Хост
	byteHost, err := base64.StdEncoding.DecodeString(host)
	error.CheckNil(err)
	host := string(byteHost)

	// Порт
	bytePort, err := base64.StdEncoding.DecodeString(port)
	error.CheckNil(err)
	port := string(bytePort)

	// Отправка сообщения
	send(to, subject, msg, name, login, pass, host, port)
}

// Отправка сообщения
func send(to, subject, msg, name, login, pass, host, port string) {
	message := "From: " + name + " <" + login + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n" +
		msg

	smtpAuth := smtp.PlainAuth("", login, pass, host)
	err := smtp.SendMail(host+":"+port, smtpAuth, login, []string{to}, []byte(message))
	error.CheckNil(err)
}
