# Stoplight


## Table of Contents
- [Overview](#overview)
- [Road Map](#road-map)


## Overview

Stoplight is a Network Traffic Analyzer built in Go. The intent of this project is to build an easy to use tool for wardriving and network monitoring.

Stoplight implements [Cobra](https://github.com/spf13/cobra-cli/tree/main) to create a CLI based toolfor more technical users while also implementing a Terminal User Interface (TUI) and a localhost webpage.


## Road Map

1. Captue Network Packets
The MVP is the capturing and printing of network packets.
*This is mostly complete however there are layers that still need to be implemented*

2. User Profiles
People should be able to register their user profile and then when they spin up a monitor instance it should connect the instance with their profile.

3. Packet Storage
Packets should be stored in a postgres database to allow for analytical dashboards to draw insights

4. Session Sharing
Users should be able to name sessions when starting monitor and retroactively and then they should be able to share those sessions with other users

6. Data Export
The user should be able to export the data in .pcap and .csv for user profiles as a whole and individual sessions.

7. War Drving Functionality
Add gps data associated with the location of the packet capture and allow for device selection which will put the device in monitor mode

5. TUI Development
Introduce BubbleTea and LipGloss to create an interactive TUI which should make the user experience smoother

8. Device Connection Functionality
Allow the user to connect multiple devices to the server and it will montior traffic on all the devices

9. IAM Functonality
Groups and Organizations will be able to be made and users can be assigned to these groups and orgs with various capabilities.

10. Build a Front End
Transition to a web browser with dashboards and visuals
