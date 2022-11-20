import struct

offset = 96

exploit = ""
exploit += "A" * offset
exploit += "BBBB"
exploit += "C" * 20

print(exploit)