package main

import (
	"fmt"
	"time"
)

var itemsChannel = make(chan string)
var cleanedItemsChannel = make(chan string)

func Channels() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "harta", "kerang"}

	go menyelam(items)
	go membersihkan()
	go menyimpan()

	time.Sleep(500 * time.Millisecond)
}

func menyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil medapatkan " + item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil membersihkan " + rawItem)
		cleanedItemsChannel <- "hartaBersih"
	}
}

func menyimpan() {
	for cleanedItem := range cleanedItemsChannel {
		fmt.Println("Berhasil menyimpan " + cleanedItem)
	}
}
