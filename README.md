udp-hub  -- forward UDP data to multiple destinations
-----------------------------------------------------

This is a little UDP forwarder I wrote mostly for fun, and because I basically
suck at iptables and other proper networking/firewall configuration stuff.

The idea to receive UDP datagrams which could then be sent to multiple
destinations sprang up because of some ephemeral data streams we were sending
at work, which became so useful we wanted to send them to multiple hosts.

Reliability is obviously not paramount.

Example Configuration
---------------------
```
{ "Maps":
   [
      {
         "SrcPort": 7070,
         "DstAddr": [
            {"Host": "localhost", "Port": 7071},
            {"Host": "localhost", "Port": 7072}
         ]
      }
   ]
}
```
