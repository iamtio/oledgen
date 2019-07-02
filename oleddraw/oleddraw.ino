
#include <SPI.h>
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

#define SCREEN_WIDTH 128 // OLED display width, in pixels
#define SCREEN_HEIGHT 64 // OLED display height, in pixels

const size_t bytesLineSize = SCREEN_WIDTH / 8;

// Declaration for an SSD1306 display connected to I2C (SDA, SCL pins)
#define OLED_RESET     4 // Reset pin # (or -1 if sharing Arduino reset pin)
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

void setup() {
  Serial.begin(115200);

  // SSD1306_SWITCHCAPVCC = generate display voltage from 3.3V internally
  if(!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) { // Address 0x3C for 128x32
     Serial.println(F("SSD1306 allocation failed"));
    for(;;); // Don't proceed, loop forever
  }

  // Show initial display buffer contents on the screen --
  // the library initializes this with an Adafruit splash screen.
  display.display();
  delay(2000); // Pause for 1 second

  // Clear the buffer
  display.clearDisplay();
}

void loop() {
  byte b;
  display.clearDisplay();
  for(size_t readed = 0; readed < (SCREEN_WIDTH*SCREEN_HEIGHT/8); readed++) {
    while(!(Serial.available() > 0)){
      delay(10);
    }
    b = Serial.read();
    for(int bit = 0; bit < 8; bit++){
      if(((b >> bit) & 1) != 0)
        display.drawPixel((readed*8+bit) % SCREEN_WIDTH, readed / bytesLineSize, WHITE);
    }
  }
  display.display();
}
