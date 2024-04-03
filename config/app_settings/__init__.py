import environ

BASE_DIR = environ.Path(__file__) - 3  # noqa F405
APPS_DIR = BASE_DIR.path("apps")
env = environ.Env()  # noqa F405

READ_DOT_ENV_FILE = env.bool("DJANGO_READ_DOT_ENV_FILE", default=True)
if READ_DOT_ENV_FILE:
    # OS environment variables take precedence over variables from .env
    env.read_env(str(BASE_DIR.path(".env")))

ENV = env.str("ENVIRONMENT")
LOCAL = False
TESTING = False
PRODUCTION = False
STAGING = False

if ENV == "local":
    from .local import *

    LOCAL = True
elif ENV == "test":
    from .test import *

    TESTING = True
elif ENV == "prod":
    from .prod import *

    PRODUCTION = True
    # elif ENV == "staging":
    #     from .staging import *

    STAGING = True
