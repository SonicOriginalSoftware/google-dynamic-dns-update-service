define SERVICE_CONTENT
#!/sbin/openrc-run

description="Updates Google Domains Dynamic DNS with current IP Address"

export GDDNS_IP_URL=$(GDDNS_IP_URL)
export GDDNS_FREQUENCY=$(GDDNS_FREQUENCY)
export GDDNS_HOSTNAMES=$(GDDNS_HOSTNAMES)
export GDDNS_USERNAMES=$(GDDNS_USERNAMES)
export GDDNS_PASSWORDS=$(GDDNS_PASSWORDS)

command="$(EXE_NAME)"
command_background=true
pidfile="/run/$(EXE_NAME).pid"
output_log="/var/log/$(EXE_NAME).log"
error_log="/var/log/$(EXE_NAME).log"
endef

export SERVICE_CONTENT
