define SERVICE_CONTENT
#!/sbin/openrc-run

description="Updates Google Domains Dynamic DNS with current IP Address"

GDDNS_HOSTNAME=$(GDDNS_HOSTNAME)
GDDNS_IP_URL=$(GDDNS_IP_URL)
GDDNS_USERNAME=$(GDDNS_USERNAME)
GDDNS_PASSWORD=$(GDDNS_PASSWORD)
GDDNS_FREQUENCY=$(GDDNS_FREQUENCY)

command="$(EXE_NAME)"
pidfile="/run/$(EXE_NAME).pid"
endef

export SERVICE_CONTENT
