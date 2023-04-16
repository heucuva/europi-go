# Complex Envelope

A random gate generator based on YouTube video observations made about the operation of the After Later Audio/mxmxmx uO_C (micro Ornament & Crime) module.

## Scope of This App

The scope of this app is to drive the CV 1 and CV 2 outputs as enveloped values using complex configurations.

### Inputs

- Digital Input = gate input (optional / recommended, see below)
- Analogue Input = CV for Attack parameter attenuation (applied to both envelopes)

### Outputs

- CV 1 = Envelope 1 output
- CV 2 = Envelope 2 output

## Using Complex Envelope

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Complex Envelope configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi.

While Complex Envelope is operating, you can manually trigger an envelope by pressing, holding for a desired period, then releasing Button 1 for Envelope 1 or Button 2 for Envelope 2. However, if you hold Button 2 for 650ms or more, it will instead transition the display to the next menu.

#### Complex Envelope Configuration

By default, the settings of Complex Envelope are:

- Mode: AD
- AttackMode: Quartic
- ReleaseMode: Exponential
- Attack: +127
- Decay: +127

When on the Complex Envelope Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

The current envelope being displayed/edited will be shown in the upper-right corner of the screen and may be cycled through by briefly pressing Button 2 on the EuroPi.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- After Later Audio
- mxmxmx
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
