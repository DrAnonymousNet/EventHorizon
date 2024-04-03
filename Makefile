SHELL := /bin/bash
# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif


SETTINGS=config.app_settings
# target: help - show all available commands
help:
	@egrep "^# target:" [Mm]akefile

# target: run - run the server
run $(arguments):
	python manage.py runserver_plus $(RUN_ARGS)

# target: ipy - activate the shell
ipy:
	python manage.py shell_plus $(RUN_ARGS)

# target: migrations - generate migrations
migrations:
	python manage.py makemigrations $(RUN_ARGS)

# target: migrate - migrate to db
migrate:
	python manage.py migrate $(RUN_ARGS)
	python manage.py createcachetable

# target: superuser - create superuser
superuser:
	python manage.py createsuperuser

# target: reference - create all reference database tables
reference:
	python manage.py populate_references $(RUN_ARGS)

# target: clean - remove all ".pyc" files
clean:
	python manage.py clean_pyc --settings=$(SETTINGS)

# target: collect - calls the "collectstatic" django command
collect:
	python manage.py collectstatic --settings=$(SETTINGS) --noinput

# target: reset_db: drop and recreate the database
reset_db:
	python manage.py reset_db --settings=$(SETTINGS) --router=default --noinput --close-sessions
