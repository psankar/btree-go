#B Tree

B Trees are arguably the most important datastructure in the computer storage
area. Variants of B Trees have been widely used in filesystems, databases,
indexes and anything to do with disks for the last few decades.

There are many variants of B Trees such as the Knuth variant, B+ Trees, B\* Trees
etc.

This Go program implements the original B Tree that was invented by Rudolf Bayer
and Edward McCreight while in Boeing, and as documented by Douglas Comer in his
seminal paper *The ubiquitous B Tree* [www.cs.aau.dk/~simas/aalg06/UbiquitBtree.pdf], which I personally believe what every
programmer must read. To quote Donald Knuth, B Trees are the biggest
contribution from Computer Science to Mathematics.

I have implemented this Go program mainly as a teaching aid to cover the
original version of the B Tree. Most of the online visualization projects
cover the Knuth variant and not the original version.  On hindsight, 
I believe that I should have implemented this in Javascript and contributed 
to the tree visualization project maintained by David Galles from
the University of San Francisco [http://www.cs.usfca.edu/~galles/visualization/]

#Usage
* Install go from http://golang.org
* go run btree.go
* Visit http://localhost:8080 and insert/delete numbers

#Sources
I have used go to maintain the BTree datastructure. I generate DoT markup files
from my Go program which I render in the browser using the viz.js, that I have
copied and used.

The Go program is completely written by me and is licensed under Creative Commons Zero License.
More information at: http://creativecommons.org/publicdomain/zero/1.0/

Send your patches / merge-requests only if you are ready to release your changes
to public domain.

The viz.js is licensed under https://github.com/mdaines/viz.js/blob/master/COPYING
