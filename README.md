# CS495_Golang_GUI
A project built while utilizing reference code as needed for Carroll College Computer Science Seminar which uses vector math and simple transformations to ray-cast a 3D play area onto a 2D screen, resulting in a "2.5D" type of "game". Uses Pixel graphics framework.
Reference code was used because I am still learning Go as a language and do not yet have the basics down; therefore, being able to read and understand uses of code in specific scenarios is important to the learning process at this point.

This project far surpasses the minimum requirements, because, although reference code was involved, it was simply to learn the concepts and methods for most of the program—the majority of code was written by myself independent of the reference code (but still after having looked at it). This is also excluding the fact that the reference code was not commented so it was up to me to decipher what was happening and understand it myself.

This project is built on Go 1.18.

In order to run this project:
1. clone this repository and navigate to the top level directory of the project in a command terminal of your choice.
2. build the project with the following command "go build src/main/*.go" (without the quotes).
3. run the generated executable file for your operating system.

<pre>
If you choose to run this executable from a command line terminal, you can pass flags to change how the game displays. 
This can be tweaked to your resolution. The flags are as follows:
1. -f     	                controls whether or not the game is fullscreen
2. -w &lt;width pixel count&gt;  	determines the width of the window that is rendered (default: 1920 pixels)
3. -h &lt;height pixel count&gt; 	determines the height of the window that is rendered (default: 1080 pixels)
4. -s &lt;scale factor&gt;		determines how much the actual rendered image is scaled 
                                    (if you have black borders on your fullscreen window, increase this (default: 3.0)
</pre>

The controls for the game are as follows:
1. 'W' moves the player forward
2. 'S' moves the player backward
3. 'A' turns the player left
4. 'D' turns the player right
5. 'LShift' activates "sprint" while active where the player moves and turns faster
6. 'Esc' exits the game with one keypress


***

* Note: The file "windowFrame.go" was directly taken from the demo as the math/code does not make enough sense to me with the lack of commentary for me to be able to understand and break down line-by-line what is happening. For this reason, this is commented to the best of my ability (deserving of full credit), but not everything is commented (because I do not understand the code here).
* Note: There is more content in the demo of this project (https://github.com/faiface/pixel-examples/tree/master/community/raycaster); however, not all of this was brought over because I was typing this code and thinking about this code myself while using this project merely as a reference for 85% of it and using a function directly from it in the "windowFrame.go" file.