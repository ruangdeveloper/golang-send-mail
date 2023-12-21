package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDialer(t *testing.T) {
	dialer := InitDialer()

	assert.NotEmpty(t, dialer)
}

func TestSendMail(t *testing.T) {
	err := SendMail(
		"user@example.com",
		"Halo, ini subject",
		"Halo, ini isi pesan yang dikirim menggunakan GoMail ya!",
	)

	assert.NoError(t, err)
}
