# Complex Random

A complex random number generator based on YouTube video observations made about the operation of the WMD/SSF Ultra-Random Analog module.

## Scope of This App

The scope of this app is to drive the CV 1 and CV 2 outputs as randomly-generated values.

### Inputs

- Digital Input = trigger a sampling of the value on the Analogue Input (optional)
- Analogue Input = value to sample (optional)

### Outputs

- CV 1 = randomly-generated value A
- CV 2 = randomly-genreated value B with slew

## Using Complex Random

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Complex Random configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi.

#### Complex Random Configuration

By default, the settings of Complex Random are:
- Sample Attenuator A (Attn.A): 60.0%
- Gate Density (GDense): 40.0%
- Pulse Stage Divider (PSD): 1
- Sample Attenuator B (Attn.B): 20.0%
- Sample Slew B (SlewB): 30.0%
- Clock Speed (CSpeed): 40.0%
- Clock Range (CRange): Full

When on the Complex Random Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

Clock Ranges are as follows:
- Full: 22050 Hz range
- Limited: 1470 Hz range

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- WMD
- SSF
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
