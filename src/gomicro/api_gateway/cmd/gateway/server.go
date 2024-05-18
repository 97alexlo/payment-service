package main

import "net/http"

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/customer/payment/authorize", customerPaymentAuthorize)
	http.HandleFunc("/customer/payment/capture", customerPaymentCapture)
	http.HandleFunc("/ledger", customerLedger)
}

func login(http.ResponseWriter, *http.Request) {

}

func customerPaymentAuthorize(http.ResponseWriter, *http.Request) {

}

func customerPaymentCapture(http.ResponseWriter, *http.Request) {

}

func customerLedger(http.ResponseWriter, *http.Request) {

}