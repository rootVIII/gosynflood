
### gosynflood - Repeatedly Send Crafted TCP SYN Packets with Raw Sockets

###### intended for Ubuntu and other Debian distributions

###### USAGE:
<pre>
  <code>
# Clone project:
git clone https://github.com/rootVIII/gosynflood.git

# Build and run:
cd &lt;project root&gt;
go build -o bin/gosynflood

# raw sockets require root privileges when executing:
sudo ./bin/gosynflood  -t &lt;target IPV4 address&gt; -p &lt;port number&gt; -i &lt;network interface&gt;

# Example:
sudo ./bin/gosynflood  -t 192.168.1.120 -p 80 -i wlp3s0
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
<img src="https://user-images.githubusercontent.com/30498791/166178559-b8ad5922-b12c-4ce0-883c-7205a22721b8.png">


2. The attacker machine (a separate physical machine also running Ubuntu) executes the gosynflood exe with root privileges:
<img src="https://user-images.githubusercontent.com/30498791/166178555-53635637-9b62-4afb-af36-f36f0354902d.png">


3. The initial SYNs are visible in Wireshark on the target, purposefully never completing the thee 3-way handshake:
<img src="https://user-images.githubusercontent.com/30498791/166178558-fffa57f5-3083-4af2-9a7e-63225cd23d7c.png">
<img src="https://user-images.githubusercontent.com/30498791/166178557-d394bd6f-9a16-4711-bbf6-8c46c5847aa5.png">


4. During the attack the webserver should be unreachable at it's URL if it is susceptible. The half-open connections are visible via the command <code>netstat -na --tcp</code>
<img src="https://user-images.githubusercontent.com/30498791/166178556-ea9fa18e-4f90-40c3-b2dc-8bc748cd7632.png">




This was developed on Ubuntu 18.04 LTS.
<hr>
<b>Author: rootVIII  2020</b>
<br><br>