import os
import sys

from io import StringIO
import matplotlib.pyplot as plt


files = [p for p in os.listdir() if "julia" in p and not p.endswith(".res")]
times = []

for file in files:
    os.system("go run " + file + " > " + file + ".res") #couldn't capture stdout on my pc for some reason. Which is kinda cringe
    fs = open(file + ".res", "r")
    times.append([int(i) for i in fs.readlines()])
    fs.close()

files = [p for p in os.listdir() if "julia" in p and p.endswith(".res")]
for file in files:
    os.remove(file)

for index, timing in enumerate(times):
    x = [i for i in range(len(timing))]
    plt.plot(x, timing, label=files[index])

plt.xlabel("Image (n)")
plt.ylabel("Finish Time (ms)")
plt.title("Julia Parallellization Comparison")
plt.legend()
plt.show()