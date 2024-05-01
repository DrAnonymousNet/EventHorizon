import uuid

from apps.core.models import AbstractBaseModel
from django.contrib.auth.models import AbstractUser
from django.db import models


class User(AbstractUser, AbstractBaseModel):
    USERNAME_FIELD = "username"
    REQUIRED_FIELDS = ["email"]

    def __str__(self):
        return f"User-{self.email}-{self.uid}"


class Device(AbstractBaseModel):
    user = models.ForeignKey(User, on_delete=models.CASCADE, related_name="devices")
    device_token = models.CharField(max_length=255, unique=True)

    def __str__(self) -> str:
        return f"Device-{self.user.email}-{self.uid}"
