```
 ____ _  _ ____ _  _ ____ ___ 
 |--- |__| |___ |-:_ |=== |__>
```

This simple tool takes a list of directories / file locations of packet capture files (pcap, pcapng), loads the packets, and mashes them into one simple packet capture file, here is an example of its usage 

```
go run main.go Filepaths.txt Output.pcap
```

**Example Filepaths.txt**

```
Pcap_Examples/HTTP - JPG Download.pcap
Pcap_Examples/Telnet.pcap
Pcap_Examples/Ftp.pcap
Pcap_Examples/aaa.pcap
```

This will take every single packet of every file and mash it into one packet output, this is a side tool for myself to make it easier on automatic packet parsers such as my own whicb take a single packet capture file at a time and parses through things such as credentials, telnet, ssh, ftp, smtp, smb, HTTP and other protocols. 
