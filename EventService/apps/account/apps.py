from django.apps import AppConfig


class AccountConfig(AppConfig):
    default_auto_field = "django.db.models.BigAutoField"
    name = "apps.account"

    def ready(self) -> None:
        from apps.account.api.v1 import serializers

        return super().ready()
