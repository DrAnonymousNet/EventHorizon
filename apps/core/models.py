import csv
import json
import uuid

from django.db import models
from model_utils.models import SoftDeletableModel, TimeStampedModel


class AbstractUUIDModel(models.Model):
    uid = models.UUIDField(default=uuid.uuid4, editable=False)

    class Meta:
        abstract = True

    def __str__(self):
        return f"{self._meta.verbose_name} - {self.uid}"


class AbstractBaseModelMinimal(TimeStampedModel, AbstractUUIDModel):
    """
    Base model with "uid", "created", "modified" fields. (TimeStamped)
    """

    class Meta:
        abstract = True


class AbstractBaseModel(SoftDeletableModel, TimeStampedModel, AbstractUUIDModel):
    """
    Base model with "uid", "created", "modified", "is_removed" fields. (SoftDeletable)
    """

    class Meta:
        abstract = True
