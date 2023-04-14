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

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Cascading LFO configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi.

#### Cascading LFO Configuration

By default, the settings of Cascading LFO are:
- Rate Attenuverter: +80%
- Rate: 16.0Hz

When on the Cascading LFO Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Instruo
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
