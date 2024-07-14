Article https://www.reddit.com/r/rocksmith/comments/zi17ng/success_connecting_boss_katana_to_rs2014_as_an/

Success connecting Boss Katana to RS2014 as an audio interface (step by step)
Hey friends,

I finally got Rocksmith 2014 working with my setup after many hours of combing over old posts, trying different options and pulling most of my hair out. The solution that ended up working for me was a combination of tips from several different sources, so I figured I'd put the steps into a single set of instructions just in case this helps anyone else.

(...And also to help myself if I need to see these steps again, let's be honest.)

My setup:
Electric guitar + regular patch cord --> Boss Katana 50 MkII + regular printer cable --> Windows 10 + Steam version of Rocksmith 2014, using Direct Connect Mode.

Download RSMods and make sure you get the most up to date version. https://github.com/Lovrom8/RSMods/releases The file you want will most likely be called RS2014-Mod-Installer.exe.

Run the installer. You should only have to do this once.

After the installer has been run, you should be directed right to the RSMods interface. If you ever need to open RSMods again, for example to try a new mod, you can get back here through Program Files (x86)\Steam\steamapps\common\Rocksmith2014\RSMods.
If you get an error message whenever you try to open RSMods, check to ensure you've got the latest version.

In the RSMods interface, navigate to "Set And Forget Mods," then "Misc." Click "Add Direct Connect Mode." You should get a message when this is successful. You can now close RSMods (or tweak any other settings you want).

Plug the Katana into your PC and turn it on. If this is your first time plugging the amp into the PC it will install the necessary drivers automatically. A firmware update was released earlier this year. If you haven't updated the Katana yet, follow these instructions: https://www.youtube.com/watch?v=9O-z21cczyg

Go to the start menu and search "Katana." You want the KATANA app, not the Boss Tone Studio. This will open the Katana driver settings.

Change the SAMPLE setting to 48000 Hz. If you're unable to change this, uncheck the "Match with the ASIO sample rate" box. Hit Apply and close. DO NOT SKIP THIS STEP. If you do, when you attempt to use Direct Connect Mode in Rocksmith your options will most likely sound like either screeching feedback from Hell, or no signal at all.

Right click on the sound icon in your taskbar and go to "Sounds" then "Recording." Disable your default microphone. (In other sets of instructions, they also advise to disable the PRIMARY Katana channel, but I found that I needed both Katana channels enabled in order for Rocksmith to pick up the signal.)

(Optional) Right click on the sound icon in your taskbar and go to "Open Sound Settings." Under Input, go to Device properties. Just for shits and giggles, you can also click on "Additional device properties" > "Advanced" to confirm that a) the sampling rate is at 48000 Hz and b) both checkboxes under "Exclusive Mode" are checked.
I don't know why, but whenever I check the Device properties of either the PRIMARY or the SECONDARY Katana inputs, they're always set to something really low (like 17 or 22). I upped this to 100, but I don't know if that actually helps anything, and it seems to always reset to something low whenever I check it.

Start Rocksmith. If your game is set to a weird resolution you can change that in the Rocksmith.ini file. Also, don't forget that the F11 key is your friend here. Took me stupidly long to remember that one.

Create your profile and answer the questions. When it comes time to select an input, choose Direct Connect. When I did this, it brought up both the PRIMARY and SECONDARY Katana channels, but only the SECONDARY one provided any signal. Nevertheless, I needed both channels to be enabled for Rocksmith to recognize the SECONDARY channel. You'll know you're getting signal through when you strum the guitar and the bar fills up.

Once you select this mode, you'll be brought into the Calibration menu. Follow MasterSh4k3's advice: https://www.reddit.com/r/rocksmith/comments/ho6sud/psa_you_might_be_calibrating_wrong_and_it_makes_a/ and only mute the strings briefly (not holding them down) to get the best results during the "Mute the Strings" part.

If you're lucky, you'll be able to go through the Calibration prompts as usual... if you're like me, your guitar won't be loud enough to fill up the meter. In that case, go on to Step 13.

13. Check to make sure your guitar volume is turned all the way up and that you're using the bridge pickup. To be honest, I didn't notice much difference based on the tone knob, so I didn't bother turning it all the way up. Playing with the gain and the volume on the Katana itself also didn't seem to affect anything being sent to the computer, but if you want to experiment, go ahead I guess.

14. Assuming your Calibration still doesn't work, sigh loudly and go back to the previous menu. Since this was my first time setting up a profile, the only way I could see to skip the Calibration step was to choose "Disconnected" mode at first.

15. Once in the main menu, go to Tools > Audio Settings > Input Gain Override. (If you were able to successfully install Direct Connect Mode, this menu option should be visible.) My default was set to -1.2db. Check the box and use the slider to increase the gain. Pressing Enter didn't seem to do anything here, but once I pressed Esc to get back to the Audio Settings menu, my preference was saved.

16. Once you change the Input Gain Override, it should take you to the Calibration step automatically. If not though, you can always get to it by going to Tools > Tuner, selecting any tuning, and then opening Calibration (you don't have to tune first). You can play around with the Input Gain Override setting, but I only needed to slide it up to 0db to get the game to actually fill up the Calibration meter.

17. You should be good to go now! You can either mute your guitar in game and hear the sound through the Katana, mute the Katana and just listen to the in game guitar (you can also put the Katana on standby mode), or mix in a bit of both.

So far I haven't noticed any glaring issues with latency, notes buzzing, etc. And it has put the biggest smile on my face, so all that struggling to get this up and running was (mostly) worth it. :)

#rocksmith #guitar