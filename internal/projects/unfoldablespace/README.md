# Unfoldable Space

A project dedicated to reproducing Krell-style CVs and gates necessary to generate audio reminiscent of [State Azure](https://www.youtube.com/channel/UClKIjbgtWGzHtXhBDS_I0pg)'s [Folding Space](https://www.youtube.com/watch?v=6JeZR13dLLI).

Additional dissection of State Azure's generative ambient patch was performed by [VisionsMusicGroup](https://www.youtube.com/@VisionsMusicGroup) within the [Patch from Scratch - Un-folding Space (a State Azure cover)](https://www.youtube.com/watch?v=3bksW2QjQ-0) video.

## Scope of This App

The scope of this app is to drive the first voice (Mutable Instruments Plaits - or its many derivatives - and XAOC Devices Belgrad) as accurately as possible with just a EuroPi. An additional output for sending to the V/Oct input of XAOC Devices Belgrad is provided on CV6.

Some additional modules down-chain from Plaits and Belgrad are necessary.

### Inputs

- Digital Input = clock input

In Un-folding Space, a clock rate of 120 BPM is suggested, though any rate can be supplied. This clock drives the trigger and skipper systems, which ultimately drive the envelope and arpeggiator circuits, which generate the final values sent to the output jacks on the EuroPi.

### Outputs

- CV 1 = Plaits V/Oct
- CV 2 = Plaits Level
- CV 3 = Plaits Timbre
- CV 4 = Plaits Harmo
- CV 5 = Plaits Morph
- CV 6 = Belgrad V/Oct

### What Else is Needed

- Mutable Instruments Plaits (or similar)
- XAOC Devices Belgrad
- LFO with ~8Hz rate
  - Erica Synths Pico VCO in LFO mode with sine waveform suffices

## Configuration

Set up your modular system in this way:

1. Wire EuroPi CV 1 to Plaits V/Oct.
2. Wire EuroPi CV 2 to Plaits Level.
3. Wire EuroPi CV 3 to Plaits Timbre.
4. Wire EuroPi CV 4 to Plaits Harmo.
5. Wire EuroPi CV 5 to Plaits Morph.
6. Wire EuroPi CV 6 to Belgrad V/Oct.
7. Wire Plaits Out to Belgrad Input.
8. Wire Belgrad Output to subsequent modules necessary (or send to output mixer).
9. Wire LFO output to Plaits FM.
10. Set Plaits Mode to Green or Yellow/Orange (firmware 1.2) in the top-most ('cloud') position. Original uses Green mode, but Yellow/Orange is also good.
11. Set Plaits Frequency knob to 12 o'clock position.
12. Set Plaits Harmonics knob to 12 o'clock position.
13. Set Plaits Timbre knob to 10 o'clock position.
14. Set Plaits Morph knob to 2 o'clock position.
15. Set Plaits Timbre attenuverter knob to 2 o'clock position.
16. Set Plaits Morph attenuverter knob to 10 o'clock position.
17. Set Plaits FM attenuverter knob to just slightly counter-clockwise/anti-clockwise of 12 o'clock position.
18. Set LFO to about 8Hz rate (~7.5Hz is recommended). Ensure waveform of LFO is a sine wave.
19. Set Belgrad Freq knob to 12 o'clock position.
20. Set Belgrad Level knob to maximum position (`10`).
21. Set Belgrad Span knob to 12 o'clock position.
22. Set Belgrad modulation switch to the SM position.
23. Set Belgrad Reso knob to 2 o'clock position.
24. Set Belgrad Balance knob to 12 o'clock position.
25. Set Belgrad FM slider to middle position (`5`).
26. Set Belgrad Span slider to middle position (`5`).
27. Set Belgrad Mode selector knob to Double Bandpass (`BB`) mode.

Optional:

28. Wire a Clock source to EuroPi Digital input.

## Using Unfoldable Space

### Performance Clock

While Unfoldable Space is operating, you can toggle between using the external clock (default mode at startup) and the internal 120 BPM clock by pressing Button 1 on the EuroPi. When the mode is active, you will be informed by seeing a small bar ( `_` ) in the upper-left corner of the display.

## Internal *'Module'* Configuration

In order to simulate a complex configuration of modules and their interactions, multiple 'Modules' have been designed and internally wired together so they result in a close approximation of the original modules. No code from these modules has been used and only deciphering of their respective manuals and observations made from YouTube videos was done to identify their functionality.

### Internal Wiring

#### Clock Generator

- Clock Out -> Random Gates Trigger Input

#### Random Gates

- Gate 1 Out -> Random Skips Gate 1 Input

#### Three-Phase LFO

- 0-Degree Output -> Random Skips CV 1 Input

#### Random Skips

- Gate 1 Out -> Complex Envelope 1 Trigger Input
- Gate 1 Out -> Complex Envelope 2 Trigger Input
- Gate 1 Out -> Complex Arpeggiator Clock Input

#### Complex Arpeggiator

- Arpeggio Out -> EuroPi CV 1 Output

#### Cascade LFO

- CV 4 Out -> EuroPi CV5 Output
- CV 5 Out -> Complex Envelope 1 CV Input

#### Complex Envelope

- Envelope 1 Out -> EuroPi CV 2 Output
- Envelope 2 Out -> EuroPi CV 6 Output
- Envelope 2 Out -> Cascade LFO CV Input

#### Complex Random

- Sample A Out -> EuroPi CV 3 Output
- Sample B Out -> EuroPi CV 4 Output

## Factory Default Settings

### Clock Generator

- BPM: 120.0
- Enabled: False

### Random Gates

- Gate Duration: 200ms

### Three-Phase LFO

- Wave Mode: Sine
- 3-Phi Rate: 20%
- Skew Rate: 0%
- Skew Shape: 5%

### Random Skips

- Chance: 60%

### Complex Arpeggiator

- Pattern: Brownian
- Chord Mode: C Major
- Quantizer Mode: Round
- Range: 1.0 Octave(s)
- Pitch: C-4 (4.0 V/Oct)

### Cascade LFO

- Rate: 80%
- Rate Attenuverter: 90%

### Complex Envelope

- Envelope 1 Mode: AD (attack-decay)
- Envelope 1 Attack Mode: Linear
- Envelope 1 Release Mode: Exponential
- Envelope 1 Attack Time: 66%
- Envelope 1 Release Time: 66%
- Envelope 2 Mode: AD (attack-decay)
- Envelope 2 Attack Mode: Linear
- Envelope 2 Release Mode: Exponential
- Envelope 2 Attack Time: 50%
- Envelope 2 Release Time: 50%

### Complex Random

- Sample A Attenuator: 60%
- Integration Slope: 0%
- Gate Density: 50%
- Pulse Stage Division: 1:1
- Sample B Attenuator: 20%
- Sample B Slew: 30%
- Clock Speed: 40%
- Clock Range: Full

### TODO

More functionality and menu-diving to come...

## Special Thanks

- State Azure
- VisionsMusicGroup
- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Mutable Instruments / Emilie Gillet
- XAOC Devices
- Erica Synths
- Perfect Circuit
- Pusherman
- After Later Audio
- Antumbra
- Mordax Systems
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation

## Additional Thanks

- 2HP
- 4ms
- Acid Rain Technology
- ACL
- ALM
- Bastl
- Befaco
- Doepfer
- Expert Sleepers
- Instruo
- Intelljel
- Klavis
- Ladik
- mxmxmx
- WMD/SSF
