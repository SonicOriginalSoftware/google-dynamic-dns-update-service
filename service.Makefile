define SERVICE_CONTENT
#!/sbin/openrc-run

description="Updates Google Domains Dynamic DNS with current IP Address"

export GDDNS_HOSTNAME=$(GDDNS_HOSTNAME)
export GDDNS_IP_URL=$(GDDNS_IP_URL)
export GDDNS_USERNAME=$(GDDNS_USERNAME)
export GDDNS_PASSWORD=$(GDDNS_PASSWORD)
export GDDNS_FREQUENCY=$(GDDNS_FREQUENCY)

command="$(EXE_NAME)"
pidfile="/run/$(EXE_NAME).pid"
endef

export SERVICE_CONTENT
