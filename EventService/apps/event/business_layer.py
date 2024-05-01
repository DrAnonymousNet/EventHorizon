from apps.event.models import Event, EventAttendance
from config import constant, notification_channel_map
from django.template import Template
from django.template.loader import render_to_string
from eh_grpc.client import notification_service_client


class EmailService:
    @classmethod
    def generate_email(cls, template_name: str, context: dict) -> tuple:
        subject = context.get("subject")
        message = render_to_string(template_name, context)
        return subject, message

    @classmethod
    def rsvp_confirmation_to_attendee(
        cls, event: Event, attendee: EventAttendance
    ) -> tuple:
        context = {
            "event": event,
            "recipient_name": attendee.name,
            "organizer": event.organizer,
            "subject": "RSVP Confirmation",
        }
        return cls.generate_email("rsvp_confirmation.html", context)

    @classmethod
    def rsvp_notification_to_organizer(
        cls, event: Event, attendee: EventAttendance
    ) -> tuple:
        context = {
            "event": event,
            "attendee": attendee,
            "organizer": event.organizer,
            "subject": "New RSVP",
        }
        return cls.generate_email("rsvp_notification.html", context)

    @classmethod
    def event_updated(cls, event: Event) -> tuple:
        context = {"event": event, "subject": "Event Updated"}
        return cls.generate_email("event_updated.html", context)


# The notification will receive the the subject, the message, the user id, the event uid


def send_rsvp_message_to_notification_service(event: Event, attendee: EventAttendance):
    subject, message = EmailService.rsvp_confirmation_to_attendee(event, attendee)
    notification_service_client(
        str(event.uid),
        message=message,
        subject=subject,
        recipient=attendee.email,
        message_type=constant.RSVP_CONFIRMATION_TO_ATTENDEE,
        notification_channels=notification_channel_map.get(
            constant.RSVP_CONFIRMATION_TO_ATTENDEE
        ),
    )

    subject, message = EmailService.rsvp_notification_to_organizer(event, attendee)
    notification_service_client(
        str(event.uid),
        message=message,
        subject=subject,
        recipient=event.organizer.email,
        message_type=constant.RSVP_NOTIFICATION_TO_ORGANIZER,
        notification_channels=notification_channel_map.get(
            constant.RSVP_NOTIFICATION_TO_ORGANIZER
        ),
    )
