# Stoplight


## Table of Contents
- [Overview](#overview)
- [Road Map](#road-map)


## Overview

Stoplight is a Network Traffic Analyzer built in Go. The intent of this project is to build an easy to use tool for wardriving and network monitoring.


## Road Map

1. Captue Network Packets
The MVP is the capturing and printing of network packets.

2. Packet Storage
Packets should be stored in a postgres database to allow for analytical dashboards to draw insights

3. Data Export
The user should be able to export the data in .pcap and .csv

4. War Drving Functionality
Add gps data to the database and save sessions allowing for the naming of sessions

5. Device Connection Functionality
Allow the user to connect multiple devices to the server and it will montior traffic on all the devices

6. Login & Register
Add user accounts with IAM functionality so various users can have various accesses

7. Build a Front End
Transition to a web browser with dashboards and visuals
