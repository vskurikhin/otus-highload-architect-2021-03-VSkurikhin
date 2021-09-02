WIFI := wifi
BLUETOOTH := bluetooth

COUNTER := app-counter
DIALOG := app-dialog
MAIN := app-main

all:
	cd $(COUNTER); make
	cd $(DIALOG); make
	cd $(MAIN); make
