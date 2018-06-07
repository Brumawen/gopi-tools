import argparse
import RPi.GPIO as GPIO
from gpiozero import MCP3008
from time import sleep

parser = argparse.ArgumentParser(description='Read values from MCP3008 chip')
parser.add_argument('channels', metavar='N', type=int, nargs='+', help='the channels to read')
args = parser.parse_args()

values = []
# retrieve each value from the required ADC pins
# for each pin, read 50 values then get the average of the range
# without the first and last 10 values
for c in args.channels:
    dev = MCP3008(c)
    a = []
    for i in range(0,50):
        a.append(dev.value*100)
    s = sorted(a)[10:41]
    av = sum(s) / float(len(s))
    values.append(av)

# return the values to the calling function
print("\t".join(map(str,values)))