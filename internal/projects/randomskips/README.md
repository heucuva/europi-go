# Random Skips

A random gate skipper based on YouTube video observations made about the operation of the Ladik S-090 module.

## Scope of This App

The scope of this app is to drive the CV 1 output as a gate output based on a percentage chance of 33%. When the input gate (or internal clock gate) goes high (CV >= 0.8V), then a random value is generated and compared against the chance that's provided - if the probability is sufficient enough, then the gate is let through for as long as it is still high on the input. The moment the gate goes low, the output also goes low and the detection process starts again.

### Inputs

- Digital Input = clock input (optional, see below)

### Outputs

- CV 1 = Random Gate output

## Using Random Skips

### Performance Clock

While Random Skips is operating, you can toggle between using the external clock (default mode at startup) and the internal 120 BPM clock by pressing Button 1 on the EuroPi. When the mode is active, you will be informed by seeing a small bar ( `_` ) in the upper-left corner of the display.

Turning the Knob 2 potentiometer will change the BPM of the internal clock and that value will be displayed while the clock is enabled.

### Chance

While Random Skips is operating, you can turn the Knob 1 potentiometer in order to change the Chance (`Chn`) variable.

### TODO

More functionality and menu-diving to come...

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Ladik.eu
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
