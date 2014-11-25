#therewillbenoise

Rob Szumlakowski
Allan Baril

This application is a plugin for the CF cli that will monitor an application running
on CF and light a hockey goal light and output speech whenever several events occur:

 * Application started
 * Application stopped
 * Application crashed
 * Application scales to a new number of instances

The application itself uses the 'light.sh' and 'speech.sh' Bash scripts to light the
hockey goal light and produce speech.

We ran this plugin on a Raspberry Pi device.  The hockey goal light was wired to one
of the output pins on the Raspberri Pi device.  The 'light.sh' script uses the 'gpio'
program to trigger the output pins and turn the hockey goal light on and off.  If you
don't have access to a Raspberry Pi and a hockey goal light then just change 'light.sh'
so it does something else.

The 'speech.sh' script sends text to Google Translate and uses mpg123 to play the output.
