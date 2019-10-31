package main

import(
	"machine"
	"time"
)

var(
	ticks int
)

//
// built-in println uses UART automatically!
//
func main() {
	// configure the Arduino's UART1 interface with the board's TX and RX pins
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})

	go blinkOnboardForever()
	go reportTicksForever()
}

func reportTicksForever() {
	for {
		time.Sleep(time.Millisecond * 2000)
		println("Tick count: ", ticks)
	}
}


// blinkOnboardForever configures and blinks the Nano33's onboard LED
func blinkOnboardForever() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		ticks++

		led.Low()
		time.Sleep(time.Millisecond * 1500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}


