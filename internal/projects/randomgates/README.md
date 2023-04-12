# Random Gates

A random gate generator based on YouTube video observations made about the operation of the Ladik R-110 module.

## Scope of This App

The scope of this app is to drive the CV 1 output as a gate change output based on a probability.

### Inputs

- Digital Input = clock input (optional, see below)

### Outputs

- CV 1 = Random Trigger 1 output
- CV 2 = Random Trigger 2 output
- CV 3 = Random Trigger 3 output
- CV 4 = Random Gate 1 output
- CV 5 = Random Gate 2 output
- CV 6 = Random Gate 3 output

## Using Random Gates

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Random Gates configuration
- Performance clock configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi as well as the enabled status of the internal performance clock.

While Random Gates is operating, you can toggle between using the external clock (default mode at startup) and the internal clock by pressing Button 1 on the EuroPi while on the main screen. When the internal clock mode is active, you will be informed by seeing a small bar ( `_` ) in the upper-left corner of the display.

#### Random Gates Configuration

By default, the settings of Random Gates are:
- Mode: 200ms

Modes are as follows:
- 1ms: Triggers will last a fixed 1ms
- 200ms: Triggers will last a fixed 200ms
- 1/4: Triggers will last a quarter of the length of the input gate
- 1/2: Triggers will last half of the length of the input gate
- 1:1: Triggers will last the same length of the input gate

When on the Random Gates Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

#### Performance Clock Configuration

By default, the settings of the Performance Clock are:
- Clock Rate: 120.0 BPM
- Gate Duration: 100.0 ms

When on the Performance Clock Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Ladik.eu
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
