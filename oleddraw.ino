#include <SPI.h>
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

#define SCREEN_WIDTH 128 // OLED display width, in pixels
#define SCREEN_HEIGHT 64 // OLED display height, in pixels

// Declaration for an SSD1306 display connected to I2C (SDA, SCL pins)
#define OLED_RESET     4 // Reset pin # (or -1 if sharing Arduino reset pin)
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

const int imageBufferSize = (SCREEN_WIDTH * SCREEN_HEIGHT)/8;
byte imageBuffer[imageBufferSize];

void setup() {
  Serial.begin(115200);

  // SSD1306_SWITCHCAPVCC = generate display voltage from 3.3V internally
  if(!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) { // Address 0x3C for 128x32
    // Serial.println(F("SSD1306 allocation failed"));
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
    for(size_t readed = 0; readed < imageBufferSize; readed++) {
        while(!(Serial.available() > 0)){
            delay(10);
        }
        imageBuffer[readed] = Serial.read();
    }
    drawbuffer();
}

void drawbuffer(void) {
  display.clearDisplay();

  display.drawBitmap(0, 0, imageBuffer, SCREEN_WIDTH, SCREEN_HEIGHT, 1);
  display.display();
}