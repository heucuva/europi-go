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

While Complex Arpeggiator is operating, you can manually trigger an update to the arpeggiator by pressing Button 1.

Regardless of how the arpeggiator is triggered, only the leading (rising) edge of the trigger will initiate the "sample-and-hold" for the new note.

### Arpeggio Scale

Turning the Knob 1 potentiometer will change the scale that the arpeggiator uses to produce notes.

Available scales:
- C Lydian
- C Major
- C 7
- C Suspended
- C Harmonic_5
- C Dorian
- C Minor
- C Phrygian
- C Diminished
- C Augmented

### Note Range

Turning the Knob 2 potentiometer will change the range that the arpeggiator uses to produce notes within.

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
