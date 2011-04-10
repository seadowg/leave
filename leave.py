#!/usr/bin/env python
# encoding: utf-8
#
# Created by Callum Stott 2011

import sys
import time
import Growl

class Timer:
	def __init__(self, time):
		self.time = time
		
	def run(self):
		time.sleep(self.time)

class Growler:
	def __init__(self):
		self.reggrowl()
		
	def notify(self):
		self.growl.notify( "Leaving time", "Time to leave!", "Go go go!", None, True, None)
			
	def reggrowl(self):
		self.growl = Growl.GrowlNotifier( "Leave", ["Leaving time"] )
		self.growl.register()
		
def getsecsfrommins(input):
	return int(input) * 60
		
def getsecsuntil(input):
	current = time.localtime()
	target = time.strptime(input, "%H:%M")
	hour = target.tm_hour

	if target.tm_hour < current.tm_hour:
		hour = hour + 24
		
	hours = hour - current.tm_hour
	
	if hours == 0 and target.tm_min < current.tm_min:
		hours = 24
		
	return (hours * 3600) + ((target.tm_min - current.tm_min) * 60) - current.tm_sec

def main():
	growl = Growler()
	
	try:		
		if sys.argv[1] == 'in':
			secs = int(sys.argv[2]) * 60
			
			if secs < 0:
				raise ValueError
				
			print "Leaving in " + sys.argv[2] + " minutes..."
			
		elif sys.argv[1] == 'at':
			secs = getsecsuntil(sys.argv[2])
			print "Leaving at " + sys.argv[2]

		timer = Timer(secs)
		timer.run()
		growl.notify()
			
	except IndexError:
		print "Command syntax should be 'leave [in|at] time'"
		
	except ValueError:
		print "Values are incorrect:" 
		print "    'in' values are positive integers (minutes)" 
		print "    'at' values are of the form 'hours:mins'"

if __name__ == '__main__':
	main()

