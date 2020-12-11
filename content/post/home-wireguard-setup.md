---
title: home Wireguard setup
date: 2020-03-08T14:49:48-05:00
tags: [operations]
---

Here's a quick guide to setting up Wireguard to connect your Android phone to your home network.

<!--more-->

## Server setup

Do all of this as root.

1. Install Wireguard
1. `mkdir /etc/wireguard` if it doesn't already exist
1. `cd /etc/wireguard`
1. `umask 077`
1. `wg genkey | tee privatekey | wg pubkey > publickey`
1. `vim wg0.conf`

wg0.conf should have these contents:

```
[Interface]
Address = 192.168.2.1  # this is the IP the server will have. Choose an unused subnet.
PrivateKey = <server's privatekey (not path to it, the actual private key)>
ListenPort = 51820 # this is the default Wireguard port
PostUp   = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o enp2s5 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o enp2s5 -j

[Peer]
PublicKey = <client's publickey (actual private key)> # You will set the in the next part
AllowedIPs = 192.168.2.2/32 # based on the subnet you've chosen in the address above
```

Replace `enp2s5` with your actual device name. You can get this with `ip link show`.

## Client setup

1. Install the Wireguard Android app
1. Touch "Create from scratch" under the plus menu
1. Name it
1. Generate a private and public key pair
1. Copy the public key somehow to the server (email?)
1. Set `Addresses` to be in the subnet, e.g. `192.168.2.2`

Then add the sever as a peer by tapping "Add Peer":
1. Add the server's public key (it's in `/etc/wireguard/publickey`, email it to yourself again)
1. In `Allowed IPs` put `0.0.0.0/0, ::/0`
1. In `Endpoint` put in your server's IP or Host, like `my.server.com:51820`. This port should match the port configured on the server `ListenPort`.
1. Optionally set `Persistant keepalive` to a value (e.g. `25`) if you want the connection to send keepalives.

That's it! Try to connect. If it worked it will say `Transfer` and how much data has transferred.

If it doesn't work check:
1. Your router should be forwarding port 51820 to your server.
1. Your server's firewall should be allowing port 51820. Just UDP I think?

You'll also want this to start on the server at startup, so do this:

```
sudo chown -R root:root /etc/wireguard/
sudo chmod -R og-rwx /etc/wireguard/*
sudo systemctl enable wg-quick@wg0.service
```

## Additional clients

You can setup additional clients with a config like this:

```
[Interface]
PrivateKey = <new client's private key>
ListenPort = 21841
Address = 192.168.2.3/32 # always use a new IP in the subnet

[Peer]
PublicKey = <the server's public key>
AllowedIPs = 192.168.2.0/24, 192.168.1.0/24 # you can also specify just which IPs to send through the VPN like this
Endpoint = my.server.com:51820
```
