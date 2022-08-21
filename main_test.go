package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	var now = time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())
}

// curl --location --request GET 'https://api.myquran.com/v1/sholat/jadwal/1203/2022/08/21'
// curl --location --request GET 'https://api.myquran.com/v1/sholat/jadwal/1609/2021/06/23'
