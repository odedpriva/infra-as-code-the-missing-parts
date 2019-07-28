
NAME := vpc

COMMAND := terraform
PLAN = plan.tfplan
SHOW = show.json
CODE = main.tf


all: test

plan: $(PLAN)

$(PLAN): $(CODE)
	$(COMMAND) plan -out $(PLAN)

show: $(SHOW)

$(SHOW): plan
	$(COMMAND) show -json $(PLAN) > $(SHOW)

test: show
	cat $(SHOW) | conftest test -

clean:
	@rm -f $(PLAN) $(SHOW)

.PHONY: plan show test all clean