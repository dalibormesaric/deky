# DEKY

Idea is to have two programs, one running on raspberry pi zero and other on the PC. They would communicate over usb network.

## Build

``` sh
env GOOS=linux GOARCH=arm GOARM=6 go build -o deky
scp deky pi@raspberrypizero.local:~/deky
```

### Use

``` sh
curl raspberrypizero.local:8080/Hello
```

## Hardware components

[Raspberry Pi Zero](https://www.raspberrypi.com/products/raspberry-pi-zero/)

[Adafruit 2.2" PiTFT HAT - 320x240 Display](https://learn.adafruit.com/adafruit-2-2-pitft-hat-320-240-primary-display-for-raspberry-pi)
