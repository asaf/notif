package main

import (
	"flag"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"os"
	"strings"
)

func Smtp(args []string) error {
	cli := flag.NewFlagSet("smtp", flag.ExitOnError)

	// SMTP sub-command flag pointers
	fromPtr := cli.String("from", "", "From. (Required)")
	toPtr := cli.String("to", "", "Comma separated email recipients. (Required)")
	subjectPtr := cli.String("subject", "", "Subject. (Required)")
	smtpPtr := cli.String("server", "localhost", "SMTP server address.")
	portPtr := cli.Int("port", 25, "SMTP server port.")
	userPtr := cli.String("user", "", "Username for auth.")
	passPtr := cli.String("password", "", "Password for auth.")

	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	if err := cli.Parse(args); err != nil {
		return err
	}

	// Required Flags
	if *fromPtr == "" {
		cli.PrintDefaults()
		os.Exit(1)
	}

	if *toPtr == "" {
		cli.PrintDefaults()
		os.Exit(1)
	}

	if *subjectPtr == "" {
		cli.PrintDefaults()
		os.Exit(1)
	}

	to := strings.Split(*toPtr, ",")

	// read os.Stdin into body
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", *fromPtr)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", *subjectPtr)
	m.SetBody("text/html", string(body))
	d := gomail.NewPlainDialer(*smtpPtr, *portPtr, *userPtr, *passPtr)
	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
