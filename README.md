#Leave

##Description

So you're sitting trying to understand how you've broken the dev environment on your extra special rails project and your girlfriend calls you.

	Her: Meet you in 10 mins at the end of the road.
	You: uhuh yeah.
	
You go back to what you were doing and by the time ten minutes has passed you've forgotten when the call was made. You're late.
This sort of stuff happens to me all the time. I could take my phone out and set an alarm or something but this wasn't awesome
enough. 'Leave' allows you to pop open a terminal and do this:

	$ leave in 10
	
In ten minutes a growl notification pops up and tells me that its time to go. Yay! You will never need time management again.

##Installation

You'll need:
	- Growl (http://growl.info)
	- Growl Python bindings (http://growl.info/documentation/developer/python-support.php)
	- Python (http://www.python.org)

To install simply clone or download the source and put 'leave' in /usr/bin (as it should be in you PATH already). You're done!

##Commands

Leave can be used in two ways. Firstly:

	leave in x
	
This will pop up a growl (remember not to close the terminal) in x minutes. You can also specify a time:

	leave at x
	
Here x should be of the format 'hours:mins'. My flat mate asked me why he couldn't specify seconds. I really hope no one actually micro manages their day to that degree.
	

	