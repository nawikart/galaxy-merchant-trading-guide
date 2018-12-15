# GALAXY MERCHANT TRADING GUIDE


## Build Setup

``` bash
#serve golang server
$ go run main.go (for other than windows)
or
$ go build
$ golang.exe (for windows)


```

###### run the application in browser
go to http://localhost:8082/ on browser
<br/><br/>


## DEMO: http://172.104.185.200:8082
<br/>


## Screenshot

Screenshot 1
![alt text](https://raw.githubusercontent.com/nawikart/README_IMAGES/master/galaxy/ss1.png)
<br/><br/><br/><br/>


Screenshot 2
![alt text](https://raw.githubusercontent.com/nawikart/README_IMAGES/master/galaxy/ss2.png)
<br/><br/><br/><br/>


Screenshot 3
![alt text](https://raw.githubusercontent.com/nawikart/README_IMAGES/master/galaxy/ss3.png)

<br/><br/><br/><br/>
### "GoVersion": "go1.10"

<br/><br/>
## Problem Description
You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you. The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them. Roman numerals are based on seven symbols:

Symbol Value
I 1<br/>
V 5<br/>
X 10<br/>
L 50<br/>
C 100<br/>
D 500<br/>
M 1,000<br/>

Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.)

"D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only.

"V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately.

In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.
-- Source: Wikipedia (http://en.wikipedia.org/wiki/Roman_numerals)
<br/>

Test Input:
-------------
glob is I<br/>
prok is V<br/>
pish is X<br/>
tegj is L<br/>

glob glob Silver is 34 Credits<br/>
glob prok Gold is 57800 Credits<br/>
pish pish Iron is 3910 Credits<br/>
how much is pish tegj glob glob ?<br/>
how many Credits is glob prok Silver ?<br/>
how many Credits is glob prok Gold ?<br/>
how many Credits is glob prok Iron ?<br/>
how much wood could a woodchuck chuck if a woodchuck could chuck wood?<br/>

Expecting Output:
---------------
pish tegj glob glob is 42<br/>
glob prok Silver is 68 Credits<br/>
glob prok Gold is 57800 Credits<br/>
glob prok Iron is 782 Credits<br/>
I have no idea what you are talking about<br/>

