
### gosynflood - Repeatedly Send Crafted TCP SYN Packets with Raw Sockets

###### intended for Ubuntu and other Debian distributions

###### USAGE:
<pre>
  <code>
# go get or git clone the project
go get github.com/rootVIII/gosynflood
git clone https://github.com/rootVIII/gosynflood.git


# Navigate to project root and build:
go build .

# raw sockets require root privileges when executing:
sudo ./gosynflood -t &lt;target IPV4 address&gt; -p &lt;port number&gt; -i &lt;network interface&gt;

# Example:
sudo ./gosynflood -t 192.168.1.120 -p 80 -i wlp3s0

** The bin/builds directory has a compiled executable if you do not have
Golang installed or do not want to build it yourself.
  </code>
</pre>

###### CLI OPTIONS:
<pre>
  <code>
-t private or public IP address of target webserver
-p target webserver's port number (defaults to port 80 if not provided)
-i your network interface (running without -i will fail,
     but it will suggest all found interfaces, ie: lo, wlpxxx, eth0 etc.)

Enter control-c to stop the flood attack.
  </code>
</pre>

Each packet's IP address is spoofed. <b>MAC addresses are not spoofed</b>.
It is up to you to spoof your MAC Address beforehand if desired.

This attack may only work on web servers susceptible to numerous half-open connections (SYN_RECV).

To demonstrate this, a small Ubuntu Mate running Apache2 will act as the target.
It's a physical machine on a private network.

<hr>
1. The tcp_syncookies flag was set to 0 (to make the target vulnerable for demonstration purposes) and the webserver was started on the target:
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/1.png">


2. The attacker machine (a separate physical machine also running Ubuntu) executes the gosynflood exe with root privileges:
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/5.png">


3. The initial SYNs are visible in Wireshark on the target, purposefully never completing the thee 3-way handshake:
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/2.png">
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/3.png">


4. During the attack the webserver should be unreachable at it's URL if it is susceptible. The half-open connections are visible via the command <code>netstat -na --tcp</code>
<img src="https://github.com/rootVIII/gosynflood/blob/master/bin/screenshots/4.png">




This was developed on Ubuntu 18.04 LTS.
<hr>
<b>Author: rootVIII  2020</b>
<br><br>