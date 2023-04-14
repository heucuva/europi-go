# Complex Arpeggiator

A complex arpeggio generator based on YouTube video observations made about the operation of the ACL Sinfonion module.

## Scope of This App

The scope of this app is to drive the CV 1 output as a V/octave representation of the note.

### Inputs

- Digital Input = gate input (optional / recommended, see below)
- Analogue Input = V/octave pitch used as a centerpoint for the arpeggio

### Outputs

- CV 1 = V/octave output

## Using Complex Arpeggiator

### Triggering / Gating

While Complex Arpeggiator is operating, a gate input passed into the Digital Input of the EuroPi will trigger updates to the arpeggiator Alternatively, you can manually trigger an update to the arpeggiator by momentarily pressing Button 1 of the EuroPi while the display is showing the main display screen.

Regardless of how the arpeggiator is triggered, only the leading (rising) edge of the trigger will initiate the "sample-and-hold" for the new note.

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Complex Arpeggiator configuration

#### Main Display

The main display shows the voltage of the CV 1 output on the EuroPi as well as the the currently selected scale and the range (in +/- V/Octaves) of the arpeggiation.

#### Complex Arpeggiator Configuration

By default, the settings of Complex Arpeggiator are:
- Scale: C Major
- Pitch: 4.0 (C-4)
- Range: 1.0 (+/- 1 Octave)

Scales are as follows:
- "C lyd" = C Lydian
- "C maj" = C Major
- "C 7"   = C 7
- "C sus" = C Suspended
- "C hm5" = C Harmonic 5th
- "C dor" = C Dorian
- "C min" = C Minor
- "C phr" = C Phrygian
- "C dim" = C Diminished
- "C aug" = C Augmented

When on the Complex Arpeggiator Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

### TODO

More functionality and menu-diving to come...

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- ACL
- Mathias Kettner
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
