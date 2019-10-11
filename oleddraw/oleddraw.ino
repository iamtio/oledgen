
#include <SPI.h>
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

#define SCREEN_WIDTH 128 // OLED display width, in pixels
#define SCREEN_HEIGHT 64 // OLED display height, in pixels

const size_t bytesLineSize = SCREEN_WIDTH / 8;
const int timeout = 10000;
// Declaration for an SSD1306 display connected to I2C (SDA, SCL pins)
#define OLED_RESET     4 // Reset pin # (or -1 if sharing Arduino reset pin)
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

void drawNoData(){
  display.clearDisplay();

  display.setTextSize(1);      // Normal 1:1 pixel scale
  display.setTextColor(WHITE); // Draw white text
  display.setCursor(0, 0);     // Start at top-left corner
  display.cp437(true);         // Use full 256 char 'Code Page 437' font

  // Not all the characters will fit on the display. This is normal.
  // Library will draw what it can and the rest will be clipped.
  display.println("No data.\nRun host application\non your PC.");
  display.setCursor(55, 54);
  display.println("TIO (c) 2019");
  display.display();
}
void setup() {
  Serial.begin(115200);

  // SSD1306_SWITCHCAPVCC = generate display voltage from 3.3V internally
  if(!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) { // Address 0x3C for 128x32
     Serial.println(F("SSD1306 allocation failed"));
    for(;;); // Don't proceed, loop forever
  }
  
//  delay(2000); // Pause

  // Clear the buffer
  drawNoData();
}

void loop() {
  byte b;
  int c = 0;
  display.clearDisplay();
  for(size_t readed = 0; readed < (SCREEN_WIDTH*SCREEN_HEIGHT/8); readed++) {
    while(!(Serial.available() > 0)){
      delay(10);
      if(c++ > (timeout/10)){
        drawNoData();
        display.clearDisplay();
        delay(10);
      }
    }
//    display.clearDisplay();
    b = Serial.read();
    for(int bit = 0; bit < 8; bit++){
      if(((b >> bit) & 1) != 0)
        display.drawPixel((readed*8+bit) % SCREEN_WIDTH, readed / bytesLineSize, WHITE);
    }
  }
  display.display();
}
