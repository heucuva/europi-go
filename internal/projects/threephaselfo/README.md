# Three-Phase LFO

A three-phase LFO generator based on YouTube video observations made about the operation of the WMD/SSF Modbox module.

## Scope of This App

The scope of this app is to drive the CV 1 through CV 3 outputs as values sampled from a waveform at three phase points (0, 120, and 240 degrees).

Currently, only Sine waves are generated.

### Inputs

- Digital Input = reset the LFO.

### Outputs

- CV 1 = waveform value at 0 degrees
- CV 2 = waveform value at 120 degrees
- CV 3 = waveform value at 240 degrees

## Using Three-Phase LFO

### Three-Phase LFO

While Three-Phase LFO is operating, you can turn Knob 1 to change the base rate of the LFO.

Knob 2 can be turned to change the selected waveform. Currently, this does effectively nothing, as there is only 1 waveform provided (sine).

### TODO

More functionality and menu-diving to come...

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- WMD
- SSF
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
