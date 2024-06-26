from functools import cached_property

from apps.core.models import AbstractBaseModel
from config.storage_backends import PublicMediaStorage
from django.contrib.auth import get_user_model
from django.contrib.postgres.fields import ArrayField
from django.db import models

# from django_prometheus.models import ExportModelOperationsMixin


User = get_user_model()


class Event(AbstractBaseModel):
    class EventStatus(models.TextChoices):
        PENDING = "PENDING", "Pending"
        ACTIVE = "ACTIVE", "Active"
        CANCELLED = "CANCELLED", "Cancelled"
        COMPLETED = "COMPLETED", "Completed"

    title = models.CharField(max_length=255)
    description = models.TextField(blank=True, null=True)
    organizer = models.ForeignKey(
        User, on_delete=models.CASCADE, related_name="organized_events"
    )
    location = models.CharField(max_length=255)
    start_time = models.DateTimeField()
    end_time = models.DateTimeField()
    max_participants = models.IntegerField(blank=True, null=True)
    is_virtual = models.BooleanField(default=False)
    virtual_meeting_link = models.URLField(blank=True, null=True)
    status = models.CharField(
        max_length=50, choices=EventStatus.choices, default=EventStatus.PENDING
    )

    def __str__(self):
        return self.title


class EventAttendance(AbstractBaseModel):
    class AttendanceType(models.TextChoices):
        PHYSICAL = "PHYSICAL", "Physical"
        VIRTUAL = "VIRTUAL", "Virtual"
        INTERESTED = "INTERESTED", "Interested"
        REGISTERED = "REGISTERED", "Registered"

    name = models.CharField(max_length=255, blank=True, null=True)
    event = models.ForeignKey(Event, on_delete=models.CASCADE, related_name="attendees")
    registered_user = models.ForeignKey(
        User,
        on_delete=models.CASCADE,
        null=True,
        blank=True,
        related_name="event_attendances",
    )
    unregistered_user_email = models.EmailField(max_length=255, blank=True, null=True)
    attendance_type = ArrayField(
        models.CharField(max_length=100, choices=AttendanceType.choices), default=list
    )

    class Meta:
        unique_together = ("event", "registered_user", "unregistered_user_email")
        constraints = [
            models.CheckConstraint(
                check=models.Q(registered_user__isnull=False)
                | models.Q(unregistered_user_email__isnull=False),
                name="check_attendance_has_user",
            ),
        ]

    def __str__(self):
        if self.registered_user:
            return f"{self.registered_user.username}\
            - {self.event.title} - {self.attendance_type}"
        else:
            return f"Unregistered: {self.unregistered_user_email} \
            - {self.event.title} - {self.attendance_type}"

    @cached_property
    def email(self):
        if self.registered_user:
            return self.registered_user.email
        else:
            self.unregistered_user_email


class EventImage(AbstractBaseModel):
    event = models.ForeignKey(Event, related_name="images", on_delete=models.CASCADE)
    image = models.ImageField(upload_to="event_images/", storage=PublicMediaStorage())
    caption = models.CharField(max_length=255, blank=True, null=True)

    def __str__(self):
        return f"Image for {self.event.title} - \
        {self.caption or 'No Caption'}"
