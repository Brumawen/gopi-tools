package gopitools

import rpio "github.com/stianeikeland/go-rpio"
import "time"

// Command Flags
const lcdClearDisplay int = 0x01
const lcdReturnHome int = 0x02
const lcdEntryModeSet int = 0x04
const lcdDisplayControl int = 0x08
const lcdCursorShift int = 0x10
const lcdFunctionSet int = 0x20
const lcdSetCGRAMAddr int = 0x40
const lcdSetDDRAMAddr int = 0x80

// Entry Flags
const lcdEntryRight int = 0x00
const lcdEntryLeft int = 0x02
const lcdEntryShiftIncrement int = 0x01
const lcdEntryShiftDecrement int = 0x00

// Control Flags
const lcdDisplayOff int = 0x00
const lcdDisplayOn int = 0x04
const lcdCursorOff int = 0x00
const lcdCursorOn int = 0x02
const lcdBlinkOff int = 0x00
const lcdBlinkOn int = 0x01

// Move Flags
const lcdDisplayMove int = 0x08
const lcdCursorMove int = 0x00
const lcdMoveRight int = 0x04
const lcdMoveLeft int = 0x00

// Function Set Flags
const lcd4BitMode int = 0x00
const lcd8BitMode int = 0x10
const lcd1Line int = 0x00
const lcd2Line int = 0x08
const lcd5x8Dots int = 0x00
const lcd5x10Dots int = 0x04

var lcdRowOffsets = [...]int{0x00, 0x040, 0x14, 0x54}

// CharDisplay allows you to control a HDD44780 compatible LCD display.
// Default GPIO numbers:
// Register Select (RS) = 21
// Clock Enable (EN)    = 20
// Data 0 to 3          = Not used in 4bit operation
// Data 4               = 26
// Data 5               = 19
// Data 6               = 13
// Data 7               = 6
//
// Defaults to a 8x2 character display.
//
// Port of https://github.com/adafruit/Adafruit_Python_CharLCD
//
type CharDisplay struct {
	// Public values
	GpioRS int
	GpioEN int
	GpioD4 int
	GpioD5 int
	GpioD6 int
	GpioD7 int
	Cols   int
	Lines  int
	// Private values
	isInitialized   bool
	pinRS           rpio.Pin
	pinEN           rpio.Pin
	pinD4           rpio.Pin
	pinD5           rpio.Pin
	pinD6           rpio.Pin
	pinD7           rpio.Pin
	displayControl  int
	displayFunction int
	displayMode     int
}

// AutoScroll will 'right justify' text from the cursor if set to true,
// otherwise it will 'left justify' the text
func (d *CharDisplay) AutoScroll(autoScroll bool) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	if autoScroll {
		d.displayMode = d.displayMode | lcdEntryShiftIncrement
	} else {
		d.displayMode = d.displayMode &^ lcdEntryShiftIncrement
	}
	d.writeValue(lcdEntryModeSet|d.displayMode, false)

	return nil
}

// BlinkCursor turns cursor blinking on or off.
func (d *CharDisplay) BlinkCursor(blink bool) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	if blink {
		d.displayControl = d.displayControl | lcdBlinkOn
	} else {
		d.displayControl = d.displayControl &^ lcdBlinkOn
	}
	d.writeValue(lcdDisplayControl|d.displayControl, false)
	return nil
}

// Clear clears the LCD display.
func (d *CharDisplay) Clear() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	// Send command to clear display
	d.writeValue(lcdClearDisplay, false)
	// Clearing the display takes a long time
	time.Sleep(3000 * time.Microsecond)
	return nil
}

// Close releases and unmaps GPIO memory.
func (d *CharDisplay) Close() {
	if d.isInitialized {
		rpio.Close()
		d.isInitialized = false
	}
}

// EnableDisplay enables or disables the LCD display.
func (d *CharDisplay) EnableDisplay(enable bool) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	if enable {
		d.displayControl = d.displayControl | lcdDisplayOn
	} else {
		d.displayControl = d.displayControl &^ lcdDisplayOn
	}
	d.writeValue(lcdDisplayControl|d.displayControl, false)
	return nil
}

// Home moves the cursor to its home (fist line and first column)
func (d *CharDisplay) Home() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	d.writeValue(lcdReturnHome, false)
	time.Sleep(3000 * time.Microsecond)
	return nil
}

