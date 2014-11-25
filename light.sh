#!/bin/bash
gpio -g mode 27 out; gpio -g write 27 1; sleep 0.2s; gpio -g write 27 0
