# Cascade LFO

A cascading LFO generator based on YouTube video observations made about the operation of the Instruo Ochd module.

## Scope of This App

The scope of this app is to drive the CV 1 through CV 6 outputs as triangle waves running at consecutively slower rates.

### Inputs

- Digital Input = reset the LFO chain.
- Analogue Input = rate modification CV modified by the internal attenuverter

### Outputs

- CV 1 = triangle wave with 1:1 rate
- CV 2 = triangle wave with 1:2 rate
- CV 3 = triangle wave with 1:4 rate
- CV 4 = triangle wave with 1:8 rate
- CV 5 = triangle wave with 1:16 rate
- CV 6 = triangle wave with 1:32 rate

## Using Cascading LFO

### Cascading LFOs

While Cascade LFO is operating, you can turn Knob 1 to change the base rate of the LFO chain.

Knob 2 can be turned to change the attenuversion of the modifications to the rate made by the Analogue Input. This range is -1.0 to 1.0.

### TODO

More functionality and menu-diving to come...

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Instruo
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
