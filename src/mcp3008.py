from gpiozero import MCP3008
from time import sleep
values = []
for i in range(0,7):
    dev = MCP3008(i)
    values.append(dev.value*100)
print("\t".join(map(str,values)))

   
