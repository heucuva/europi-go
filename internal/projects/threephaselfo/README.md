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

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Three-Phase LFO configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi.

#### Three-Phase LFO Configuration

By default, the settings of Three-Phase LFO are:
- Wave: sine
- Rate: 1.0Hz
- SkewRate: 20.0Hz
- SkewShape: 0.05

When on the Three-Phase LFO Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- WMD
- SSF
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
