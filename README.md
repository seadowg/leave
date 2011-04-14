#leave

##Description

So you're sitting trying to understand how you've broken the dev environment on your extra special rails project and your girlfriend calls you.

	Her: Meet you in 10 mins at the end of the road.
	You: uhuh yeah.
	
You go back to what you were doing and by the time ten minutes has passed you've forgotten when the call was made. You're late.
This sort of stuff happens to me all the time. I could take my phone out and set an alarm or something but this wasn't awesome
enough. 'leave' allows you to pop open a terminal and do this:

	$ leave in 10
	
In ten minutes a growl notification pops up and tells me that its time to go. Yay! You will never need time management again.

##Installation

You'll need:
	- Growl (http://growl.info)
	- Python (http://www.python.org)

To install, simply clone or download the source and run the install.sh script provided. It will prompt you for your password as it
copies files into /usr/bin. After its finished running you should be good to go!

##Commands

Leave can be used in two ways. Firstly:

	leave in x
	
This will pop up a growl in x minutes. You can also specify a time:

	leave at x
	
Here x should be of the format 'hours:mins'. My flat mate asked me why he couldn't specify seconds. I really hope no one actually
micro manages their day to that degree. So what happens if you decide the person you're meeting sucks and you can't be bothered
leaving? Just do:

	leave cancel
	
This will cancel the current countdown.

##Problems?

If you have any problems at all please don't hesitate to raise an issue on github or email me (oetzi101@gmail.com).	