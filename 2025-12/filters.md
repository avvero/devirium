![alt text](square-wave.png)

Filters not only change a waveform by attenuation, but distort it by individually phase-shifting the harmonics within it.

A low-pass filter passes unhindered all the frequencies below a 'cutoff' frequency while attenuating all those above it. Let's take [[square wave]] (100Hz + 200Hz + 300Hz + 400Hz + 500Hz + ...). Now let's say that our simple RC filter has a cutoff frequency of 400Hz, and imagine what would happen to our square wave if the filter's phase response was zero at all frequencies. This is quite simple: the fundamental and first overtone of the square wave (the harmonics at 100Hz and 300Hz) would be unattenuated, but all the overtones at 500Hz and above would be attenuated according to the filter's response. The resulting waveform (and you'll have to trust me on this) is shown in Figure 9

[[Synth Secrets]]

#synth