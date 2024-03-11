BIN_NAME = office-days-agent
BIN_DIR = /usr/local/bin
INSTALL_DIR = /usr/local/$(BIN_NAME)
PLIST_NAME = hu.okki.office-days-agent.plist
PLIST_DIR = /Library/LaunchDaemons

build:
	go build -o $(BIN_NAME) -buildvcs=false

clean:
	rm -f $(BIN_NAME)

install: build
	mkdir -p $(INSTALL_DIR)/
	cp $(BIN_NAME) $(INSTALL_DIR)/
	chmod 755 $(INSTALL_DIR)/$(BIN_NAME)
	ln -sf $(INSTALL_DIR)/$(BIN_NAME) $(BIN_DIR)/$(BIN_NAME)
	cp etc/$(PLIST_NAME) $(PLIST_DIR)/
	chmod 644 $(PLIST_DIR)/$(PLIST_NAME)

uninstall:
	rm -rf $(INSTALL_DIR)/
	rm -rf $(BIN_DIR)/$(BIN_NAME)
	rm -f $(PLIST_DIR)/$(PLIST_NAME)