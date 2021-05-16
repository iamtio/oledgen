# What is it?
PC resources monitoring device. Such as RAM, CPU usage...
# How it looks
![photo_2021-05-16_17-28-02](https://user-images.githubusercontent.com/597141/118400876-2fd2cf00-b66c-11eb-87f1-4fcf0e16217e.jpg)

# Parts
- Arduino Pro Micro 5v - https://www.sparkfun.com/products/12640
- 0.96 I2C OLED Screen - https://www.winstar.com.tw/products/oled-module/graphic-oled-display/4-pin-oled.html
- 3D printed case from `./oledcase.scad` (OpenSCAD)
- Micro USB cable

# Directories and files
- `*.go` - Software for collecting resource metrics and draw them on device
- `./oleddraw` - Device firmware (Arduino)
- `./oledcase.scad` - OopeSCAD Model of device case for 3D printing