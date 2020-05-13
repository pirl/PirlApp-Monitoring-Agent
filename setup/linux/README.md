## pirlapp-agent.conf

This is the agent configuration file. The default port is tcp/8081. The pirlapp-agent.conf is mandatory for the start.

```PIRLAPPPORT="8081"```

There you can change the listening port to a desired one. 

The file should be placed here: ```/etc/pirl/pirlapp-agent.conf```

## pirlapp-monitoring-agent.service

This file is for managing the monitoring agent daemon with systemd. It takes care that the process  is running or 
started after system boot.

The file should be placed here: ```/lib/systemd/system/pirlapp-monitoring-agent.service``` Service definition: 


```
[Unit]
Description=PirlApp Monitoring Agent
After=network.target

[Service]
EnvironmentFile=/etc/pirl/pirlapp-agent.conf
User=daemon
Group=daemon
ExecStart=/usr/bin/pirlapp-agent
Restart=always
RemainAfterExit=no
RestartSec=5s
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

You get control over this service with the following commands:

```
systemctl start pirlapp-monitoring-agent.service       # starts the service
systemctl stop pirlapp-monitoring-agent.service        # stop the service
systemctl status pirlapp-monitoring-agent.service      # get the status of the service
systemctl enable pirlapp-monitoring-agent.service      # enable service, should survive a boot
systemctl disable pirlapp-monitoring-agent.service     # disable service, shouldn´t survive a boot
```