// Init initializes the display ready for use.
func (d *CharDisplay) Init() error {
	if d.isInitialized {
		return nil
	}
	// Set the GPIO pin numbers
	if d.GpioRS <= 0 {
		d.GpioRS = 21
	}
	if d.GpioEN <= 0 {
		d.GpioEN = 20
	}
	if d.GpioD4 <= 0 {
		d.GpioD4 = 26
	}
	if d.GpioD5 <= 0 {
		d.GpioD5 = 19
	}
	if d.GpioD6 <= 0 {
		d.GpioD6 = 13
	}
	if d.GpioD7 <= 0 {
		d.GpioD7 = 6
	}
	// Set the Columns and Lines
	if d.Cols <= 0 {
		d.Cols = 8
	}
	if d.Lines <= 0 {
		d.Lines = 2
	}

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return err
	}
	// Get the GPIO pins
	d.pinRS = rpio.Pin(d.GpioRS)
	d.pinRS.Output()
	d.pinEN = rpio.Pin(d.GpioEN)
	d.pinEN.Output()
	d.pinD4 = rpio.Pin(d.GpioD4)
	d.pinD4.Output()
	d.pinD5 = rpio.Pin(d.GpioD5)
	d.pinD5.Output()
	d.pinD6 = rpio.Pin(d.GpioD6)
	d.pinD6.Output()
	d.pinD7 = rpio.Pin(d.GpioD7)
	d.pinD7.Output()

	// Initialize the display
	d.writeValue(0x33, false)
	d.writeValue(0x32, false)
	// Initialize display control, function and mode registers
	d.displayControl = lcdDisplayOn | lcdCursorOff | lcdBlinkOff
	d.displayFunction = lcd4BitMode | lcd1Line | lcd2Line | lcd5x8Dots
	d.displayMode = lcdEntryLeft | lcdEntryShiftDecrement
	d.writeValue(lcdDisplayControl|d.displayControl, false)
	d.writeValue(lcdFunctionSet|d.displayFunction, false)
	d.writeValue(lcdEntryModeSet|d.displayMode, false)

	d.isInitialized = true
	return nil
}

// Message writes the text to the display.
// A NewLine character '\n' in the text moves the rest of the text
// tn the next line.
func (d *CharDisplay) Message(text string) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	row := 0
	col := 0
	if d.displayMode&lcdEntryLeft == 0 {
		col = d.Cols - 1
	}
	d.SetCursor(col, row)
	for _, char := range text {
		if char == '\n' {
			row = row + 1
			if d.displayMode&lcdEntryLeft > 0 {
				col = 0
			} else {
				col = d.Cols - 1
			}
			d.SetCursor(col, row)
		} else {
			// Write the character to the display
			col = col + 1
			d.writeValue(int(char), true)
		}
	}
	return nil
}

// MoveLeft moves the display left by one position
func (d *CharDisplay) MoveLeft() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	d.writeValue(lcdCursorShift|lcdDisplayMove|lcdMoveLeft, false)
	return nil
}

// MoveRight moves the display right by one position
func (d *CharDisplay) MoveRight() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	d.writeValue(lcdCursorShift|lcdDisplayMove|lcdMoveRight, false)
	return nil
}

// SetCursor moves the cursor to an explicit column and row position
func (d *CharDisplay) SetCursor(col int, row int) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	if row >= d.Lines {
		row = d.Lines - 1
	}
	// Set the location
	d.writeValue(lcdSetDDRAMAddr|(col+lcdRowOffsets[row]), false)
	return nil
}

// SetLeftToRight sets the text direction as left to right
func (d *CharDisplay) SetLeftToRight() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	d.displayMode = d.displayMode | lcdEntryLeft
	d.writeValue(lcdEntryModeSet|d.displayMode, false)
	return nil
}

// SetRightToLeft sets the text direction as right to left
func (d *CharDisplay) SetRightToLeft() error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	d.displayMode = d.displayMode &^ lcdEntryLeft
	d.writeValue(lcdEntryModeSet|d.displayMode, false)
	return nil
}

// ShowCursor shows or hides the cursor.
func (d *CharDisplay) ShowCursor(show bool) error {
	if !d.isInitialized {
		err := d.Init()
		if err != nil {
			return err
		}
	}
	if show {
		d.displayControl = d.displayControl | lcdCursorOn
	} else {
		d.displayControl = d.displayControl &^ lcdCursorOn
	}
	d.writeValue(lcdDisplayControl|d.displayControl, false)
	return nil
}

// writeValue writes an 8-bit value in character or data mode.
// Value should be an int value from 0-255.
func (d *CharDisplay) writeValue(value int, isCharMode bool) {
	// Sleep for 1 millisecond.  This is to prevent writing too quickly
	time.Sleep(1 * time.Microsecond)
	// Set the RS pin to character or data mode
	d.setPin(d.pinRS, isCharMode)
	// Write the upper 4 bits
	d.setPin(d.pinD4, ((value>>4)&1) > 0)
	d.setPin(d.pinD5, ((value>>5)&1) > 0)
	d.setPin(d.pinD6, ((value>>6)&1) > 0)
	d.setPin(d.pinD7, ((value>>7)&1) > 0)
	d.pulseEnable()
	// Write the lower 4 bits
	d.setPin(d.pinD4, (value&1) > 0)
	d.setPin(d.pinD5, ((value>>1)&1) > 0)
	d.setPin(d.pinD6, ((value>>2)&1) > 0)
	d.setPin(d.pinD7, ((value>>3)&1) > 0)
	d.pulseEnable()
}

// pulseEnable pulses the clock enable line off, on, off to send command.
func (d *CharDisplay) pulseEnable() {
	d.pinEN.Low()
	time.Sleep(1 * time.Microsecond) // Enable pulse must be > 450ns
	d.pinEN.High()
	time.Sleep(1 * time.Microsecond) // Enable pulse must be > 450ns
	d.pinD4.Low()
	time.Sleep(1 * time.Microsecond) // Command needs > 37us to settle
}

func (d *CharDisplay) setPin(pin rpio.Pin, output bool) {
	if output {
		pin.High()
	} else {
		pin.Low()
	}
}
