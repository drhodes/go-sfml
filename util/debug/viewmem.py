#!/usr/bin/env python
"""
Show how to make date plots in matplotlib using date tick locators and
formatters.  See major_minor_demo1.py for more information on
controlling major and minor ticks
"""
import datetime
from pylab import figure, show
import sys

files = sys.argv[1:]


plots = []
for f in files:
    temp = []
    for line in open(f):
        num = line.strip()
        if num:
            temp.append(float(num))
    plots.append(temp)

normal = []
for p in plots:
    temp = []
    m = max(p)
    for num in p:
        temp.append(num/float(m))
    normal.append(temp)

fig = figure()

shum = []
for norm in normal:
    shum.append( fig.add_subplot(111) )
    ax = shum[-1]
    ax.plot(norm)
    ax.autoscale_view()
    ax.grid(True)
    

show()
