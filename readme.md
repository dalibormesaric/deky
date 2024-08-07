# DEKY

Idea is to have two programs, one running on raspberry pi zero and other on the PC. They would communicate over usb network.

## Setup

To see which Raspbian version is on Raspberry Pi Zero, run:
``` sh
cat /etc/os-release
```
For me it was `bullseye`

Then download compiler from:
> https://sourceforge.net/projects/raspberry-pi-cross-compilers/

I downloaded https://sourceforge.net/projects/raspberry-pi-cross-compilers/files/Raspberry%20Pi%20GCC%20Cross-Compiler%20Toolchains/Bullseye/GCC%2010.3.0/Raspberry%20Pi%201%2C%20Zero/

Using
``` sh
wget https://downloads.sourceforge.net/project/raspberry-pi-cross-compilers/Raspberry%20Pi%20GCC%20Cross-Compiler%20Toolchains/Bullseye/GCC%2010.3.0/Raspberry%20Pi%201%2C%20Zero/cross-gcc-10.3.0-pi_0-1.tar.gz
```

And extracted using
``` sh
tar --extract -f cross-gcc-10.3.0-pi_0-1.tar.gz
```

Then when building we can point to use that gcc

## Build

``` sh
env CGO_ENABLED=1 CC=~/cross-pi-gcc-10.3.0-0/bin/arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 go build -o deky
scp deky pi@raspberrypizero.local:~/deky
```

### Use

``` sh
curl raspberrypizero.local:8080/Hello
```

Clear display including blinking cursor and run the app:
``` sh
sudo sh -c "TERM=linux setterm -foreground black -cursor off -clear all >/dev/tty0"
sudo ./deky
```

## On PC

### Share internet with RpiZero

> https://solarianprogrammer.com/2018/12/07/raspberry-pi-zero-internet-usb/

## On RpiZero

### Installation

``` sh
cd Raspberry-Pi-Installer-Scripts
sudo python3 adafruit-pitft.py --display=22 --rotation=90 --install-type=fbcp
```

## Hardware components

[Raspberry Pi Zero](https://www.raspberrypi.com/products/raspberry-pi-zero/)

[Adafruit 2.2" PiTFT HAT - 320x240 Display](https://learn.adafruit.com/adafruit-2-2-pitft-hat-320-240-primary-display-for-raspberry-pi)
